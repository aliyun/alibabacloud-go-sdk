// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfirmDjbhReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ConfirmDjbhReportRequest
	GetId() *int64
}

type ConfirmDjbhReportRequest struct {
	// Primary key ID of the report.
	//
	// example:
	//
	// 24563
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ConfirmDjbhReportRequest) String() string {
	return dara.Prettify(s)
}

func (s ConfirmDjbhReportRequest) GoString() string {
	return s.String()
}

func (s *ConfirmDjbhReportRequest) GetId() *int64 {
	return s.Id
}

func (s *ConfirmDjbhReportRequest) SetId(v int64) *ConfirmDjbhReportRequest {
	s.Id = &v
	return s
}

func (s *ConfirmDjbhReportRequest) Validate() error {
	return dara.Validate(s)
}
