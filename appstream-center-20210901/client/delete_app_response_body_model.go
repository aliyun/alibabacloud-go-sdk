// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAppResponseBody
	GetCode() *string
	SetData(v bool) *DeleteAppResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteAppResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAppResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAppResponseBody
	GetSuccess() *bool
}

type DeleteAppResponseBody struct {
	// The error code. This field is returned when the request fails. A successful response may not contain this field.
	//
	// example:
	//
	// InternalError
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the operation was successful. A value of true indicates that the operation was successful.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error or result message. This field describes the cause when a request fails. A successful response may not contain this field.
	//
	// example:
	//
	// Internal error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can provide this value for troubleshooting.
	//
	// example:
	//
	// 01A1A962-5503-107B-AC8F-EBA85DF905AC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates that the request was successful. To check whether an asynchronous task is completed, you must query the task status separately.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAppResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAppResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAppResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteAppResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAppResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAppResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAppResponseBody) SetCode(v string) *DeleteAppResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAppResponseBody) SetData(v bool) *DeleteAppResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAppResponseBody) SetMessage(v string) *DeleteAppResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAppResponseBody) SetRequestId(v string) *DeleteAppResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAppResponseBody) SetSuccess(v bool) *DeleteAppResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAppResponseBody) Validate() error {
	return dara.Validate(s)
}
