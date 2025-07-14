// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRollbackApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RollbackApplicationResponseBody
	GetCode() *string
	SetData(v *RollbackApplicationResponseBodyData) *RollbackApplicationResponseBody
	GetData() *RollbackApplicationResponseBodyData
	SetErrorCode(v string) *RollbackApplicationResponseBody
	GetErrorCode() *string
	SetMessage(v string) *RollbackApplicationResponseBody
	GetMessage() *string
	SetRequestId(v string) *RollbackApplicationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RollbackApplicationResponseBody
	GetSuccess() *bool
	SetTraceId(v string) *RollbackApplicationResponseBody
	GetTraceId() *string
}

type RollbackApplicationResponseBody struct {
	// The HTTP status code. Take note of the following rules:
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
	// The response.
	Data *RollbackApplicationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// The message returned for the operation.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 91F93257-7A4A-4BD3-9A7E-2F6EAE6D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the application is successfully rolled back. Take note of the following rules:
	//
	// 	- **true**
	//
	// 	- **false**
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

func (s RollbackApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RollbackApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *RollbackApplicationResponseBody) GetCode() *string {
	return s.Code
}

func (s *RollbackApplicationResponseBody) GetData() *RollbackApplicationResponseBodyData {
	return s.Data
}

func (s *RollbackApplicationResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *RollbackApplicationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RollbackApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RollbackApplicationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RollbackApplicationResponseBody) GetTraceId() *string {
	return s.TraceId
}

func (s *RollbackApplicationResponseBody) SetCode(v string) *RollbackApplicationResponseBody {
	s.Code = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetData(v *RollbackApplicationResponseBodyData) *RollbackApplicationResponseBody {
	s.Data = v
	return s
}

func (s *RollbackApplicationResponseBody) SetErrorCode(v string) *RollbackApplicationResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetMessage(v string) *RollbackApplicationResponseBody {
	s.Message = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetRequestId(v string) *RollbackApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetSuccess(v bool) *RollbackApplicationResponseBody {
	s.Success = &v
	return s
}

func (s *RollbackApplicationResponseBody) SetTraceId(v string) *RollbackApplicationResponseBody {
	s.TraceId = &v
	return s
}

func (s *RollbackApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type RollbackApplicationResponseBodyData struct {
	// The ID of the change process.
	//
	// example:
	//
	// 01db03d3-3ee9-48b3-b3d0-dfce2d88****
	ChangeOrderId *string `json:"ChangeOrderId,omitempty" xml:"ChangeOrderId,omitempty"`
	// Specifies whether approval is required when a RAM user performs release. Take note of the following rules:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsNeedApproval *bool `json:"IsNeedApproval,omitempty" xml:"IsNeedApproval,omitempty"`
}

func (s RollbackApplicationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RollbackApplicationResponseBodyData) GoString() string {
	return s.String()
}

func (s *RollbackApplicationResponseBodyData) GetChangeOrderId() *string {
	return s.ChangeOrderId
}

func (s *RollbackApplicationResponseBodyData) GetIsNeedApproval() *bool {
	return s.IsNeedApproval
}

func (s *RollbackApplicationResponseBodyData) SetChangeOrderId(v string) *RollbackApplicationResponseBodyData {
	s.ChangeOrderId = &v
	return s
}

func (s *RollbackApplicationResponseBodyData) SetIsNeedApproval(v bool) *RollbackApplicationResponseBodyData {
	s.IsNeedApproval = &v
	return s
}

func (s *RollbackApplicationResponseBodyData) Validate() error {
	return dara.Validate(s)
}
