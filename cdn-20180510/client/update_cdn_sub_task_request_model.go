// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCdnSubTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *UpdateCdnSubTaskRequest
	GetDomainName() *string
	SetEndTime(v string) *UpdateCdnSubTaskRequest
	GetEndTime() *string
	SetReportIds(v string) *UpdateCdnSubTaskRequest
	GetReportIds() *string
	SetStartTime(v string) *UpdateCdnSubTaskRequest
	GetStartTime() *string
}

type UpdateCdnSubTaskRequest struct {
	// The domain name that you want to track. You can specify up to 500 domain names in each request. If you specify multiple domain names, separate them with commas (,). If you do not specify a domain name, operations reports are updated for all domain names in your Alibaba Cloud account.
	//
	// example:
	//
	// www.example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The end time of the operations report. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-11-17T00:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IDs of operations reports that you want to update. Separate IDs with commas (,).
	//
	// example:
	//
	// 1,2,3
	ReportIds *string `json:"ReportIds,omitempty" xml:"ReportIds,omitempty"`
	// The start time of the operations report. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2020-09-17T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s UpdateCdnSubTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCdnSubTaskRequest) GoString() string {
	return s.String()
}

func (s *UpdateCdnSubTaskRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UpdateCdnSubTaskRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *UpdateCdnSubTaskRequest) GetReportIds() *string {
	return s.ReportIds
}

func (s *UpdateCdnSubTaskRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *UpdateCdnSubTaskRequest) SetDomainName(v string) *UpdateCdnSubTaskRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateCdnSubTaskRequest) SetEndTime(v string) *UpdateCdnSubTaskRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateCdnSubTaskRequest) SetReportIds(v string) *UpdateCdnSubTaskRequest {
	s.ReportIds = &v
	return s
}

func (s *UpdateCdnSubTaskRequest) SetStartTime(v string) *UpdateCdnSubTaskRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateCdnSubTaskRequest) Validate() error {
	return dara.Validate(s)
}
