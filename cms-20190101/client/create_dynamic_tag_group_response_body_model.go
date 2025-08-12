// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDynamicTagGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateDynamicTagGroupResponseBody
	GetCode() *string
	SetId(v string) *CreateDynamicTagGroupResponseBody
	GetId() *string
	SetMessage(v string) *CreateDynamicTagGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDynamicTagGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDynamicTagGroupResponseBody
	GetSuccess() *bool
}

type CreateDynamicTagGroupResponseBody struct {
	// The HTTP status code.
	//
	// >  The status code 200 indicates that the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the tag matching rule.
	//
	// example:
	//
	// 2534dc0a-e3e5-4ae1-a2fc-75ef166c****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 84AC6F0B-7945-466A-AA44-99BB5A561F86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- true: The call is successful.
	//
	// 	- false: The call fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDynamicTagGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDynamicTagGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDynamicTagGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateDynamicTagGroupResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateDynamicTagGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDynamicTagGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDynamicTagGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDynamicTagGroupResponseBody) SetCode(v string) *CreateDynamicTagGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateDynamicTagGroupResponseBody) SetId(v string) *CreateDynamicTagGroupResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDynamicTagGroupResponseBody) SetMessage(v string) *CreateDynamicTagGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDynamicTagGroupResponseBody) SetRequestId(v string) *CreateDynamicTagGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDynamicTagGroupResponseBody) SetSuccess(v bool) *CreateDynamicTagGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDynamicTagGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
