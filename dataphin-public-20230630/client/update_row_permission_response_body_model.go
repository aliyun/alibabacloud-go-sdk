// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRowPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateRowPermissionResponseBody
	GetCode() *string
	SetData(v bool) *UpdateRowPermissionResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *UpdateRowPermissionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateRowPermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateRowPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateRowPermissionResponseBody
	GetSuccess() *bool
}

type UpdateRowPermissionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s UpdateRowPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRowPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRowPermissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateRowPermissionResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateRowPermissionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateRowPermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateRowPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRowPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateRowPermissionResponseBody) SetCode(v string) *UpdateRowPermissionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateRowPermissionResponseBody) SetData(v bool) *UpdateRowPermissionResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateRowPermissionResponseBody) SetHttpStatusCode(v int32) *UpdateRowPermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateRowPermissionResponseBody) SetMessage(v string) *UpdateRowPermissionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateRowPermissionResponseBody) SetRequestId(v string) *UpdateRowPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRowPermissionResponseBody) SetSuccess(v bool) *UpdateRowPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateRowPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
