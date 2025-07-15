// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAclResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteAclResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteAclResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAclResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAclResponseBody
	GetSuccess() *bool
}

type DeleteAclResponseBody struct {
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
	// B0740227-AA9A-4E14-8E9F-36ED665****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAclResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAclResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAclResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteAclResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAclResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAclResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAclResponseBody) SetCode(v int32) *DeleteAclResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAclResponseBody) SetMessage(v string) *DeleteAclResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAclResponseBody) SetRequestId(v string) *DeleteAclResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAclResponseBody) SetSuccess(v bool) *DeleteAclResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAclResponseBody) Validate() error {
	return dara.Validate(s)
}
