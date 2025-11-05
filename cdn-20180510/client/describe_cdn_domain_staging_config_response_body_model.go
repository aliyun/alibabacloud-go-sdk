// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainStagingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigs(v []*DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) *DescribeCdnDomainStagingConfigResponseBody
	GetDomainConfigs() []*DescribeCdnDomainStagingConfigResponseBodyDomainConfigs
	SetDomainName(v string) *DescribeCdnDomainStagingConfigResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeCdnDomainStagingConfigResponseBody
	GetRequestId() *string
}

type DescribeCdnDomainStagingConfigResponseBody struct {
	// The domain name configurations.
	DomainConfigs []*DescribeCdnDomainStagingConfigResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Repeated"`
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C80705BF-0F76-41FA-BAD1-5B59296A4E59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnDomainStagingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainStagingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainStagingConfigResponseBody) GetDomainConfigs() []*DescribeCdnDomainStagingConfigResponseBodyDomainConfigs {
	return s.DomainConfigs
}

func (s *DescribeCdnDomainStagingConfigResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnDomainStagingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnDomainStagingConfigResponseBody) SetDomainConfigs(v []*DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) *DescribeCdnDomainStagingConfigResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBody) SetDomainName(v string) *DescribeCdnDomainStagingConfigResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBody) SetRequestId(v string) *DescribeCdnDomainStagingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBody) Validate() error {
	if s.DomainConfigs != nil {
		for _, item := range s.DomainConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnDomainStagingConfigResponseBodyDomainConfigs struct {
	// The configuration ID.
	//
	// example:
	//
	// 6xx5
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The description of each feature.
	FunctionArgs []*DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Repeated"`
	// The feature name.
	//
	// example:
	//
	// aliauth
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The rule condition ID. This parameter is optional. To create a rule condition, you can configure the **condition*	- feature that is described in the [Parameters for configuring features for domain names](https://help.aliyun.com/document_detail/388460.html) topic. A rule condition can identify parameters that are included in requests and filter requests based on the identified parameters. Each rule condition has a [ConfigId](https://help.aliyun.com/document_detail/388994.html). You can reference ConfigId instead of ParentId in other features. This way, you can combine rule conditions and features for flexible configurations. For more information, see [BatchSetCdnDomainConfig](https://help.aliyun.com/document_detail/90915.html) or ParentId configuration example in this topic.
	//
	// example:
	//
	// 222728944812032
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The configuration status. Valid values:
	//
	// 	- **testing**
	//
	// 	- **configuring**
	//
	// 	- **success**
	//
	// 	- **failed**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) GetFunctionArgs() []*DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	return s.FunctionArgs
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) GetStatus() *string {
	return s.Status
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) SetConfigId(v string) *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs {
	s.ConfigId = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) SetFunctionArgs(v []*DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs {
	s.FunctionArgs = v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) SetFunctionName(v string) *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs {
	s.FunctionName = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) SetParentId(v string) *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs {
	s.ParentId = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) SetStatus(v string) *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs {
	s.Status = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigs) Validate() error {
	if s.FunctionArgs != nil {
		for _, item := range s.FunctionArgs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs struct {
	// The configuration name.
	//
	// example:
	//
	// auth_type
	ArgName *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	// The configuration value.
	//
	// example:
	//
	// req_auth
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) SetArgName(v string) *DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgName = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) SetArgValue(v string) *DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs {
	s.ArgValue = &v
	return s
}

func (s *DescribeCdnDomainStagingConfigResponseBodyDomainConfigsFunctionArgs) Validate() error {
	return dara.Validate(s)
}
