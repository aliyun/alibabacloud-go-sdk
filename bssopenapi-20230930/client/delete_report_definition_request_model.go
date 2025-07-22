// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteReportDefinitionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNbid(v string) *DeleteReportDefinitionRequest
	GetNbid() *string
	SetReportTaskId(v int64) *DeleteReportDefinitionRequest
	GetReportTaskId() *int64
}

type DeleteReportDefinitionRequest struct {
	// example:
	//
	// 2684201000001
	Nbid *string `json:"Nbid,omitempty" xml:"Nbid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 123123
	ReportTaskId *int64 `json:"ReportTaskId,omitempty" xml:"ReportTaskId,omitempty"`
}

func (s DeleteReportDefinitionRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteReportDefinitionRequest) GoString() string {
	return s.String()
}

func (s *DeleteReportDefinitionRequest) GetNbid() *string {
	return s.Nbid
}

func (s *DeleteReportDefinitionRequest) GetReportTaskId() *int64 {
	return s.ReportTaskId
}

func (s *DeleteReportDefinitionRequest) SetNbid(v string) *DeleteReportDefinitionRequest {
	s.Nbid = &v
	return s
}

func (s *DeleteReportDefinitionRequest) SetReportTaskId(v int64) *DeleteReportDefinitionRequest {
	s.ReportTaskId = &v
	return s
}

func (s *DeleteReportDefinitionRequest) Validate() error {
	return dara.Validate(s)
}
