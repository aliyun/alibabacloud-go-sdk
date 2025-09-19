// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitCheckResponseBodyData) *SubmitCheckResponseBody
	GetData() *SubmitCheckResponseBodyData
	SetRequestId(v string) *SubmitCheckResponseBody
	GetRequestId() *string
	SetTaskId(v string) *SubmitCheckResponseBody
	GetTaskId() *string
}

type SubmitCheckResponseBody struct {
	// The data returned if the call is successful.
	Data *SubmitCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// CE8369A6-A843-5E1B-A613-78E6920D****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the configuration assessment task.
	//
	// example:
	//
	// 5fe3f65d-4012-455d-8232-2a98a858****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitCheckResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitCheckResponseBody) GetData() *SubmitCheckResponseBodyData {
	return s.Data
}

func (s *SubmitCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitCheckResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitCheckResponseBody) SetData(v *SubmitCheckResponseBodyData) *SubmitCheckResponseBody {
	s.Data = v
	return s
}

func (s *SubmitCheckResponseBody) SetRequestId(v string) *SubmitCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitCheckResponseBody) SetTaskId(v string) *SubmitCheckResponseBody {
	s.TaskId = &v
	return s
}

func (s *SubmitCheckResponseBody) Validate() error {
	return dara.Validate(s)
}

type SubmitCheckResponseBodyData struct {
	// The operation code of the configuration assessment task.
	//
	// 	- **Throttling**: frequency limit
	//
	// 	- **AuthorizationExhaust**: insufficient quota
	//
	// example:
	//
	// Throttling
	OperateCode *string `json:"OperateCode,omitempty" xml:"OperateCode,omitempty"`
	// The throttling duration. Unit: seconds.
	//
	// example:
	//
	// 1800
	ThrottlingTimeSecond *int32 `json:"ThrottlingTimeSecond,omitempty" xml:"ThrottlingTimeSecond,omitempty"`
}

func (s SubmitCheckResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitCheckResponseBodyData) GetOperateCode() *string {
	return s.OperateCode
}

func (s *SubmitCheckResponseBodyData) GetThrottlingTimeSecond() *int32 {
	return s.ThrottlingTimeSecond
}

func (s *SubmitCheckResponseBodyData) SetOperateCode(v string) *SubmitCheckResponseBodyData {
	s.OperateCode = &v
	return s
}

func (s *SubmitCheckResponseBodyData) SetThrottlingTimeSecond(v int32) *SubmitCheckResponseBodyData {
	s.ThrottlingTimeSecond = &v
	return s
}

func (s *SubmitCheckResponseBodyData) Validate() error {
	return dara.Validate(s)
}
