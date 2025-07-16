// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDomainConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainConfigs(v *DescribeCdnDomainConfigsResponseBodyDomainConfigs) *DescribeCdnDomainConfigsResponseBody
	GetDomainConfigs() *DescribeCdnDomainConfigsResponseBodyDomainConfigs
	SetRequestId(v string) *DescribeCdnDomainConfigsResponseBody
	GetRequestId() *string
}

type DescribeCdnDomainConfigsResponseBody struct {
	// The configurations of the domain name.
	DomainConfigs *DescribeCdnDomainConfigsResponseBodyDomainConfigs `json:"DomainConfigs,omitempty" xml:"DomainConfigs,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// C80705BF-0F76-41FA-BAD1-5B59296A4E59
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnDomainConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBody) GetDomainConfigs() *DescribeCdnDomainConfigsResponseBodyDomainConfigs {
	return s.DomainConfigs
}

func (s *DescribeCdnDomainConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnDomainConfigsResponseBody) SetDomainConfigs(v *DescribeCdnDomainConfigsResponseBodyDomainConfigs) *DescribeCdnDomainConfigsResponseBody {
	s.DomainConfigs = v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBody) SetRequestId(v string) *DescribeCdnDomainConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainConfigsResponseBodyDomainConfigs struct {
	DomainConfig []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig `json:"DomainConfig,omitempty" xml:"DomainConfig,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigs) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigs) GetDomainConfig() []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	return s.DomainConfig
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigs) SetDomainConfig(v []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) *DescribeCdnDomainConfigsResponseBodyDomainConfigs {
	s.DomainConfig = v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigs) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig struct {
	// The ID of the configuration.
	//
	// example:
	//
	// 6295
	ConfigId *string `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The configuration of each feature.
	FunctionArgs *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs `json:"FunctionArgs,omitempty" xml:"FunctionArgs,omitempty" type:"Struct"`
	// The name of the feature.
	//
	// example:
	//
	// aliauth
	FunctionName *string `json:"FunctionName,omitempty" xml:"FunctionName,omitempty"`
	// The ID of the rule condition. This parameter is optional.
	//
	// To create a rule condition, you can configure the **condition*	- feature that is described in the [Parameters for configuring features for domain names](https://help.aliyun.com/document_detail/388460.html) topic. A rule condition can identify parameters that are included in requests and filter requests based on the identified parameters. Each rule condition has a [ConfigId](https://help.aliyun.com/document_detail/388994.html). You can use ConfigId as ParentId that is referenced by other features. This way, you can combine rule conditions and features for flexible configurations.
	//
	// For more information, see [BatchSetCdnDomainConfig](https://help.aliyun.com/document_detail/90915.html) or ParentId configuration example in this topic.
	//
	// example:
	//
	// 222728944812032
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The status of the configuration. Valid values:
	//
	// 	- **success**
	//
	// 	- **testing**
	//
	// 	- **failed**
	//
	// 	- **configuring**
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetConfigId() *string {
	return s.ConfigId
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetFunctionArgs() *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	return s.FunctionArgs
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetFunctionName() *string {
	return s.FunctionName
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetParentId() *string {
	return s.ParentId
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) GetStatus() *string {
	return s.Status
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetConfigId(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ConfigId = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionArgs(v *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionArgs = v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetFunctionName(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.FunctionName = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetParentId(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.ParentId = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) SetStatus(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig {
	s.Status = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfig) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs struct {
	FunctionArg []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg `json:"FunctionArg,omitempty" xml:"FunctionArg,omitempty" type:"Repeated"`
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) GetFunctionArg() []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	return s.FunctionArg
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) SetFunctionArg(v []*DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs {
	s.FunctionArg = v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgs) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg struct {
	// The parameter name, which is the configuration item of **functionName**. You can configure multiple configuration items.
	//
	// example:
	//
	// auth_type
	ArgName *string `json:"ArgName,omitempty" xml:"ArgName,omitempty"`
	// The parameter value, which is the value of the configuration item of **functionName**.
	//
	// example:
	//
	// req_auth
	ArgValue *string `json:"ArgValue,omitempty" xml:"ArgValue,omitempty"`
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GoString() string {
	return s.String()
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GetArgName() *string {
	return s.ArgName
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) GetArgValue() *string {
	return s.ArgValue
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgName(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgName = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) SetArgValue(v string) *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg {
	s.ArgValue = &v
	return s
}

func (s *DescribeCdnDomainConfigsResponseBodyDomainConfigsDomainConfigFunctionArgsFunctionArg) Validate() error {
	return dara.Validate(s)
}
