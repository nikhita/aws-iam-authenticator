/*
Copyright 2017 by the contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

// RoleMapping is a mapping of an AWS Role ARN to a Kubernetes username and a
// list of Kubernetes groups. The username and groups are specified as templates
// that may optionally contain two template parameters:
//
//  1) "{{AccountID}}" is the 12 digit AWS ID.
//  2) "{{SessionName}}" is the role session name.
//
// The meaning of SessionName depends on the type of entity assuming the role.
// In the case of an EC2 instance role this will be the EC2 instance ID. In the
// case of a federated role it will be the federated identity (controlled by the
// federated identity provider). In the case of a role assumed directly with
// sts:AssumeRole it will be user controlled.
//
// You can use plain values without parameters to have a more static mapping.
type RoleMapping struct {
	// RoleARN is the AWS Resource Name of the role. (e.g., "arn:aws:iam::000000000000:role/Foo").
	RoleARN string

	// Username is the username pattern that this instances assuming this
	// role will have in Kubernetes.
	Username string

	// Groups is a list of Kubernetes groups this role will authenticate
	// as (e.g., `system:masters`). Each group name can include placeholders.
	Groups []string
}

// UserMapping is a static mapping of a single AWS User ARN to a
// Kubernetes username and a list of Kubernetes groups
type UserMapping struct {
	// UserARN is the AWS Resource Name of the user. (e.g., "arn:aws:iam::000000000000:user/Test").
	UserARN string

	// Username is the Kubernetes username this role will authenticate as (e.g., `mycorp:foo`)
	Username string

	// Groups is a list of Kubernetes groups this role will authenticate as (e.g., `system:masters`)
	Groups []string
}

// Config specifies the configuration for a heptio-authenticator-aws server
type Config struct {
	// ClusterID is a unique-per-cluster identifier for your
	// heptio-authenticator-aws installation.
	ClusterID string

	// LocalhostPort is the TCP on which to listen for authentication checks
	// (on localhost).
	LocalhostPort int

	// GenerateKubeconfigPath is the output path where a generated webhook
	// kubeconfig (for `--authentication-token-webhook-config-file`) will be
	// stored.
	GenerateKubeconfigPath string

	// StateDir is the directory where generated certificates and private keys
	// will be stored. You want these persisted between runs so that your API
	// server webhook configuration doesn't change on restart.
	StateDir string

	// RoleMappings is a list of mappings from AWS IAM Role to
	// Kubernetes username + groups.
	RoleMappings []RoleMapping

	// UserMappings is a list of mappings from AWS IAM User to
	// Kubernetes username + groups.
	UserMappings []UserMapping

	// AWS Accounts that are allowed without an explicit user/role mapping
	// The IAM ARN maps to the user or role name automatically.
	MappedAccounts []string
}
