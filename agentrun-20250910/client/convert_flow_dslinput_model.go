// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertFlowDSLInput interface {
	dara.Model
	String() string
	GoString() string
	SetDslSource(v *ConvertFlowDSLInputDslSource) *ConvertFlowDSLInput
	GetDslSource() *ConvertFlowDSLInputDslSource
	SetOptions(v *ConvertFlowDSLInputOptions) *ConvertFlowDSLInput
	GetOptions() *ConvertFlowDSLInputOptions
}

type ConvertFlowDSLInput struct {
	// 工作流DSL的来源配置，支持inline和base64两种格式
	//
	// This parameter is required.
	DslSource *ConvertFlowDSLInputDslSource `json:"dslSource,omitempty" xml:"dslSource,omitempty" type:"Struct"`
	// DSL转换的可选配置参数
	Options *ConvertFlowDSLInputOptions `json:"options,omitempty" xml:"options,omitempty" type:"Struct"`
}

func (s ConvertFlowDSLInput) String() string {
	return dara.Prettify(s)
}

func (s ConvertFlowDSLInput) GoString() string {
	return s.String()
}

func (s *ConvertFlowDSLInput) GetDslSource() *ConvertFlowDSLInputDslSource {
	return s.DslSource
}

func (s *ConvertFlowDSLInput) GetOptions() *ConvertFlowDSLInputOptions {
	return s.Options
}

func (s *ConvertFlowDSLInput) SetDslSource(v *ConvertFlowDSLInputDslSource) *ConvertFlowDSLInput {
	s.DslSource = v
	return s
}

func (s *ConvertFlowDSLInput) SetOptions(v *ConvertFlowDSLInputOptions) *ConvertFlowDSLInput {
	s.Options = v
	return s
}

func (s *ConvertFlowDSLInput) Validate() error {
	if s.DslSource != nil {
		if err := s.DslSource.Validate(); err != nil {
			return err
		}
	}
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConvertFlowDSLInputDslSource struct {
	// DSL内容，可以是原始JSON字符串，或根据encoding字段指定的编码格式
	//
	// This parameter is required.
	//
	// example:
	//
	// {"app":{"name":"My Flow"}}
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// DSL内容的编码方式。不填表示内容为原始字符串；base64表示内容经过Base64编码；base64+gzip表示内容经过gzip压缩后再Base64编码
	//
	// example:
	//
	// base64
	Encoding *string `json:"encoding,omitempty" xml:"encoding,omitempty"`
	// 源DSL的提供商类型，如：dify、n8n、zapier等
	//
	// This parameter is required.
	//
	// example:
	//
	// dify
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
}

func (s ConvertFlowDSLInputDslSource) String() string {
	return dara.Prettify(s)
}

func (s ConvertFlowDSLInputDslSource) GoString() string {
	return s.String()
}

func (s *ConvertFlowDSLInputDslSource) GetContent() *string {
	return s.Content
}

func (s *ConvertFlowDSLInputDslSource) GetEncoding() *string {
	return s.Encoding
}

func (s *ConvertFlowDSLInputDslSource) GetProvider() *string {
	return s.Provider
}

func (s *ConvertFlowDSLInputDslSource) SetContent(v string) *ConvertFlowDSLInputDslSource {
	s.Content = &v
	return s
}

func (s *ConvertFlowDSLInputDslSource) SetEncoding(v string) *ConvertFlowDSLInputDslSource {
	s.Encoding = &v
	return s
}

func (s *ConvertFlowDSLInputDslSource) SetProvider(v string) *ConvertFlowDSLInputDslSource {
	s.Provider = &v
	return s
}

func (s *ConvertFlowDSLInputDslSource) Validate() error {
	return dara.Validate(s)
}

type ConvertFlowDSLInputOptions struct {
	// 是否执行兼容性检查，默认为true
	//
	// example:
	//
	// true
	CompatibilityCheck *bool   `json:"compatibilityCheck,omitempty" xml:"compatibilityCheck,omitempty"`
	CredentialName     *string `json:"credentialName,omitempty" xml:"credentialName,omitempty"`
	FlowName           *string `json:"flowName,omitempty" xml:"flowName,omitempty"`
	VpcEndpointName    *string `json:"vpcEndpointName,omitempty" xml:"vpcEndpointName,omitempty"`
}

func (s ConvertFlowDSLInputOptions) String() string {
	return dara.Prettify(s)
}

func (s ConvertFlowDSLInputOptions) GoString() string {
	return s.String()
}

func (s *ConvertFlowDSLInputOptions) GetCompatibilityCheck() *bool {
	return s.CompatibilityCheck
}

func (s *ConvertFlowDSLInputOptions) GetCredentialName() *string {
	return s.CredentialName
}

func (s *ConvertFlowDSLInputOptions) GetFlowName() *string {
	return s.FlowName
}

func (s *ConvertFlowDSLInputOptions) GetVpcEndpointName() *string {
	return s.VpcEndpointName
}

func (s *ConvertFlowDSLInputOptions) SetCompatibilityCheck(v bool) *ConvertFlowDSLInputOptions {
	s.CompatibilityCheck = &v
	return s
}

func (s *ConvertFlowDSLInputOptions) SetCredentialName(v string) *ConvertFlowDSLInputOptions {
	s.CredentialName = &v
	return s
}

func (s *ConvertFlowDSLInputOptions) SetFlowName(v string) *ConvertFlowDSLInputOptions {
	s.FlowName = &v
	return s
}

func (s *ConvertFlowDSLInputOptions) SetVpcEndpointName(v string) *ConvertFlowDSLInputOptions {
	s.VpcEndpointName = &v
	return s
}

func (s *ConvertFlowDSLInputOptions) Validate() error {
	return dara.Validate(s)
}
