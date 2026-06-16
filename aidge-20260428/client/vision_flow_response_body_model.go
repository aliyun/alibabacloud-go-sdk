// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVisionFlowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VisionFlowResponseBody
	GetCode() *string
	SetData(v *VisionFlowResponseBodyData) *VisionFlowResponseBody
	GetData() *VisionFlowResponseBodyData
	SetMessage(v string) *VisionFlowResponseBody
	GetMessage() *string
	SetRequestId(v string) *VisionFlowResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VisionFlowResponseBody
	GetSuccess() *bool
}

type VisionFlowResponseBody struct {
	// The status code. The value "success" is returned for a successful call.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The submit status data of the asynchronous task, which contains the asynchronous task ID.
	Data *VisionFlowResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. The value "Success" is returned for a successful call.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID, which uniquely identifies the API call.
	//
	// example:
	//
	// F3E65785-0180-1227-91B0-2F5F52F679FB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VisionFlowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VisionFlowResponseBody) GoString() string {
	return s.String()
}

func (s *VisionFlowResponseBody) GetCode() *string {
	return s.Code
}

func (s *VisionFlowResponseBody) GetData() *VisionFlowResponseBodyData {
	return s.Data
}

func (s *VisionFlowResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VisionFlowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VisionFlowResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VisionFlowResponseBody) SetCode(v string) *VisionFlowResponseBody {
	s.Code = &v
	return s
}

func (s *VisionFlowResponseBody) SetData(v *VisionFlowResponseBodyData) *VisionFlowResponseBody {
	s.Data = v
	return s
}

func (s *VisionFlowResponseBody) SetMessage(v string) *VisionFlowResponseBody {
	s.Message = &v
	return s
}

func (s *VisionFlowResponseBody) SetRequestId(v string) *VisionFlowResponseBody {
	s.RequestId = &v
	return s
}

func (s *VisionFlowResponseBody) SetSuccess(v bool) *VisionFlowResponseBody {
	s.Success = &v
	return s
}

func (s *VisionFlowResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VisionFlowResponseBodyData struct {
	// The asynchronous task ID. Use this ID to query the processing result through QueryAsyncTaskResult.
	//
	// example:
	//
	// e4c48e88-3c34-91e3-ab8a-08484dc4d402
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s VisionFlowResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VisionFlowResponseBodyData) GoString() string {
	return s.String()
}

func (s *VisionFlowResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *VisionFlowResponseBodyData) SetTaskId(v string) *VisionFlowResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *VisionFlowResponseBodyData) Validate() error {
	return dara.Validate(s)
}
