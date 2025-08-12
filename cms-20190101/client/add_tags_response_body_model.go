// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddTagsResponseBody
	GetCode() *string
	SetMessage(v string) *AddTagsResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddTagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddTagsResponseBody
	GetSuccess() *bool
}

type AddTagsResponseBody struct {
	// The status code.
	//
	// > The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified resource is not found.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DAE4B115-3847-5438-8709-423627F0A3A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddTagsResponseBody) GoString() string {
	return s.String()
}

func (s *AddTagsResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddTagsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddTagsResponseBody) SetCode(v string) *AddTagsResponseBody {
	s.Code = &v
	return s
}

func (s *AddTagsResponseBody) SetMessage(v string) *AddTagsResponseBody {
	s.Message = &v
	return s
}

func (s *AddTagsResponseBody) SetRequestId(v string) *AddTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddTagsResponseBody) SetSuccess(v bool) *AddTagsResponseBody {
	s.Success = &v
	return s
}

func (s *AddTagsResponseBody) Validate() error {
	return dara.Validate(s)
}
