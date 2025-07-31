// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantResourcePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GrantResourcePermissionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GrantResourcePermissionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GrantResourcePermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *GrantResourcePermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GrantResourcePermissionResponseBody
	GetSuccess() *bool
}

type GrantResourcePermissionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GrantResourcePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GrantResourcePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *GrantResourcePermissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *GrantResourcePermissionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GrantResourcePermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GrantResourcePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GrantResourcePermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GrantResourcePermissionResponseBody) SetCode(v string) *GrantResourcePermissionResponseBody {
	s.Code = &v
	return s
}

func (s *GrantResourcePermissionResponseBody) SetHttpStatusCode(v int32) *GrantResourcePermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GrantResourcePermissionResponseBody) SetMessage(v string) *GrantResourcePermissionResponseBody {
	s.Message = &v
	return s
}

func (s *GrantResourcePermissionResponseBody) SetRequestId(v string) *GrantResourcePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GrantResourcePermissionResponseBody) SetSuccess(v bool) *GrantResourcePermissionResponseBody {
	s.Success = &v
	return s
}

func (s *GrantResourcePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
