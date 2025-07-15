// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *CreateAclResponseBody
	GetCode() *int32
	SetMessage(v string) *CreateAclResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateAclResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAclResponseBody
	GetSuccess() *bool
}

type CreateAclResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 56729737-C428-4E1B-AC68-7A8C2D5****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAclResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAclResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *CreateAclResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAclResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAclResponseBody) SetCode(v int32) *CreateAclResponseBody {
	s.Code = &v
	return s
}

func (s *CreateAclResponseBody) SetMessage(v string) *CreateAclResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAclResponseBody) SetRequestId(v string) *CreateAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAclResponseBody) SetSuccess(v bool) *CreateAclResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAclResponseBody) Validate() error {
	return dara.Validate(s)
}
