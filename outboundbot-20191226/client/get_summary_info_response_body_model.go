// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSummaryInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAgentBotInstanceSummaryList(v []*GetSummaryInfoResponseBodyAgentBotInstanceSummaryList) *GetSummaryInfoResponseBody
	GetAgentBotInstanceSummaryList() []*GetSummaryInfoResponseBodyAgentBotInstanceSummaryList
	SetCode(v string) *GetSummaryInfoResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetSummaryInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetSummaryInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSummaryInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSummaryInfoResponseBody
	GetSuccess() *bool
}

type GetSummaryInfoResponseBody struct {
	// example:
	//
	// []
	AgentBotInstanceSummaryList []*GetSummaryInfoResponseBodyAgentBotInstanceSummaryList `json:"AgentBotInstanceSummaryList,omitempty" xml:"AgentBotInstanceSummaryList,omitempty" type:"Repeated"`
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSummaryInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSummaryInfoResponseBody) GetAgentBotInstanceSummaryList() []*GetSummaryInfoResponseBodyAgentBotInstanceSummaryList {
	return s.AgentBotInstanceSummaryList
}

func (s *GetSummaryInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSummaryInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetSummaryInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSummaryInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSummaryInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSummaryInfoResponseBody) SetAgentBotInstanceSummaryList(v []*GetSummaryInfoResponseBodyAgentBotInstanceSummaryList) *GetSummaryInfoResponseBody {
	s.AgentBotInstanceSummaryList = v
	return s
}

func (s *GetSummaryInfoResponseBody) SetCode(v string) *GetSummaryInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetSummaryInfoResponseBody) SetHttpStatusCode(v int32) *GetSummaryInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetSummaryInfoResponseBody) SetMessage(v string) *GetSummaryInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetSummaryInfoResponseBody) SetRequestId(v string) *GetSummaryInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSummaryInfoResponseBody) SetSuccess(v bool) *GetSummaryInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetSummaryInfoResponseBody) Validate() error {
	if s.AgentBotInstanceSummaryList != nil {
		for _, item := range s.AgentBotInstanceSummaryList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSummaryInfoResponseBodyAgentBotInstanceSummaryList struct {
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 12
	TotalCallCount *int64 `json:"TotalCallCount,omitempty" xml:"TotalCallCount,omitempty"`
	// example:
	//
	// 10
	TotalCallTime *int64 `json:"TotalCallTime,omitempty" xml:"TotalCallTime,omitempty"`
	// example:
	//
	// 10
	UsedRecordingStorageSpace *int32 `json:"UsedRecordingStorageSpace,omitempty" xml:"UsedRecordingStorageSpace,omitempty"`
}

func (s GetSummaryInfoResponseBodyAgentBotInstanceSummaryList) String() string {
	return dara.Prettify(s)
}

func (s GetSummaryInfoResponseBodyAgentBotInstanceSummaryList) GoString() string {
	return s.String()
}

func (s *GetSummaryInfoResponseBodyAgentBotInstanceSummaryList) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSummaryInfoResponseBodyAgentBotInstanceSummaryList) GetTotalCallCount() *int64 {
	return s.TotalCallCount
}

func (s *GetSummaryInfoResponseBodyAgentBotInstanceSummaryList) GetTotalCallTime() *int64 {
	return s.TotalCallTime
}

func (s *GetSummaryInfoResponseBodyAgentBotInstanceSummaryList) GetUsedRecordingStorageSpace() *int32 {
	return s.UsedRecordingStorageSpace
}

func (s *GetSummaryInfoResponseBodyAgentBotInstanceSummaryList) SetInstanceId(v string) *GetSummaryInfoResponseBodyAgentBotInstanceSummaryList {
	s.InstanceId = &v
	return s
}

func (s *GetSummaryInfoResponseBodyAgentBotInstanceSummaryList) SetTotalCallCount(v int64) *GetSummaryInfoResponseBodyAgentBotInstanceSummaryList {
	s.TotalCallCount = &v
	return s
}

func (s *GetSummaryInfoResponseBodyAgentBotInstanceSummaryList) SetTotalCallTime(v int64) *GetSummaryInfoResponseBodyAgentBotInstanceSummaryList {
	s.TotalCallTime = &v
	return s
}

func (s *GetSummaryInfoResponseBodyAgentBotInstanceSummaryList) SetUsedRecordingStorageSpace(v int32) *GetSummaryInfoResponseBodyAgentBotInstanceSummaryList {
	s.UsedRecordingStorageSpace = &v
	return s
}

func (s *GetSummaryInfoResponseBodyAgentBotInstanceSummaryList) Validate() error {
	return dara.Validate(s)
}
