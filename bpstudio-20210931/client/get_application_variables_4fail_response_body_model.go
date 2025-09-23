// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationVariables4FailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetApplicationVariables4FailResponseBody
	GetCode() *string
	SetData(v []*GetApplicationVariables4FailResponseBodyData) *GetApplicationVariables4FailResponseBody
	GetData() []*GetApplicationVariables4FailResponseBodyData
	SetMessage(v string) *GetApplicationVariables4FailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetApplicationVariables4FailResponseBody
	GetRequestId() *string
}

type GetApplicationVariables4FailResponseBody struct {
	// example:
	//
	// 200
	Code *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetApplicationVariables4FailResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// BFB7F5C8-FE7A-06CA-9F87-ABBF6B848F0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetApplicationVariables4FailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationVariables4FailResponseBody) GoString() string {
	return s.String()
}

func (s *GetApplicationVariables4FailResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetApplicationVariables4FailResponseBody) GetData() []*GetApplicationVariables4FailResponseBodyData {
	return s.Data
}

func (s *GetApplicationVariables4FailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetApplicationVariables4FailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationVariables4FailResponseBody) SetCode(v string) *GetApplicationVariables4FailResponseBody {
	s.Code = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBody) SetData(v []*GetApplicationVariables4FailResponseBodyData) *GetApplicationVariables4FailResponseBody {
	s.Data = v
	return s
}

func (s *GetApplicationVariables4FailResponseBody) SetMessage(v string) *GetApplicationVariables4FailResponseBody {
	s.Message = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBody) SetRequestId(v string) *GetApplicationVariables4FailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetApplicationVariables4FailResponseBodyData struct {
	// example:
	//
	// instance_name
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// example:
	//
	// cadt-app-01
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// ${name}
	PlaceHolder *string `json:"PlaceHolder,omitempty" xml:"PlaceHolder,omitempty"`
	// example:
	//
	// cn-shenzhen
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// cadt-app-01
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// ${name}
	Variable *string `json:"Variable,omitempty" xml:"Variable,omitempty"`
}

func (s GetApplicationVariables4FailResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationVariables4FailResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetApplicationVariables4FailResponseBodyData) GetAttribute() *string {
	return s.Attribute
}

func (s *GetApplicationVariables4FailResponseBodyData) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *GetApplicationVariables4FailResponseBodyData) GetPlaceHolder() *string {
	return s.PlaceHolder
}

func (s *GetApplicationVariables4FailResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GetApplicationVariables4FailResponseBodyData) GetValue() *string {
	return s.Value
}

func (s *GetApplicationVariables4FailResponseBodyData) GetVariable() *string {
	return s.Variable
}

func (s *GetApplicationVariables4FailResponseBodyData) SetAttribute(v string) *GetApplicationVariables4FailResponseBodyData {
	s.Attribute = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBodyData) SetDefaultValue(v string) *GetApplicationVariables4FailResponseBodyData {
	s.DefaultValue = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBodyData) SetPlaceHolder(v string) *GetApplicationVariables4FailResponseBodyData {
	s.PlaceHolder = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBodyData) SetRegion(v string) *GetApplicationVariables4FailResponseBodyData {
	s.Region = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBodyData) SetValue(v string) *GetApplicationVariables4FailResponseBodyData {
	s.Value = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBodyData) SetVariable(v string) *GetApplicationVariables4FailResponseBodyData {
	s.Variable = &v
	return s
}

func (s *GetApplicationVariables4FailResponseBodyData) Validate() error {
	return dara.Validate(s)
}
