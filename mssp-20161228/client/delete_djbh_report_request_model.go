// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDjbhReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteDjbhReportRequest
	GetId() *int64
}

type DeleteDjbhReportRequest struct {
	// Primary key ID of the report.
	//
	// This parameter is required.
	//
	// example:
	//
	// 26579
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteDjbhReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDjbhReportRequest) GoString() string {
	return s.String()
}

func (s *DeleteDjbhReportRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteDjbhReportRequest) SetId(v int64) *DeleteDjbhReportRequest {
	s.Id = &v
	return s
}

func (s *DeleteDjbhReportRequest) Validate() error {
	return dara.Validate(s)
}
