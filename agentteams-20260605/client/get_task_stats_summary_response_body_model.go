// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTaskStatsSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTaskStatsSummaryResponseBody
	GetCode() *string
	SetData(v *GetTaskStatsSummaryResponseBodyData) *GetTaskStatsSummaryResponseBody
	GetData() *GetTaskStatsSummaryResponseBodyData
	SetHttpStatusCode(v int32) *GetTaskStatsSummaryResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetTaskStatsSummaryResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTaskStatsSummaryResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetTaskStatsSummaryResponseBody
	GetSuccess() *bool
}

type GetTaskStatsSummaryResponseBody struct {
	Code           *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *GetTaskStatsSummaryResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                               `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetTaskStatsSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatsSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatsSummaryResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTaskStatsSummaryResponseBody) GetData() *GetTaskStatsSummaryResponseBodyData {
	return s.Data
}

func (s *GetTaskStatsSummaryResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetTaskStatsSummaryResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTaskStatsSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTaskStatsSummaryResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetTaskStatsSummaryResponseBody) SetCode(v string) *GetTaskStatsSummaryResponseBody {
	s.Code = &v
	return s
}

func (s *GetTaskStatsSummaryResponseBody) SetData(v *GetTaskStatsSummaryResponseBodyData) *GetTaskStatsSummaryResponseBody {
	s.Data = v
	return s
}

func (s *GetTaskStatsSummaryResponseBody) SetHttpStatusCode(v int32) *GetTaskStatsSummaryResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetTaskStatsSummaryResponseBody) SetMessage(v string) *GetTaskStatsSummaryResponseBody {
	s.Message = &v
	return s
}

func (s *GetTaskStatsSummaryResponseBody) SetRequestId(v string) *GetTaskStatsSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTaskStatsSummaryResponseBody) SetSuccess(v bool) *GetTaskStatsSummaryResponseBody {
	s.Success = &v
	return s
}

func (s *GetTaskStatsSummaryResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTaskStatsSummaryResponseBodyData struct {
	AverageTaskDuration  *float64                                                 `json:"AverageTaskDuration,omitempty" xml:"AverageTaskDuration,omitempty"`
	StatusDistribution   []*GetTaskStatsSummaryResponseBodyDataStatusDistribution `json:"StatusDistribution,omitempty" xml:"StatusDistribution,omitempty" type:"Repeated"`
	TaskTokenConsumption *int64                                                   `json:"TaskTokenConsumption,omitempty" xml:"TaskTokenConsumption,omitempty"`
	TotalTasks           *int32                                                   `json:"TotalTasks,omitempty" xml:"TotalTasks,omitempty"`
}

func (s GetTaskStatsSummaryResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatsSummaryResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTaskStatsSummaryResponseBodyData) GetAverageTaskDuration() *float64 {
	return s.AverageTaskDuration
}

func (s *GetTaskStatsSummaryResponseBodyData) GetStatusDistribution() []*GetTaskStatsSummaryResponseBodyDataStatusDistribution {
	return s.StatusDistribution
}

func (s *GetTaskStatsSummaryResponseBodyData) GetTaskTokenConsumption() *int64 {
	return s.TaskTokenConsumption
}

func (s *GetTaskStatsSummaryResponseBodyData) GetTotalTasks() *int32 {
	return s.TotalTasks
}

func (s *GetTaskStatsSummaryResponseBodyData) SetAverageTaskDuration(v float64) *GetTaskStatsSummaryResponseBodyData {
	s.AverageTaskDuration = &v
	return s
}

func (s *GetTaskStatsSummaryResponseBodyData) SetStatusDistribution(v []*GetTaskStatsSummaryResponseBodyDataStatusDistribution) *GetTaskStatsSummaryResponseBodyData {
	s.StatusDistribution = v
	return s
}

func (s *GetTaskStatsSummaryResponseBodyData) SetTaskTokenConsumption(v int64) *GetTaskStatsSummaryResponseBodyData {
	s.TaskTokenConsumption = &v
	return s
}

func (s *GetTaskStatsSummaryResponseBodyData) SetTotalTasks(v int32) *GetTaskStatsSummaryResponseBodyData {
	s.TotalTasks = &v
	return s
}

func (s *GetTaskStatsSummaryResponseBodyData) Validate() error {
	if s.StatusDistribution != nil {
		for _, item := range s.StatusDistribution {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTaskStatsSummaryResponseBodyDataStatusDistribution struct {
	Count  *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetTaskStatsSummaryResponseBodyDataStatusDistribution) String() string {
	return dara.Prettify(s)
}

func (s GetTaskStatsSummaryResponseBodyDataStatusDistribution) GoString() string {
	return s.String()
}

func (s *GetTaskStatsSummaryResponseBodyDataStatusDistribution) GetCount() *int32 {
	return s.Count
}

func (s *GetTaskStatsSummaryResponseBodyDataStatusDistribution) GetStatus() *string {
	return s.Status
}

func (s *GetTaskStatsSummaryResponseBodyDataStatusDistribution) SetCount(v int32) *GetTaskStatsSummaryResponseBodyDataStatusDistribution {
	s.Count = &v
	return s
}

func (s *GetTaskStatsSummaryResponseBodyDataStatusDistribution) SetStatus(v string) *GetTaskStatsSummaryResponseBodyDataStatusDistribution {
	s.Status = &v
	return s
}

func (s *GetTaskStatsSummaryResponseBodyDataStatusDistribution) Validate() error {
	return dara.Validate(s)
}
