// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ReleaseInstanceResponseBody
	GetCode() *int32
	SetMessage(v string) *ReleaseInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReleaseInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReleaseInstanceResponseBody
	GetSuccess() *bool
}

type ReleaseInstanceResponseBody struct {
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ABA4A7FD-E10F-45C7-9774-A5236015A***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ReleaseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ReleaseInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReleaseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseInstanceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReleaseInstanceResponseBody) SetCode(v int32) *ReleaseInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetMessage(v string) *ReleaseInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetRequestId(v string) *ReleaseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseInstanceResponseBody) SetSuccess(v bool) *ReleaseInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *ReleaseInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
