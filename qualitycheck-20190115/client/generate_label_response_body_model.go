// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateLabelResponseBody
	GetCode() *string
	SetData(v string) *GenerateLabelResponseBody
	GetData() *string
	SetMessage(v string) *GenerateLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *GenerateLabelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GenerateLabelResponseBody
	GetSuccess() *bool
}

type GenerateLabelResponseBody struct {
	// The response code. A value of **200*	- indicates success. Other values indicate failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the generation task.
	//
	// example:
	//
	// 20260616-4955F615-A74E-171E-86ED-080F60C72EC9
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned when an error occurs.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 3CEA0495-341B-4482-9AD9-8191EF4***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values: true: The call was successful. false: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GenerateLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateLabelResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateLabelResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateLabelResponseBody) GetData() *string {
	return s.Data
}

func (s *GenerateLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateLabelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GenerateLabelResponseBody) SetCode(v string) *GenerateLabelResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateLabelResponseBody) SetData(v string) *GenerateLabelResponseBody {
	s.Data = &v
	return s
}

func (s *GenerateLabelResponseBody) SetMessage(v string) *GenerateLabelResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateLabelResponseBody) SetRequestId(v string) *GenerateLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateLabelResponseBody) SetSuccess(v bool) *GenerateLabelResponseBody {
	s.Success = &v
	return s
}

func (s *GenerateLabelResponseBody) Validate() error {
	return dara.Validate(s)
}
