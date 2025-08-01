// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ImportServicesResponseBody
	GetCode() *int32
	SetData(v bool) *ImportServicesResponseBody
	GetData() *bool
	SetHttpStatusCode(v int32) *ImportServicesResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ImportServicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportServicesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ImportServicesResponseBody
	GetSuccess() *bool
}

type ImportServicesResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result returned.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
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
	// The request is successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0F0FBA7D-5AC5-5DC4-A1E9-E9656BFAE1B9
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

func (s ImportServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportServicesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ImportServicesResponseBody) GetData() *bool {
	return s.Data
}

func (s *ImportServicesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ImportServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportServicesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ImportServicesResponseBody) SetCode(v int32) *ImportServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ImportServicesResponseBody) SetData(v bool) *ImportServicesResponseBody {
	s.Data = &v
	return s
}

func (s *ImportServicesResponseBody) SetHttpStatusCode(v int32) *ImportServicesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ImportServicesResponseBody) SetMessage(v string) *ImportServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ImportServicesResponseBody) SetRequestId(v string) *ImportServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportServicesResponseBody) SetSuccess(v bool) *ImportServicesResponseBody {
	s.Success = &v
	return s
}

func (s *ImportServicesResponseBody) Validate() error {
	return dara.Validate(s)
}
