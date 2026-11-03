// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateOrUpdateSwimmingLaneResponseBody
	GetCode() *string
	SetData(v *CreateOrUpdateSwimmingLaneResponseBodyData) *CreateOrUpdateSwimmingLaneResponseBody
	GetData() *CreateOrUpdateSwimmingLaneResponseBodyData
	SetErrorCode(v string) *CreateOrUpdateSwimmingLaneResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateOrUpdateSwimmingLaneResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateOrUpdateSwimmingLaneResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOrUpdateSwimmingLaneResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateOrUpdateSwimmingLaneResponseBody
	GetTraceId() *string
}

type CreateOrUpdateSwimmingLaneResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Client error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned information.
	Data *CreateOrUpdateSwimmingLaneResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Description:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidParameter.NotEmpty
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Description:
	//
	// - If the request is normal, success is returned.
	//
	// - If the request is abnormal, a specific error code is returned.
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
	// Specifies whether the traffic lane is created or updated successfully. Valid values:
	//
	// - true: The traffic lane is created or updated successfully.
	//
	// - false: The traffic lane fails to be created or updated.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query call information accurately.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetData() *CreateOrUpdateSwimmingLaneResponseBodyData {
	return s.Data
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetCode(v string) *CreateOrUpdateSwimmingLaneResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetData(v *CreateOrUpdateSwimmingLaneResponseBodyData) *CreateOrUpdateSwimmingLaneResponseBody {
	s.Data = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetErrorCode(v string) *CreateOrUpdateSwimmingLaneResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetMessage(v string) *CreateOrUpdateSwimmingLaneResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetRequestId(v string) *CreateOrUpdateSwimmingLaneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetSuccess(v bool) *CreateOrUpdateSwimmingLaneResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) SetTraceId(v string) *CreateOrUpdateSwimmingLaneResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrUpdateSwimmingLaneResponseBodyData struct {
	// The ID of the traffic lane.
	//
	// example:
	//
	// 22318
	LaneId *int64 `json:"LaneId,omitempty" xml:"LaneId,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) GetLaneId() *int64 {
	return s.LaneId
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) SetLaneId(v int64) *CreateOrUpdateSwimmingLaneResponseBodyData {
	s.LaneId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneResponseBodyData) Validate() error {
	return dara.Validate(s)
}
