// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseDataExportJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v int64) *PauseDataExportJobRequest
	GetJobId() *int64
	SetOrderId(v int64) *PauseDataExportJobRequest
	GetOrderId() *int64
	SetTid(v int64) *PauseDataExportJobRequest
	GetTid() *int64
}

type PauseDataExportJobRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1276****
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 546****
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s PauseDataExportJobRequest) String() string {
	return dara.Prettify(s)
}

func (s PauseDataExportJobRequest) GoString() string {
	return s.String()
}

func (s *PauseDataExportJobRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *PauseDataExportJobRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *PauseDataExportJobRequest) GetTid() *int64 {
	return s.Tid
}

func (s *PauseDataExportJobRequest) SetJobId(v int64) *PauseDataExportJobRequest {
	s.JobId = &v
	return s
}

func (s *PauseDataExportJobRequest) SetOrderId(v int64) *PauseDataExportJobRequest {
	s.OrderId = &v
	return s
}

func (s *PauseDataExportJobRequest) SetTid(v int64) *PauseDataExportJobRequest {
	s.Tid = &v
	return s
}

func (s *PauseDataExportJobRequest) Validate() error {
	return dara.Validate(s)
}
