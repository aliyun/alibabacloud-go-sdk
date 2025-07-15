// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAllowedIpResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateAllowedIpResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateAllowedIpResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateAllowedIpResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateAllowedIpResponseBody
	GetSuccess() *bool
}

type UpdateAllowedIpResponseBody struct {
	// The HTTP status code that is returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 17D425C2-4EA3-4AB8-928D-E10511ECF***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateAllowedIpResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAllowedIpResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAllowedIpResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateAllowedIpResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateAllowedIpResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAllowedIpResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateAllowedIpResponseBody) SetCode(v int32) *UpdateAllowedIpResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateAllowedIpResponseBody) SetMessage(v string) *UpdateAllowedIpResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateAllowedIpResponseBody) SetRequestId(v string) *UpdateAllowedIpResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAllowedIpResponseBody) SetSuccess(v bool) *UpdateAllowedIpResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateAllowedIpResponseBody) Validate() error {
	return dara.Validate(s)
}
