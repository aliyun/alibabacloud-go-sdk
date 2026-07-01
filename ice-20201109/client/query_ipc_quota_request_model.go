// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIpcQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCapability(v string) *QueryIpcQuotaRequest
	GetCapability() *string
	SetEndTime(v string) *QueryIpcQuotaRequest
	GetEndTime() *string
	SetPageNo(v int32) *QueryIpcQuotaRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryIpcQuotaRequest
	GetPageSize() *int32
	SetStartTime(v string) *QueryIpcQuotaRequest
	GetStartTime() *string
}

type QueryIpcQuotaRequest struct {
	// Capability. Default value is understand. Valid values:
	//
	// - understand: understanding
	//
	// - understand-reid: understanding with reid
	//
	// - search: search
	//
	// example:
	//
	// understand
	Capability *string `json:"Capability,omitempty" xml:"Capability,omitempty"`
	// End time. UTC time in the format: yyyy-MM-ddTHH:mm:ssZ. The difference between EndTime and StartTime must not exceed 180 days.
	//
	// example:
	//
	// 2025-05-09T08:52:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Page number for paging. Default Value: 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// Page size, which specifies the number of entries displayed per page. Default value is 10, and the maximum value is 50.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Start Time. UTC time in the format: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2025-05-08T08:52:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s QueryIpcQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryIpcQuotaRequest) GoString() string {
	return s.String()
}

func (s *QueryIpcQuotaRequest) GetCapability() *string {
	return s.Capability
}

func (s *QueryIpcQuotaRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *QueryIpcQuotaRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryIpcQuotaRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryIpcQuotaRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *QueryIpcQuotaRequest) SetCapability(v string) *QueryIpcQuotaRequest {
	s.Capability = &v
	return s
}

func (s *QueryIpcQuotaRequest) SetEndTime(v string) *QueryIpcQuotaRequest {
	s.EndTime = &v
	return s
}

func (s *QueryIpcQuotaRequest) SetPageNo(v int32) *QueryIpcQuotaRequest {
	s.PageNo = &v
	return s
}

func (s *QueryIpcQuotaRequest) SetPageSize(v int32) *QueryIpcQuotaRequest {
	s.PageSize = &v
	return s
}

func (s *QueryIpcQuotaRequest) SetStartTime(v string) *QueryIpcQuotaRequest {
	s.StartTime = &v
	return s
}

func (s *QueryIpcQuotaRequest) Validate() error {
	return dara.Validate(s)
}
