// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContentDetectResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetContentDetectResultResponseBody
	GetCode() *string
	SetData(v *GetContentDetectResultResponseBodyData) *GetContentDetectResultResponseBody
	GetData() *GetContentDetectResultResponseBodyData
	SetHttpStatusCode(v int32) *GetContentDetectResultResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetContentDetectResultResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetContentDetectResultResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetContentDetectResultResponseBody
	GetSuccess() *bool
}

type GetContentDetectResultResponseBody struct {
	// example:
	//
	// 00000
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetContentDetectResultResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// ""
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetContentDetectResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetContentDetectResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetContentDetectResultResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetContentDetectResultResponseBody) GetData() *GetContentDetectResultResponseBodyData {
	return s.Data
}

func (s *GetContentDetectResultResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetContentDetectResultResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetContentDetectResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetContentDetectResultResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetContentDetectResultResponseBody) SetCode(v string) *GetContentDetectResultResponseBody {
	s.Code = &v
	return s
}

func (s *GetContentDetectResultResponseBody) SetData(v *GetContentDetectResultResponseBodyData) *GetContentDetectResultResponseBody {
	s.Data = v
	return s
}

func (s *GetContentDetectResultResponseBody) SetHttpStatusCode(v int32) *GetContentDetectResultResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetContentDetectResultResponseBody) SetMessage(v string) *GetContentDetectResultResponseBody {
	s.Message = &v
	return s
}

func (s *GetContentDetectResultResponseBody) SetRequestId(v string) *GetContentDetectResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetContentDetectResultResponseBody) SetSuccess(v bool) *GetContentDetectResultResponseBody {
	s.Success = &v
	return s
}

func (s *GetContentDetectResultResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetContentDetectResultResponseBodyData struct {
	DetectResultList []*GetContentDetectResultResponseBodyDataDetectResultList `json:"DetectResultList,omitempty" xml:"DetectResultList,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	ProcessedCount *int32  `json:"ProcessedCount,omitempty" xml:"ProcessedCount,omitempty"`
	TaskId         *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// example:
	//
	// 2
	TaskStatus *int32 `json:"TaskStatus,omitempty" xml:"TaskStatus,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetContentDetectResultResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetContentDetectResultResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetContentDetectResultResponseBodyData) GetDetectResultList() []*GetContentDetectResultResponseBodyDataDetectResultList {
	return s.DetectResultList
}

func (s *GetContentDetectResultResponseBodyData) GetProcessedCount() *int32 {
	return s.ProcessedCount
}

func (s *GetContentDetectResultResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetContentDetectResultResponseBodyData) GetTaskStatus() *int32 {
	return s.TaskStatus
}

func (s *GetContentDetectResultResponseBodyData) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetContentDetectResultResponseBodyData) SetDetectResultList(v []*GetContentDetectResultResponseBodyDataDetectResultList) *GetContentDetectResultResponseBodyData {
	s.DetectResultList = v
	return s
}

func (s *GetContentDetectResultResponseBodyData) SetProcessedCount(v int32) *GetContentDetectResultResponseBodyData {
	s.ProcessedCount = &v
	return s
}

func (s *GetContentDetectResultResponseBodyData) SetTaskId(v string) *GetContentDetectResultResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetContentDetectResultResponseBodyData) SetTaskStatus(v int32) *GetContentDetectResultResponseBodyData {
	s.TaskStatus = &v
	return s
}

func (s *GetContentDetectResultResponseBodyData) SetTotalCount(v int32) *GetContentDetectResultResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *GetContentDetectResultResponseBodyData) Validate() error {
	if s.DetectResultList != nil {
		for _, item := range s.DetectResultList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetContentDetectResultResponseBodyDataDetectResultList struct {
	RiskInfo *string `json:"RiskInfo,omitempty" xml:"RiskInfo,omitempty"`
	// example:
	//
	// 1
	RiskResult *int32 `json:"RiskResult,omitempty" xml:"RiskResult,omitempty"`
	// example:
	//
	// 2
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetContentDetectResultResponseBodyDataDetectResultList) String() string {
	return dara.Prettify(s)
}

func (s GetContentDetectResultResponseBodyDataDetectResultList) GoString() string {
	return s.String()
}

func (s *GetContentDetectResultResponseBodyDataDetectResultList) GetRiskInfo() *string {
	return s.RiskInfo
}

func (s *GetContentDetectResultResponseBodyDataDetectResultList) GetRiskResult() *int32 {
	return s.RiskResult
}

func (s *GetContentDetectResultResponseBodyDataDetectResultList) GetStatus() *int32 {
	return s.Status
}

func (s *GetContentDetectResultResponseBodyDataDetectResultList) SetRiskInfo(v string) *GetContentDetectResultResponseBodyDataDetectResultList {
	s.RiskInfo = &v
	return s
}

func (s *GetContentDetectResultResponseBodyDataDetectResultList) SetRiskResult(v int32) *GetContentDetectResultResponseBodyDataDetectResultList {
	s.RiskResult = &v
	return s
}

func (s *GetContentDetectResultResponseBodyDataDetectResultList) SetStatus(v int32) *GetContentDetectResultResponseBodyDataDetectResultList {
	s.Status = &v
	return s
}

func (s *GetContentDetectResultResponseBodyDataDetectResultList) Validate() error {
	return dara.Validate(s)
}
