// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRowPermissionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRowPermissionResponseBody
	GetCode() *string
	SetData(v bool) *DeleteRowPermissionResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *DeleteRowPermissionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteRowPermissionResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteRowPermissionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteRowPermissionResponseBody
	GetSuccess() *bool
}

type DeleteRowPermissionResponseBody struct {
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
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteRowPermissionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRowPermissionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRowPermissionResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRowPermissionResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteRowPermissionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteRowPermissionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteRowPermissionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRowPermissionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteRowPermissionResponseBody) SetCode(v string) *DeleteRowPermissionResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRowPermissionResponseBody) SetData(v bool) *DeleteRowPermissionResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteRowPermissionResponseBody) SetHttpStatusCode(v int32) *DeleteRowPermissionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteRowPermissionResponseBody) SetMessage(v string) *DeleteRowPermissionResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteRowPermissionResponseBody) SetRequestId(v string) *DeleteRowPermissionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRowPermissionResponseBody) SetSuccess(v bool) *DeleteRowPermissionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteRowPermissionResponseBody) Validate() error {
	return dara.Validate(s)
}
