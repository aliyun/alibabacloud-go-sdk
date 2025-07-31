// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRowPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateRowPermissionResponseBody
	GetCode() *string
	SetData(v int64) *CreateRowPermissionResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *CreateRowPermissionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *CreateRowPermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateRowPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateRowPermissionResponseBody
	GetSuccess() *bool
}

type CreateRowPermissionResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
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

func (s CreateRowPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRowPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRowPermissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateRowPermissionResponseBody) GetData() *int64 {
	return s.Data
}

func (s *CreateRowPermissionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *CreateRowPermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateRowPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRowPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateRowPermissionResponseBody) SetCode(v string) *CreateRowPermissionResponseBody {
	s.Code = &v
	return s
}

func (s *CreateRowPermissionResponseBody) SetData(v int64) *CreateRowPermissionResponseBody {
	s.Data = &v
	return s
}

func (s *CreateRowPermissionResponseBody) SetHttpStatusCode(v int32) *CreateRowPermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateRowPermissionResponseBody) SetMessage(v string) *CreateRowPermissionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateRowPermissionResponseBody) SetRequestId(v string) *CreateRowPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRowPermissionResponseBody) SetSuccess(v bool) *CreateRowPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateRowPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
