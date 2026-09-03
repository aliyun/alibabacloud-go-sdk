// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVirtualTryOnResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VirtualTryOnResponseBody
	GetCode() *string
	SetData(v *VirtualTryOnResponseBodyData) *VirtualTryOnResponseBody
	GetData() *VirtualTryOnResponseBodyData
	SetMessage(v string) *VirtualTryOnResponseBody
	GetMessage() *string
	SetRequestId(v string) *VirtualTryOnResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *VirtualTryOnResponseBody
	GetSuccess() *bool
}

type VirtualTryOnResponseBody struct {
	// The error code. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned result struct.
	Data *VirtualTryOnResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if the call is successful.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. Used to uniquely identify a single API call.
	//
	// example:
	//
	// E1AD60F1-BAC7-546B-9533-E7AD02B16E3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates success. A value of false indicates failure.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s VirtualTryOnResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VirtualTryOnResponseBody) GoString() string {
	return s.String()
}

func (s *VirtualTryOnResponseBody) GetCode() *string {
	return s.Code
}

func (s *VirtualTryOnResponseBody) GetData() *VirtualTryOnResponseBodyData {
	return s.Data
}

func (s *VirtualTryOnResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VirtualTryOnResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VirtualTryOnResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *VirtualTryOnResponseBody) SetCode(v string) *VirtualTryOnResponseBody {
	s.Code = &v
	return s
}

func (s *VirtualTryOnResponseBody) SetData(v *VirtualTryOnResponseBodyData) *VirtualTryOnResponseBody {
	s.Data = v
	return s
}

func (s *VirtualTryOnResponseBody) SetMessage(v string) *VirtualTryOnResponseBody {
	s.Message = &v
	return s
}

func (s *VirtualTryOnResponseBody) SetRequestId(v string) *VirtualTryOnResponseBody {
	s.RequestId = &v
	return s
}

func (s *VirtualTryOnResponseBody) SetSuccess(v bool) *VirtualTryOnResponseBody {
	s.Success = &v
	return s
}

func (s *VirtualTryOnResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VirtualTryOnResponseBodyData struct {
	// The asynchronous task ID. Used to query the task result later.
	//
	// example:
	//
	// task-xxxx-xxxx-xxxx
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The usage details.
	UsageMap map[string]interface{} `json:"UsageMap,omitempty" xml:"UsageMap,omitempty"`
}

func (s VirtualTryOnResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s VirtualTryOnResponseBodyData) GoString() string {
	return s.String()
}

func (s *VirtualTryOnResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *VirtualTryOnResponseBodyData) GetUsageMap() map[string]interface{} {
	return s.UsageMap
}

func (s *VirtualTryOnResponseBodyData) SetTaskId(v string) *VirtualTryOnResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *VirtualTryOnResponseBodyData) SetUsageMap(v map[string]interface{}) *VirtualTryOnResponseBodyData {
	s.UsageMap = v
	return s
}

func (s *VirtualTryOnResponseBodyData) Validate() error {
	return dara.Validate(s)
}
