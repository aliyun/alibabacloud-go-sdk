// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateSwimmingLaneGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody
	GetCode() *string
	SetData(v *CreateOrUpdateSwimmingLaneGroupResponseBodyData) *CreateOrUpdateSwimmingLaneGroupResponseBody
	GetData() *CreateOrUpdateSwimmingLaneGroupResponseBodyData
	SetErrorCode(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody
	GetErrorCode() *string
	SetMessage(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateOrUpdateSwimmingLaneGroupResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody
	GetTraceId() *string
}

type CreateOrUpdateSwimmingLaneGroupResponseBody struct {
	// The API status code or POP error code. Valid values:
	//
	// - **2xx**: Success.
	//
	// - **3xx**: Redirection.
	//
	// - **4xx**: Request error.
	//
	// - **5xx**: Server error.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The application information.
	Data *CreateOrUpdateSwimmingLaneGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code. Valid values:
	//
	// - If the request is successful, the **ErrorCode*	- parameter is not returned.
	//
	// - If the request fails, the **ErrorCode*	- parameter is returned. For more information, see the **Error codes*	- section in this topic.
	//
	// example:
	//
	// InvalidParameter.NotEmpty
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The additional information. Valid values:
	//
	// - If the request is successful, **success*	- is returned.
	//
	// - If the request fails, a specific error code is returned.
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
	// Indicates whether the traffic lane group is created. Valid values:
	//
	// - **true**: The traffic lane group is created.
	//
	// - **false**: The traffic lane group fails to be created.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The trace ID used to query invocation details.
	//
	// example:
	//
	// 0a98a02315955564772843261e****
	TraceId *string `json:"TraceId,omitempty" xml:"TraceId,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) GetData() *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	return s.Data
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) SetCode(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) SetData(v *CreateOrUpdateSwimmingLaneGroupResponseBodyData) *CreateOrUpdateSwimmingLaneGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) SetErrorCode(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) SetMessage(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) SetRequestId(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) SetSuccess(v bool) *CreateOrUpdateSwimmingLaneGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) SetTraceId(v string) *CreateOrUpdateSwimmingLaneGroupResponseBody {
	s.TraceId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateOrUpdateSwimmingLaneGroupResponseBodyData struct {
	// The ID of the traffic lane group.
	//
	// example:
	//
	// 110272
	GroupId *int64 `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
}

func (s CreateOrUpdateSwimmingLaneGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateSwimmingLaneGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) GetGroupId() *int64 {
	return s.GroupId
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) SetGroupId(v int64) *CreateOrUpdateSwimmingLaneGroupResponseBodyData {
	s.GroupId = &v
	return s
}

func (s *CreateOrUpdateSwimmingLaneGroupResponseBodyData) Validate() error {
	return dara.Validate(s)
}
