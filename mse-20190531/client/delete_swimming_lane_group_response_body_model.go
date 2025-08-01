// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimmingLaneGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v interface{}) *DeleteSwimmingLaneGroupResponseBody
	GetData() interface{}
	SetErrorCode(v string) *DeleteSwimmingLaneGroupResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteSwimmingLaneGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteSwimmingLaneGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSwimmingLaneGroupResponseBody
	GetSuccess() *bool
}

type DeleteSwimmingLaneGroupResponseBody struct {
	// The details of the data.
	//
	// example:
	//
	// {}
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error code.
	//
	// example:
	//
	// mse-100-000
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
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
	// 52BA6DA6-A702-4362-A32F-DFF79655****
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

func (s DeleteSwimmingLaneGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimmingLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSwimmingLaneGroupResponseBody) GetData() interface{} {
	return s.Data
}

func (s *DeleteSwimmingLaneGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteSwimmingLaneGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteSwimmingLaneGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSwimmingLaneGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSwimmingLaneGroupResponseBody) SetData(v interface{}) *DeleteSwimmingLaneGroupResponseBody {
	s.Data = v
	return s
}

func (s *DeleteSwimmingLaneGroupResponseBody) SetErrorCode(v string) *DeleteSwimmingLaneGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSwimmingLaneGroupResponseBody) SetMessage(v string) *DeleteSwimmingLaneGroupResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteSwimmingLaneGroupResponseBody) SetRequestId(v string) *DeleteSwimmingLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSwimmingLaneGroupResponseBody) SetSuccess(v bool) *DeleteSwimmingLaneGroupResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSwimmingLaneGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
