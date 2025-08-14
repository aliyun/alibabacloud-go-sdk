// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventVariableTemplateBindRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeEventVariableTemplateBindRequest
	GetLang() *string
	SetInputs(v string) *DescribeEventVariableTemplateBindRequest
	GetInputs() *string
	SetRegId(v string) *DescribeEventVariableTemplateBindRequest
	GetRegId() *string
	SetTemplateCode(v string) *DescribeEventVariableTemplateBindRequest
	GetTemplateCode() *string
	SetType(v string) *DescribeEventVariableTemplateBindRequest
	GetType() *string
}

type DescribeEventVariableTemplateBindRequest struct {
	// Sets the language type for requests and received messages. Default value is **zh**. Values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Input parameters, separated by commas if multiple.
	//
	// This parameter is required.
	//
	// example:
	//
	// age,ip
	Inputs *string `json:"inputs,omitempty" xml:"inputs,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Template code.
	//
	// example:
	//
	// register
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// Type
	//
	// This parameter is required.
	//
	// example:
	//
	// NATIVE
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s DescribeEventVariableTemplateBindRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventVariableTemplateBindRequest) GoString() string {
	return s.String()
}

func (s *DescribeEventVariableTemplateBindRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeEventVariableTemplateBindRequest) GetInputs() *string {
	return s.Inputs
}

func (s *DescribeEventVariableTemplateBindRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeEventVariableTemplateBindRequest) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DescribeEventVariableTemplateBindRequest) GetType() *string {
	return s.Type
}

func (s *DescribeEventVariableTemplateBindRequest) SetLang(v string) *DescribeEventVariableTemplateBindRequest {
	s.Lang = &v
	return s
}

func (s *DescribeEventVariableTemplateBindRequest) SetInputs(v string) *DescribeEventVariableTemplateBindRequest {
	s.Inputs = &v
	return s
}

func (s *DescribeEventVariableTemplateBindRequest) SetRegId(v string) *DescribeEventVariableTemplateBindRequest {
	s.RegId = &v
	return s
}

func (s *DescribeEventVariableTemplateBindRequest) SetTemplateCode(v string) *DescribeEventVariableTemplateBindRequest {
	s.TemplateCode = &v
	return s
}

func (s *DescribeEventVariableTemplateBindRequest) SetType(v string) *DescribeEventVariableTemplateBindRequest {
	s.Type = &v
	return s
}

func (s *DescribeEventVariableTemplateBindRequest) Validate() error {
	return dara.Validate(s)
}
