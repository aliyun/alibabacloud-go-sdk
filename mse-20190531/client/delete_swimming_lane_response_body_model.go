// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimmingLaneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DeleteSwimmingLaneResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteSwimmingLaneResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSwimmingLaneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSwimmingLaneResponseBody
	GetSuccess() *bool
}

type DeleteSwimmingLaneResponseBody struct {
	// The error code.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// The request was successfully processed.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 69AD2AA7-DB47-449B-941B-B14409DF****
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

func (s DeleteSwimmingLaneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimmingLaneResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSwimmingLaneResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteSwimmingLaneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSwimmingLaneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSwimmingLaneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSwimmingLaneResponseBody) SetErrorCode(v string) *DeleteSwimmingLaneResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSwimmingLaneResponseBody) SetMessage(v string) *DeleteSwimmingLaneResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSwimmingLaneResponseBody) SetRequestId(v string) *DeleteSwimmingLaneResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSwimmingLaneResponseBody) SetSuccess(v bool) *DeleteSwimmingLaneResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSwimmingLaneResponseBody) Validate() error {
	return dara.Validate(s)
}
