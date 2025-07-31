// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeResourcePermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RevokeResourcePermissionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RevokeResourcePermissionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RevokeResourcePermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *RevokeResourcePermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RevokeResourcePermissionResponseBody
	GetSuccess() *bool
}

type RevokeResourcePermissionResponseBody struct {
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

func (s RevokeResourcePermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevokeResourcePermissionResponseBody) GoString() string {
	return s.String()
}

func (s *RevokeResourcePermissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *RevokeResourcePermissionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RevokeResourcePermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RevokeResourcePermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevokeResourcePermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RevokeResourcePermissionResponseBody) SetCode(v string) *RevokeResourcePermissionResponseBody {
	s.Code = &v
	return s
}

func (s *RevokeResourcePermissionResponseBody) SetHttpStatusCode(v int32) *RevokeResourcePermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RevokeResourcePermissionResponseBody) SetMessage(v string) *RevokeResourcePermissionResponseBody {
	s.Message = &v
	return s
}

func (s *RevokeResourcePermissionResponseBody) SetRequestId(v string) *RevokeResourcePermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevokeResourcePermissionResponseBody) SetSuccess(v bool) *RevokeResourcePermissionResponseBody {
	s.Success = &v
	return s
}

func (s *RevokeResourcePermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
