// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogId(v int64) *DeleteLogMonitorRequest
	GetLogId() *int64
	SetRegionId(v string) *DeleteLogMonitorRequest
	GetRegionId() *string
}

type DeleteLogMonitorRequest struct {
	// The ID of the log monitoring metric.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12345
	LogId    *int64  `json:"LogId,omitempty" xml:"LogId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteLogMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogMonitorRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogMonitorRequest) GetLogId() *int64 {
	return s.LogId
}

func (s *DeleteLogMonitorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLogMonitorRequest) SetLogId(v int64) *DeleteLogMonitorRequest {
	s.LogId = &v
	return s
}

func (s *DeleteLogMonitorRequest) SetRegionId(v string) *DeleteLogMonitorRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLogMonitorRequest) Validate() error {
	return dara.Validate(s)
}
