// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticDataSourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteAgenticDataSourceResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteAgenticDataSourceResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteAgenticDataSourceResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteAgenticDataSourceResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAgenticDataSourceResponseBody
	GetSuccess() *bool
}

type DeleteAgenticDataSourceResponseBody struct {
	// Indicates whether the deletion was successful.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code returned when the call fails. This field is empty when the call succeeds.
	//
	// example:
	//
	// InvalidParameter
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails. This field is empty when the call succeeds.
	//
	// example:
	//
	// The parameter is invalid
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request. You can use this ID to trace the call.
	//
	// example:
	//
	// request-1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAgenticDataSourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticDataSourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAgenticDataSourceResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteAgenticDataSourceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteAgenticDataSourceResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteAgenticDataSourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAgenticDataSourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAgenticDataSourceResponseBody) SetData(v bool) *DeleteAgenticDataSourceResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAgenticDataSourceResponseBody) SetErrorCode(v string) *DeleteAgenticDataSourceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteAgenticDataSourceResponseBody) SetErrorMessage(v string) *DeleteAgenticDataSourceResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteAgenticDataSourceResponseBody) SetRequestId(v string) *DeleteAgenticDataSourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAgenticDataSourceResponseBody) SetSuccess(v bool) *DeleteAgenticDataSourceResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAgenticDataSourceResponseBody) Validate() error {
	return dara.Validate(s)
}
