// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendCustomizeReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v int64) *SendCustomizeReportRequest
	GetReportId() *int64
	SetResourceDirectoryAccountId(v int64) *SendCustomizeReportRequest
	GetResourceDirectoryAccountId() *int64
}

type SendCustomizeReportRequest struct {
	// The report ID.
	//
	// >Call [DescribeCustomizeReportList](~~DescribeCustomizeReportList~~) to obtain this parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	ReportId                   *int64 `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	ResourceDirectoryAccountId *int64 `json:"ResourceDirectoryAccountId,omitempty" xml:"ResourceDirectoryAccountId,omitempty"`
}

func (s SendCustomizeReportRequest) String() string {
	return dara.Prettify(s)
}

func (s SendCustomizeReportRequest) GoString() string {
	return s.String()
}

func (s *SendCustomizeReportRequest) GetReportId() *int64 {
	return s.ReportId
}

func (s *SendCustomizeReportRequest) GetResourceDirectoryAccountId() *int64 {
	return s.ResourceDirectoryAccountId
}

func (s *SendCustomizeReportRequest) SetReportId(v int64) *SendCustomizeReportRequest {
	s.ReportId = &v
	return s
}

func (s *SendCustomizeReportRequest) SetResourceDirectoryAccountId(v int64) *SendCustomizeReportRequest {
	s.ResourceDirectoryAccountId = &v
	return s
}

func (s *SendCustomizeReportRequest) Validate() error {
	return dara.Validate(s)
}
