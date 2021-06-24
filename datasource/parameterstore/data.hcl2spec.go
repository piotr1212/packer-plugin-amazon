// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package parameterstore

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-amazon/builder/common"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName       *string                           `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType     *string                           `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion     *string                           `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug           *bool                             `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce           *bool                             `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError         *string                           `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars        map[string]string                 `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars   []string                          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	AccessKey             *string                           `mapstructure:"access_key" required:"true" cty:"access_key" hcl:"access_key"`
	AssumeRole            *common.FlatAssumeRoleConfig      `mapstructure:"assume_role" required:"false" cty:"assume_role" hcl:"assume_role"`
	CustomEndpointEc2     *string                           `mapstructure:"custom_endpoint_ec2" required:"false" cty:"custom_endpoint_ec2" hcl:"custom_endpoint_ec2"`
	CredsFilename         *string                           `mapstructure:"shared_credentials_file" required:"false" cty:"shared_credentials_file" hcl:"shared_credentials_file"`
	DecodeAuthZMessages   *bool                             `mapstructure:"decode_authorization_messages" required:"false" cty:"decode_authorization_messages" hcl:"decode_authorization_messages"`
	InsecureSkipTLSVerify *bool                             `mapstructure:"insecure_skip_tls_verify" required:"false" cty:"insecure_skip_tls_verify" hcl:"insecure_skip_tls_verify"`
	MaxRetries            *int                              `mapstructure:"max_retries" required:"false" cty:"max_retries" hcl:"max_retries"`
	MFACode               *string                           `mapstructure:"mfa_code" required:"false" cty:"mfa_code" hcl:"mfa_code"`
	ProfileName           *string                           `mapstructure:"profile" required:"false" cty:"profile" hcl:"profile"`
	RawRegion             *string                           `mapstructure:"region" required:"true" cty:"region" hcl:"region"`
	SecretKey             *string                           `mapstructure:"secret_key" required:"true" cty:"secret_key" hcl:"secret_key"`
	SkipMetadataApiCheck  *bool                             `mapstructure:"skip_metadata_api_check" cty:"skip_metadata_api_check" hcl:"skip_metadata_api_check"`
	SkipCredsValidation   *bool                             `mapstructure:"skip_credential_validation" cty:"skip_credential_validation" hcl:"skip_credential_validation"`
	Token                 *string                           `mapstructure:"token" required:"false" cty:"token" hcl:"token"`
	VaultAWSEngine        *common.FlatVaultAWSEngineOptions `mapstructure:"vault_aws_engine" required:"false" cty:"vault_aws_engine" hcl:"vault_aws_engine"`
	PollingConfig         *common.FlatAWSPollingConfig      `mapstructure:"aws_polling" required:"false" cty:"aws_polling" hcl:"aws_polling"`
	Name                  *string                           `mapstructure:"name" required:"true" cty:"name" hcl:"name"`
	WithDecryption        *bool                             `mapstructure:"with_decryption" cty:"with_decryption" hcl:"with_decryption"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":             &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":           &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":           &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":                  &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                  &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":               &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":         &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":    &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"access_key":                    &hcldec.AttrSpec{Name: "access_key", Type: cty.String, Required: false},
		"assume_role":                   &hcldec.BlockSpec{TypeName: "assume_role", Nested: hcldec.ObjectSpec((*common.FlatAssumeRoleConfig)(nil).HCL2Spec())},
		"custom_endpoint_ec2":           &hcldec.AttrSpec{Name: "custom_endpoint_ec2", Type: cty.String, Required: false},
		"shared_credentials_file":       &hcldec.AttrSpec{Name: "shared_credentials_file", Type: cty.String, Required: false},
		"decode_authorization_messages": &hcldec.AttrSpec{Name: "decode_authorization_messages", Type: cty.Bool, Required: false},
		"insecure_skip_tls_verify":      &hcldec.AttrSpec{Name: "insecure_skip_tls_verify", Type: cty.Bool, Required: false},
		"max_retries":                   &hcldec.AttrSpec{Name: "max_retries", Type: cty.Number, Required: false},
		"mfa_code":                      &hcldec.AttrSpec{Name: "mfa_code", Type: cty.String, Required: false},
		"profile":                       &hcldec.AttrSpec{Name: "profile", Type: cty.String, Required: false},
		"region":                        &hcldec.AttrSpec{Name: "region", Type: cty.String, Required: false},
		"secret_key":                    &hcldec.AttrSpec{Name: "secret_key", Type: cty.String, Required: false},
		"skip_metadata_api_check":       &hcldec.AttrSpec{Name: "skip_metadata_api_check", Type: cty.Bool, Required: false},
		"skip_credential_validation":    &hcldec.AttrSpec{Name: "skip_credential_validation", Type: cty.Bool, Required: false},
		"token":                         &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"vault_aws_engine":              &hcldec.BlockSpec{TypeName: "vault_aws_engine", Nested: hcldec.ObjectSpec((*common.FlatVaultAWSEngineOptions)(nil).HCL2Spec())},
		"aws_polling":                   &hcldec.BlockSpec{TypeName: "aws_polling", Nested: hcldec.ObjectSpec((*common.FlatAWSPollingConfig)(nil).HCL2Spec())},
		"name":                          &hcldec.AttrSpec{Name: "name", Type: cty.String, Required: false},
		"with_decryption":               &hcldec.AttrSpec{Name: "with_decryption", Type: cty.Bool, Required: false},
	}
	return s
}

// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatDatasourceOutput struct {
	Value   *string `mapstructure:"value" cty:"value" hcl:"value"`
	Version *string `mapstructure:"version" cty:"version" hcl:"version"`
	ARN     *string `mapstructure:"arn" cty:"arn" hcl:"arn"`
}

// FlatMapstructure returns a new FlatDatasourceOutput.
// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*DatasourceOutput) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatDatasourceOutput)
}

// HCL2Spec returns the hcl spec of a DatasourceOutput.
// This spec is used by HCL to read the fields of DatasourceOutput.
// The decoded values from this spec will then be applied to a FlatDatasourceOutput.
func (*FlatDatasourceOutput) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"value":   &hcldec.AttrSpec{Name: "value", Type: cty.String, Required: false},
		"version": &hcldec.AttrSpec{Name: "version", Type: cty.String, Required: false},
		"arn":     &hcldec.AttrSpec{Name: "arn", Type: cty.String, Required: false},
	}
	return s
}
