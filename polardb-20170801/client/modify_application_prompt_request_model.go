// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApplicationPromptRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ModifyApplicationPromptRequest
	GetApplicationId() *string
	SetPromptId(v string) *ModifyApplicationPromptRequest
	GetPromptId() *string
	SetPromptName(v string) *ModifyApplicationPromptRequest
	GetPromptName() *string
	SetPromptValue(v string) *ModifyApplicationPromptRequest
	GetPromptValue() *string
}

type ModifyApplicationPromptRequest struct {
	// The application ID. Usage note: Only applications of the mem0 type support prompt operations. You can verify this by checking the ApplicationType field by calling DescribeApplicationAttribute.
	//
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// The prompt ID. The value is an identifier that starts with papt-.
	//
	// This parameter is required.
	//
	// example:
	//
	// papt-f9lajgw765f4fnrzn1
	PromptId *string `json:"PromptId,omitempty" xml:"PromptId,omitempty"`
	// The prompt name.
	//
	// example:
	//
	// prompt name
	PromptName *string `json:"PromptName,omitempty" xml:"PromptName,omitempty"`
	// The prompt content.
	//
	// example:
	//
	// prompt value
	PromptValue *string `json:"PromptValue,omitempty" xml:"PromptValue,omitempty"`
}

func (s ModifyApplicationPromptRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyApplicationPromptRequest) GoString() string {
	return s.String()
}

func (s *ModifyApplicationPromptRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ModifyApplicationPromptRequest) GetPromptId() *string {
	return s.PromptId
}

func (s *ModifyApplicationPromptRequest) GetPromptName() *string {
	return s.PromptName
}

func (s *ModifyApplicationPromptRequest) GetPromptValue() *string {
	return s.PromptValue
}

func (s *ModifyApplicationPromptRequest) SetApplicationId(v string) *ModifyApplicationPromptRequest {
	s.ApplicationId = &v
	return s
}

func (s *ModifyApplicationPromptRequest) SetPromptId(v string) *ModifyApplicationPromptRequest {
	s.PromptId = &v
	return s
}

func (s *ModifyApplicationPromptRequest) SetPromptName(v string) *ModifyApplicationPromptRequest {
	s.PromptName = &v
	return s
}

func (s *ModifyApplicationPromptRequest) SetPromptValue(v string) *ModifyApplicationPromptRequest {
	s.PromptValue = &v
	return s
}

func (s *ModifyApplicationPromptRequest) Validate() error {
	return dara.Validate(s)
}
