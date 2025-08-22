// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDcdnSubTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *UpdateDcdnSubTaskRequest
	GetDomainName() *string
	SetEndTime(v string) *UpdateDcdnSubTaskRequest
	GetEndTime() *string
	SetReportIds(v string) *UpdateDcdnSubTaskRequest
	GetReportIds() *string
	SetStartTime(v string) *UpdateDcdnSubTaskRequest
	GetStartTime() *string
}

type UpdateDcdnSubTaskRequest struct {
	// The domain names that you want to include in the operations report. If you do not specify a domain name, all domain names that belong to your Alibaba Cloud account are included.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time of the operations report. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-06-17T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IDs of the metrics that you want to update. Separate IDs with commas (,). You can call the [DescribeDcdnSubList](https://help.aliyun.com/document_detail/270075.html) operation to query the IDs.
	//
	// example:
	//
	// 2,4,6
	ReportIds *string `json:"ReportIds,omitempty" xml:"ReportIds,omitempty"`
	// The start time of the operations report. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2021-04-17T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s UpdateDcdnSubTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDcdnSubTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateDcdnSubTaskRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateDcdnSubTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateDcdnSubTaskRequest) GetReportIds() *string {
	return s.ReportIds
}

func (s *UpdateDcdnSubTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateDcdnSubTaskRequest) SetDomainName(v string) *UpdateDcdnSubTaskRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDcdnSubTaskRequest) SetEndTime(v string) *UpdateDcdnSubTaskRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateDcdnSubTaskRequest) SetReportIds(v string) *UpdateDcdnSubTaskRequest {
	s.ReportIds = &v
	return s
}

func (s *UpdateDcdnSubTaskRequest) SetStartTime(v string) *UpdateDcdnSubTaskRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateDcdnSubTaskRequest) Validate() error {
	return dara.Validate(s)
}
