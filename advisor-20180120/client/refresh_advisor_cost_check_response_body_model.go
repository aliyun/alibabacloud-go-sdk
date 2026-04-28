// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAdvisorCostCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RefreshAdvisorCostCheckResponseBody
	GetCode() *string
	SetData(v *RefreshAdvisorCostCheckResponseBodyData) *RefreshAdvisorCostCheckResponseBody
	GetData() *RefreshAdvisorCostCheckResponseBodyData
	SetMessage(v string) *RefreshAdvisorCostCheckResponseBody
	GetMessage() *string
	SetRequestId(v string) *RefreshAdvisorCostCheckResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RefreshAdvisorCostCheckResponseBody
	GetSuccess() *bool
}

type RefreshAdvisorCostCheckResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *RefreshAdvisorCostCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 566331F9-****-550F-B745-A730331F97A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RefreshAdvisorCostCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorCostCheckResponseBody) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCostCheckResponseBody) GetCode() *string {
	return s.Code
}

func (s *RefreshAdvisorCostCheckResponseBody) GetData() *RefreshAdvisorCostCheckResponseBodyData {
	return s.Data
}

func (s *RefreshAdvisorCostCheckResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RefreshAdvisorCostCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RefreshAdvisorCostCheckResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RefreshAdvisorCostCheckResponseBody) SetCode(v string) *RefreshAdvisorCostCheckResponseBody {
	s.Code = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBody) SetData(v *RefreshAdvisorCostCheckResponseBodyData) *RefreshAdvisorCostCheckResponseBody {
	s.Data = v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBody) SetMessage(v string) *RefreshAdvisorCostCheckResponseBody {
	s.Message = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBody) SetRequestId(v string) *RefreshAdvisorCostCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBody) SetSuccess(v bool) *RefreshAdvisorCostCheckResponseBody {
	s.Success = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RefreshAdvisorCostCheckResponseBodyData struct {
	// example:
	//
	// c-wl*****n0g
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// example:
	//
	// 11***********35
	ManagerTaskId *int64 `json:"ManagerTaskId,omitempty" xml:"ManagerTaskId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 959***135
	TaskId *int64 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RefreshAdvisorCostCheckResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorCostCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCostCheckResponseBodyData) GetCommandId() *string {
	return s.CommandId
}

func (s *RefreshAdvisorCostCheckResponseBodyData) GetManagerTaskId() *int64 {
	return s.ManagerTaskId
}

func (s *RefreshAdvisorCostCheckResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *RefreshAdvisorCostCheckResponseBodyData) GetTaskId() *int64 {
	return s.TaskId
}

func (s *RefreshAdvisorCostCheckResponseBodyData) SetCommandId(v string) *RefreshAdvisorCostCheckResponseBodyData {
	s.CommandId = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBodyData) SetManagerTaskId(v int64) *RefreshAdvisorCostCheckResponseBodyData {
	s.ManagerTaskId = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBodyData) SetSuccess(v bool) *RefreshAdvisorCostCheckResponseBodyData {
	s.Success = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBodyData) SetTaskId(v int64) *RefreshAdvisorCostCheckResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *RefreshAdvisorCostCheckResponseBodyData) Validate() error {
	return dara.Validate(s)
}
