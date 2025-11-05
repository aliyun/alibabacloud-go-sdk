// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnDeliverListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v *DescribeCdnDeliverListResponseBodyContent) *DescribeCdnDeliverListResponseBody
	GetContent() *DescribeCdnDeliverListResponseBodyContent
	SetRequestId(v string) *DescribeCdnDeliverListResponseBody
	GetRequestId() *string
}

type DescribeCdnDeliverListResponseBody struct {
	// The information about the tracking task.
	//
	// example:
	//
	// "data": [{"deliverId": 1,"status": "enable","createTime": "2020-10-14T11:19:26Z","crontab": "0 0 0 \\	- \\	- ?","frequency": "d","name": "The name of the tracking task","dmList": ["www.example.com"],"reports": [{"reportId": 1,"conditions": [{"op": "in","field": "prov","value": ["Heilongjiang","Beijing"]}} },{"reportId": 2}],"deliver": {"email": {"subject": "subject","to": ["example@alibaba-inc.com","example@alibaba-inc.com"]}}}]}
	Content *DescribeCdnDeliverListResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 12345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCdnDeliverListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeliverListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeliverListResponseBody) GetContent() *DescribeCdnDeliverListResponseBodyContent {
	return s.Content
}

func (s *DescribeCdnDeliverListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnDeliverListResponseBody) SetContent(v *DescribeCdnDeliverListResponseBodyContent) *DescribeCdnDeliverListResponseBody {
	s.Content = v
	return s
}

