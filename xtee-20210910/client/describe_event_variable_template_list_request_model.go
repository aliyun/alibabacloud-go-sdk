// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventVariableTemplateListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventVariableTemplateListRequest
	GetLang() *string
	SetInputs(v string) *DescribeEventVariableTemplateListRequest
	GetInputs() *string
	SetRegId(v string) *DescribeEventVariableTemplateListRequest
	GetRegId() *string
	SetTemplateCode(v string) *DescribeEventVariableTemplateListRequest
	GetTemplateCode() *string
	SetType(v string) *DescribeEventVariableTemplateListRequest
	GetType() *string
}

type DescribeEventVariableTemplateListRequest struct {
	// The language of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The input parameters. Separate multiple values with commas (,).
	//
	// example:
	//
	// age,ip
	Inputs *string `json:"inputs,omitempty" xml:"inputs,omitempty"`
	// The region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// The template code.
	//
	// example:
	//
	// register
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// The templatetype.
	//
	// example:
	//
	// NATIVE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeEventVariableTemplateListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableTemplateListRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableTemplateListRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventVariableTemplateListRequest) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableTemplateListRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventVariableTemplateListRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DescribeEventVariableTemplateListRequest) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableTemplateListRequest) SetLang(v string) *DescribeEventVariableTemplateListRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventVariableTemplateListRequest) SetInputs(v string) *DescribeEventVariableTemplateListRequest {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableTemplateListRequest) SetRegId(v string) *DescribeEventVariableTemplateListRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventVariableTemplateListRequest) SetTemplateCode(v string) *DescribeEventVariableTemplateListRequest {
	s.TemplateCode = &v
	return s
}

func (s *DescribeEventVariableTemplateListRequest) SetType(v string) *DescribeEventVariableTemplateListRequest {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableTemplateListRequest) Validate() error {
	return dara.Validate(s)
}
