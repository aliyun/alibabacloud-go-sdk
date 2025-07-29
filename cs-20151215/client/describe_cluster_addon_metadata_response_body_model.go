// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterAddonMetadataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigSchema(v string) *DescribeClusterAddonMetadataResponseBody
	GetConfigSchema() *string
	SetName(v string) *DescribeClusterAddonMetadataResponseBody
	GetName() *string
	SetVersion(v string) *DescribeClusterAddonMetadataResponseBody
	GetVersion() *string
}

type DescribeClusterAddonMetadataResponseBody struct {
	// The component schema parameters.
	//
	// example:
	//
	// {\\n  \\"$schema\\": \\"https://json-schema.org/draft-07/schema#\\",\\n  \\"properties\\": {\\n    \\"controller\\": {\\n      \\"description\\": \\"\\",\\n      \\"properties\\": {\\n        \\"resources\\": {\\n          \\"properties\\": {\\n            \\"armsPrometheusOperator\\": {\\n              \\"properties\\": {\\n                \\"resources\\": {\\n                  \\"properties\\": {\\n                    \\"limits\\": {\\n                      \\"properties\\": {\\n                        \\"memory\\": {\\n                          \\"description\\": \\"memory limit of arms prometheus operator\\",\\n                          \\"type\\": \\"string\\",\\n                          \\"pattern\\": \\"^[1-9][0-9]*(\\\\\\\\.\\\\\\\\d+)?(K|Ki|M|Mi|G|Gi|T|Ti)?$\\",\\n                          \\"default\\": \\"500m\\",\\n                          \\"x-ui-description\\": \\"<mds-key>\\",\\n                          \\"x-ui-prompt-message\\": \\"<mds-key>\\",\\n                          \\"x-ui-validation-message\\": \\"<mds-key>\\",\\n                          \\"x-ui-additional-tips\\": \\"<mds-key>\\"\\n                        },\\n                        \\"cpu\\": {\\n                          \\"description\\": \\"cpu limit of arms prometheus operator\\",\\n                          \\"type\\": \\"string\\",\\n                          \\"pattern\\": \\"^[1-9][0-9]*(m|\\\\\\\\.\\\\\\\\d+)?$\\",\\n                          \\"default\\": \\"1.0\\",\\n                          \\"x-ui-description\\": \\"<mds-key>\\",\\n                          \\"x-ui-validation-message\\": \\"<mds-key>\\"\\n                        }\\n                      },\\n                      \\"type\\": \\"object\\",\\n                      \\"additionalProperties\\": false\\n                    },\\n                    \\"requests\\": {\\n                      \\"properties\\": {\\n                        \\"memory\\": {\\n                          \\"description\\": \\"memory request of arms prometheus operator\\",\\n                          \\"type\\": \\"string\\",\\n                          \\"pattern\\": \\"^[1-9][0-9]*(\\\\\\\\.\\\\\\\\d+)?(K|Ki|M|Mi|G|Gi|T|Ti)?$\\",\\n                          \\"default\\": \\"500m\\",\\n                          \\"x-ui-description\\": \\"<mds-key>\\",\\n                          \\"x-ui-validation-message\\": \\"<mds-key>\\"\\n                        },\\n                        \\"cpu\\": {\\n                          \\"description\\": \\"cpu request of arms prometheus operator\\",\\n                          \\"type\\": \\"string\\",\\n                          \\"pattern\\": \\"^[1-9][0-9]*(m|\\\\\\\\.\\\\\\\\d+)?$\\",\\n                          \\"default\\": \\"1.0\\",\\n                          \\"x-ui-description\\": \\"<mds-key>\\",\\n                          \\"x-ui-validation-message\\": \\"<mds-key>\\"\\n                        }\\n                      },\\n                      \\"type\\": \\"object\\",\\n                      \\"additionalProperties\\": false\\n                    }\\n                  },\\n                  \\"type\\": \\"object\\",\\n                  \\"additionalProperties\\": false\\n                }\\n              },\\n              \\"type\\": \\"object\\",\\n              \\"additionalProperties\\": false\\n            }\\n          },\\n          \\"type\\": \\"object\\",\\n          \\"additionalProperties\\": false\\n        }\\n      },\\n      \\"type\\": \\"object\\",\\n      \\"additionalProperties\\": false\\n    }\\n  },\\n  \\"title\\": \\"Values\\",\\n  \\"type\\": \\"object\\",\\n  \\"additionalProperties\\": false\\n}
	ConfigSchema *string `json:"config_schema,omitempty" xml:"config_schema,omitempty"`
	// The component name.
	//
	// example:
	//
	// coredns
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The component version.
	//
	// example:
	//
	// 1.8.4.1
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeClusterAddonMetadataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterAddonMetadataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterAddonMetadataResponseBody) GetConfigSchema() *string {
	return s.ConfigSchema
}

func (s *DescribeClusterAddonMetadataResponseBody) GetName() *string {
	return s.Name
}

func (s *DescribeClusterAddonMetadataResponseBody) GetVersion() *string {
	return s.Version
}

func (s *DescribeClusterAddonMetadataResponseBody) SetConfigSchema(v string) *DescribeClusterAddonMetadataResponseBody {
	s.ConfigSchema = &v
	return s
}

func (s *DescribeClusterAddonMetadataResponseBody) SetName(v string) *DescribeClusterAddonMetadataResponseBody {
	s.Name = &v
	return s
}

func (s *DescribeClusterAddonMetadataResponseBody) SetVersion(v string) *DescribeClusterAddonMetadataResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeClusterAddonMetadataResponseBody) Validate() error {
	return dara.Validate(s)
}
