// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteInstancesResponseBody
	GetCode() *string
	SetData(v *DeleteInstancesResponseBodyData) *DeleteInstancesResponseBody
	GetData() *DeleteInstancesResponseBodyData
	SetErrorCode(v string) *DeleteInstancesResponseBody
	GetErrorCode() *string
	SetMessage(v string) *DeleteInstancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteInstancesResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *DeleteInstancesResponseBody
	GetTraceId() *string
}

type DeleteInstancesResponseBody struct {
	// The HTTP status code or the error code. Valid values:
	//
	// 	- **2xx**: The request was successful.
	//
	// 	- **3xx**: The request was redirected.
	//
	// 	- **4xx**: The request failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Responses.
	Data *DeleteInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Error code. Valid values:
	//
	// 	- If the request was successful, **ErrorCode*	- is not returned.
	//
	// 	- If the request failed, **ErrorCode*	- is returned. For more information, see **Error codes*	- in this topic.
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Additional information about the call result. Valid values:
	//
	// 	- If the request is normal, success is returned.
	//
	// 	- If the request is abnormal, the specific exception error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether the instance is deleted successfully. Valid values:
	//
	// 	- **true**: The namespace was deleted.
	//
	// 	- **false**: The namespace failed to be deleted.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the trace. This parameter is used to query the exact call information.
	//
	// example:
	//
	// 0a981dd515966966104121683d****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s DeleteInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteInstancesResponseBody) GetData() *DeleteInstancesResponseBodyData {
	return s.Data
}

func (s *DeleteInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteInstancesResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *DeleteInstancesResponseBody) SetCode(v string) *DeleteInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetData(v *DeleteInstancesResponseBodyData) *DeleteInstancesResponseBody {
	s.Data = v
	return s
}

func (s *DeleteInstancesResponseBody) SetErrorCode(v string) *DeleteInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetMessage(v string) *DeleteInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetRequestId(v string) *DeleteInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetSuccess(v bool) *DeleteInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteInstancesResponseBody) SetTraceId(v string) *DeleteInstancesResponseBody {
	s.TraceId = &v
	return s
}

func (s *DeleteInstancesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteInstancesResponseBodyData struct {
	// The ID of the release order.
	//
	// example:
	//
	// 01db03d3-3ee9-48b3-b3d0-dfce2d88****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s DeleteInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteInstancesResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *DeleteInstancesResponseBodyData) SetChangeOrderId(v string) *DeleteInstancesResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *DeleteInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
