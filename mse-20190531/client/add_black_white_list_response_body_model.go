// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBlackWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddBlackWhiteListResponseBody
	GetCode() *int32
	SetData(v int64) *AddBlackWhiteListResponseBody
	GetData() *int64
	SetHttpStatusCode(v int32) *AddBlackWhiteListResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddBlackWhiteListResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddBlackWhiteListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddBlackWhiteListResponseBody
	GetSuccess() *bool
}

type AddBlackWhiteListResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the record.
	//
	// example:
	//
	// 2
	Data *int64 `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// You are not authorized to perform this operation.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 966F6CA7-16D0-50AB-AB02-E140934F90C1
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

func (s AddBlackWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddBlackWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *AddBlackWhiteListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddBlackWhiteListResponseBody) GetData() *int64 {
	return s.Data
}

func (s *AddBlackWhiteListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddBlackWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddBlackWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddBlackWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddBlackWhiteListResponseBody) SetCode(v int32) *AddBlackWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *AddBlackWhiteListResponseBody) SetData(v int64) *AddBlackWhiteListResponseBody {
	s.Data = &v
	return s
}

func (s *AddBlackWhiteListResponseBody) SetHttpStatusCode(v int32) *AddBlackWhiteListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddBlackWhiteListResponseBody) SetMessage(v string) *AddBlackWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *AddBlackWhiteListResponseBody) SetRequestId(v string) *AddBlackWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddBlackWhiteListResponseBody) SetSuccess(v bool) *AddBlackWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *AddBlackWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
