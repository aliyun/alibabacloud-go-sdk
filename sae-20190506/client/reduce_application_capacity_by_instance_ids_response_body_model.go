// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReduceApplicationCapacityByInstanceIdsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody
	GetCode() *string
	SetData(v *ReduceApplicationCapacityByInstanceIdsResponseBodyData) *ReduceApplicationCapacityByInstanceIdsResponseBody
	GetData() *ReduceApplicationCapacityByInstanceIdsResponseBodyData
	SetErrorCode(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody
	GetErrorCode() *string
	SetMessage(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ReduceApplicationCapacityByInstanceIdsResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody
	GetTraceId() *string
}

type ReduceApplicationCapacityByInstanceIdsResponseBody struct {
	// The HTTP status code.
	//
	// 	- **2xx**: The call was successful.
	//
	// 	- **3xx**: The call was redirected.
	//
	// 	- **4xx**: The call failed.
	//
	// 	- **5xx**: A server error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the change process.
	Data *ReduceApplicationCapacityByInstanceIdsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code returned if the request failed. Take note of the following rules:
	//
	// 	- The **ErrorCode*	- parameter is not returned if the request succeeds.
	//
	// 	- If the call fails, the **ErrorCode*	- parameter is returned. For more information, see the "**Error codes**" section of this topic.
	//
	// example:
	//
	// Null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information that is returned. Take note of the following rules:
	//
	// 	- success: If the call is successful, **success*	- is returned.
	//
	// 	- An error code: If the call fails, an error code is returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A8E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the information of the change order was queried. Take note of the following rules:
	//
	// 	- **true**: The information was queried.
	//
	// 	- **false**: The image failed to be found.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID that is used to query the details of the request.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s ReduceApplicationCapacityByInstanceIdsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReduceApplicationCapacityByInstanceIdsResponseBody) GoString() string {
	return s.String()
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) GetData() *ReduceApplicationCapacityByInstanceIdsResponseBodyData {
	return s.Data
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetCode(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.Code = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetData(v *ReduceApplicationCapacityByInstanceIdsResponseBodyData) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.Data = v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetErrorCode(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetMessage(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.Message = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetRequestId(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetSuccess(v bool) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.Success = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) SetTraceId(v string) *ReduceApplicationCapacityByInstanceIdsResponseBody {
	s.TraceId = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReduceApplicationCapacityByInstanceIdsResponseBodyData struct {
	// The ID of the change process.
	//
	// example:
	//
	// 76fa5c0-9ebb-4bb4-b383-1f885447****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
}

func (s ReduceApplicationCapacityByInstanceIdsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReduceApplicationCapacityByInstanceIdsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBodyData) SetChangeOrderId(v string) *ReduceApplicationCapacityByInstanceIdsResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *ReduceApplicationCapacityByInstanceIdsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
