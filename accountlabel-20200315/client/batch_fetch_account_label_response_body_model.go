// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchFetchAccountLabelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *BatchFetchAccountLabelResponseBody
	GetCode() *string
	SetCount(v int64) *BatchFetchAccountLabelResponseBody
	GetCount() *int64
	SetData(v []*BatchFetchAccountLabelResponseBodyData) *BatchFetchAccountLabelResponseBody
	GetData() []*BatchFetchAccountLabelResponseBodyData
	SetMessage(v string) *BatchFetchAccountLabelResponseBody
	GetMessage() *string
	SetRequestId(v string) *BatchFetchAccountLabelResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BatchFetchAccountLabelResponseBody
	GetSuccess() *bool
}

type BatchFetchAccountLabelResponseBody struct {
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Count     *int64                                    `json:"Count,omitempty" xml:"Count,omitempty"`
	Data      []*BatchFetchAccountLabelResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                                     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BatchFetchAccountLabelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BatchFetchAccountLabelResponseBody) GoString() string {
	return s.String()
}

func (s *BatchFetchAccountLabelResponseBody) GetCode() *string {
	return s.Code
}

func (s *BatchFetchAccountLabelResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *BatchFetchAccountLabelResponseBody) GetData() []*BatchFetchAccountLabelResponseBodyData {
	return s.Data
}

func (s *BatchFetchAccountLabelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *BatchFetchAccountLabelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BatchFetchAccountLabelResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BatchFetchAccountLabelResponseBody) SetCode(v string) *BatchFetchAccountLabelResponseBody {
	s.Code = &v
	return s
}

func (s *BatchFetchAccountLabelResponseBody) SetCount(v int64) *BatchFetchAccountLabelResponseBody {
	s.Count = &v
	return s
}

func (s *BatchFetchAccountLabelResponseBody) SetData(v []*BatchFetchAccountLabelResponseBodyData) *BatchFetchAccountLabelResponseBody {
	s.Data = v
	return s
}

func (s *BatchFetchAccountLabelResponseBody) SetMessage(v string) *BatchFetchAccountLabelResponseBody {
	s.Message = &v
	return s
}

func (s *BatchFetchAccountLabelResponseBody) SetRequestId(v string) *BatchFetchAccountLabelResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchFetchAccountLabelResponseBody) SetSuccess(v bool) *BatchFetchAccountLabelResponseBody {
	s.Success = &v
	return s
}

func (s *BatchFetchAccountLabelResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type BatchFetchAccountLabelResponseBodyData struct {
	Creator     *string `json:"Creator,omitempty" xml:"Creator,omitempty"`
	EndTime     *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	GmtCreated  *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	Kp          *int64  `json:"Kp,omitempty" xml:"Kp,omitempty"`
	Label       *string `json:"Label,omitempty" xml:"Label,omitempty"`
	LabelSeries *string `json:"LabelSeries,omitempty" xml:"LabelSeries,omitempty"`
	StartTime   *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s BatchFetchAccountLabelResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BatchFetchAccountLabelResponseBodyData) GoString() string {
	return s.String()
}

func (s *BatchFetchAccountLabelResponseBodyData) GetCreator() *string {
	return s.Creator
}

func (s *BatchFetchAccountLabelResponseBodyData) GetEndTime() *string {
	return s.EndTime
}

func (s *BatchFetchAccountLabelResponseBodyData) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *BatchFetchAccountLabelResponseBodyData) GetGmtModified() *string {
	return s.GmtModified
}

func (s *BatchFetchAccountLabelResponseBodyData) GetKp() *int64 {
	return s.Kp
}

func (s *BatchFetchAccountLabelResponseBodyData) GetLabel() *string {
	return s.Label
}

func (s *BatchFetchAccountLabelResponseBodyData) GetLabelSeries() *string {
	return s.LabelSeries
}

func (s *BatchFetchAccountLabelResponseBodyData) GetStartTime() *string {
	return s.StartTime
}

func (s *BatchFetchAccountLabelResponseBodyData) SetCreator(v string) *BatchFetchAccountLabelResponseBodyData {
	s.Creator = &v
	return s
}

func (s *BatchFetchAccountLabelResponseBodyData) SetEndTime(v string) *BatchFetchAccountLabelResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *BatchFetchAccountLabelResponseBodyData) SetGmtCreated(v string) *BatchFetchAccountLabelResponseBodyData {
	s.GmtCreated = &v
	return s
}

func (s *BatchFetchAccountLabelResponseBodyData) SetGmtModified(v string) *BatchFetchAccountLabelResponseBodyData {
	s.GmtModified = &v
	return s
}

func (s *BatchFetchAccountLabelResponseBodyData) SetKp(v int64) *BatchFetchAccountLabelResponseBodyData {
	s.Kp = &v
	return s
}

func (s *BatchFetchAccountLabelResponseBodyData) SetLabel(v string) *BatchFetchAccountLabelResponseBodyData {
	s.Label = &v
	return s
}

func (s *BatchFetchAccountLabelResponseBodyData) SetLabelSeries(v string) *BatchFetchAccountLabelResponseBodyData {
	s.LabelSeries = &v
	return s
}

func (s *BatchFetchAccountLabelResponseBodyData) SetStartTime(v string) *BatchFetchAccountLabelResponseBodyData {
	s.StartTime = &v
	return s
}

func (s *BatchFetchAccountLabelResponseBodyData) Validate() error {
	return dara.Validate(s)
}
