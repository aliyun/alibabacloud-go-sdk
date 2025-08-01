// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBlackWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateBlackWhiteListResponseBody
	GetCode() *int32
	SetData(v int64) *UpdateBlackWhiteListResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *UpdateBlackWhiteListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateBlackWhiteListResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateBlackWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBlackWhiteListResponseBody
	GetSuccess() *bool
}

type UpdateBlackWhiteListResponseBody struct {
	// The return value.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the record.
	//
	// example:
	//
	// 13
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AD5DEDA0-C82A-50D9-AF54-BD3576CCFB4C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- `true`: The request was successful.
	//
	// 	- `false`: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBlackWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBlackWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBlackWhiteListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateBlackWhiteListResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateBlackWhiteListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateBlackWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateBlackWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBlackWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBlackWhiteListResponseBody) SetCode(v int32) *UpdateBlackWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateBlackWhiteListResponseBody) SetData(v int64) *UpdateBlackWhiteListResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateBlackWhiteListResponseBody) SetHttpStatusCode(v int32) *UpdateBlackWhiteListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBlackWhiteListResponseBody) SetMessage(v string) *UpdateBlackWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateBlackWhiteListResponseBody) SetRequestId(v string) *UpdateBlackWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBlackWhiteListResponseBody) SetSuccess(v bool) *UpdateBlackWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBlackWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