func (s *DescribeCdnDeliverListResponseBody) SetRequestId(v string) *DescribeCdnDeliverListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnDeliverListResponseBody) Validate() error {
	if s.Content != nil {
		if err := s.Content.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnDeliverListResponseBodyContent struct {
	Data []*DescribeCdnDeliverListResponseBodyContentData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeCdnDeliverListResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeliverListResponseBodyContent) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeliverListResponseBodyContent) GetData() []*DescribeCdnDeliverListResponseBodyContentData {
	return s.Data
}

func (s *DescribeCdnDeliverListResponseBodyContent) SetData(v []*DescribeCdnDeliverListResponseBodyContentData) *DescribeCdnDeliverListResponseBodyContent {
	s.Data = v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContent) Validate() error {
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

type DescribeCdnDeliverListResponseBodyContentData struct {
	// example:
	//
	// 2021-12-30T10:29:29Z
	CreateTime *string `json:"createTime,omitempty" xml:"createTime,omitempty"`
	// example:
	//
	// 00 00 08 	- 	- ?
	Crontab *string                                               `json:"crontab,omitempty" xml:"crontab,omitempty"`
	Deliver *DescribeCdnDeliverListResponseBodyContentDataDeliver `json:"deliver,omitempty" xml:"deliver,omitempty" type:"Struct"`
	// example:
	//
	// 1
	DeliverId *int64    `json:"deliverId,omitempty" xml:"deliverId,omitempty"`
	DmList    []*string `json:"dmList,omitempty" xml:"dmList,omitempty" type:"Repeated"`
	// example:
	//
	// d
	Frequency *string `json:"frequency,omitempty" xml:"frequency,omitempty"`
	// example:
	//
	// xxxx
	Name    *string                                                 `json:"name,omitempty" xml:"name,omitempty"`
	Reports []*DescribeCdnDeliverListResponseBodyContentDataReports `json:"reports,omitempty" xml:"reports,omitempty" type:"Repeated"`
	// example:
	//
	// enable
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// -1d
	TimeEndFormat *string `json:"timeEndFormat,omitempty" xml:"timeEndFormat,omitempty"`
	// example:
	//
	// -1d
	TimeFromFormat *string `json:"timeFromFormat,omitempty" xml:"timeFromFormat,omitempty"`
}

func (s DescribeCdnDeliverListResponseBodyContentData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeliverListResponseBodyContentData) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeliverListResponseBodyContentData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeCdnDeliverListResponseBodyContentData) GetCrontab() *string {
	return s.Crontab
}

func (s *DescribeCdnDeliverListResponseBodyContentData) GetDeliver() *DescribeCdnDeliverListResponseBodyContentDataDeliver {
	return s.Deliver
}

func (s *DescribeCdnDeliverListResponseBodyContentData) GetDeliverId() *int64 {
	return s.DeliverId
}

func (s *DescribeCdnDeliverListResponseBodyContentData) GetDmList() []*string {
	return s.DmList
}

func (s *DescribeCdnDeliverListResponseBodyContentData) GetFrequency() *string {
	return s.Frequency
}

func (s *DescribeCdnDeliverListResponseBodyContentData) GetName() *string {
	return s.Name
}

func (s *DescribeCdnDeliverListResponseBodyContentData) GetReports() []*DescribeCdnDeliverListResponseBodyContentDataReports {
	return s.Reports
}

func (s *DescribeCdnDeliverListResponseBodyContentData) GetStatus() *string {
	return s.Status
}

func (s *DescribeCdnDeliverListResponseBodyContentData) GetTimeEndFormat() *string {
	return s.TimeEndFormat
}

func (s *DescribeCdnDeliverListResponseBodyContentData) GetTimeFromFormat() *string {
	return s.TimeFromFormat
}

func (s *DescribeCdnDeliverListResponseBodyContentData) SetCreateTime(v string) *DescribeCdnDeliverListResponseBodyContentData {
	s.CreateTime = &v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentData) SetCrontab(v string) *DescribeCdnDeliverListResponseBodyContentData {
	s.Crontab = &v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentData) SetDeliver(v *DescribeCdnDeliverListResponseBodyContentDataDeliver) *DescribeCdnDeliverListResponseBodyContentData {
	s.Deliver = v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentData) SetDeliverId(v int64) *DescribeCdnDeliverListResponseBodyContentData {
	s.DeliverId = &v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentData) SetDmList(v []*string) *DescribeCdnDeliverListResponseBodyContentData {
	s.DmList = v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentData) SetFrequency(v string) *DescribeCdnDeliverListResponseBodyContentData {
	s.Frequency = &v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentData) SetName(v string) *DescribeCdnDeliverListResponseBodyContentData {
	s.Name = &v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentData) SetReports(v []*DescribeCdnDeliverListResponseBodyContentDataReports) *DescribeCdnDeliverListResponseBodyContentData {
	s.Reports = v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentData) SetStatus(v string) *DescribeCdnDeliverListResponseBodyContentData {
	s.Status = &v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentData) SetTimeEndFormat(v string) *DescribeCdnDeliverListResponseBodyContentData {
	s.TimeEndFormat = &v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentData) SetTimeFromFormat(v string) *DescribeCdnDeliverListResponseBodyContentData {
	s.TimeFromFormat = &v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentData) Validate() error {
	if s.Deliver != nil {
		if err := s.Deliver.Validate(); err != nil {
			return err
		}
	}
	if s.Reports != nil {
		for _, item := range s.Reports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCdnDeliverListResponseBodyContentDataDeliver struct {
	Email *DescribeCdnDeliverListResponseBodyContentDataDeliverEmail `json:"email,omitempty" xml:"email,omitempty" type:"Struct"`
}

func (s DescribeCdnDeliverListResponseBodyContentDataDeliver) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeliverListResponseBodyContentDataDeliver) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeliverListResponseBodyContentDataDeliver) GetEmail() *DescribeCdnDeliverListResponseBodyContentDataDeliverEmail {
	return s.Email
}

func (s *DescribeCdnDeliverListResponseBodyContentDataDeliver) SetEmail(v *DescribeCdnDeliverListResponseBodyContentDataDeliverEmail) *DescribeCdnDeliverListResponseBodyContentDataDeliver {
	s.Email = v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentDataDeliver) Validate() error {
	if s.Email != nil {
		if err := s.Email.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCdnDeliverListResponseBodyContentDataDeliverEmail struct {
	To []*string `json:"to,omitempty" xml:"to,omitempty" type:"Repeated"`
}

func (s DescribeCdnDeliverListResponseBodyContentDataDeliverEmail) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeliverListResponseBodyContentDataDeliverEmail) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeliverListResponseBodyContentDataDeliverEmail) GetTo() []*string {
	return s.To
}

func (s *DescribeCdnDeliverListResponseBodyContentDataDeliverEmail) SetTo(v []*string) *DescribeCdnDeliverListResponseBodyContentDataDeliverEmail {
	s.To = v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentDataDeliverEmail) Validate() error {
	return dara.Validate(s)
}

type DescribeCdnDeliverListResponseBodyContentDataReports struct {
	// example:
	//
	// 1
	ReportId *int64 `json:"reportId,omitempty" xml:"reportId,omitempty"`
}

func (s DescribeCdnDeliverListResponseBodyContentDataReports) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnDeliverListResponseBodyContentDataReports) GoString() string {
	return s.String()
}

func (s *DescribeCdnDeliverListResponseBodyContentDataReports) GetReportId() *int64 {
	return s.ReportId
}

func (s *DescribeCdnDeliverListResponseBodyContentDataReports) SetReportId(v int64) *DescribeCdnDeliverListResponseBodyContentDataReports {
	s.ReportId = &v
	return s
}

func (s *DescribeCdnDeliverListResponseBodyContentDataReports) Validate() error {
	return dara.Validate(s)
}
