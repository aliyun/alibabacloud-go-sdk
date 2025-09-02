// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceCountTrendRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *GetInstanceCountTrendRequest
	GetBeginDate() *string
	SetEndDate(v string) *GetInstanceCountTrendRequest
	GetEndDate() *string
	SetProjectId(v int64) *GetInstanceCountTrendRequest
	GetProjectId() *int64
}

type GetInstanceCountTrendRequest struct {
	// The beginning of the time range to query, accurate to the day. Specify the time in the ISO 8601 standard in the yyyy-MM-dd\\"T\\"HH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-04-02T00:00:00+0800
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// The end of the time range to query, accurate to the day. Specify the time in the ISO 8601 standard in the yyyy-MM-dd\\"T\\"HH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2020-09-10T00:00:00+0800
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The DataWorks workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GetInstanceCountTrendRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceCountTrendRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceCountTrendRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *GetInstanceCountTrendRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetInstanceCountTrendRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *GetInstanceCountTrendRequest) SetBeginDate(v string) *GetInstanceCountTrendRequest {
	s.BeginDate = &v
	return s
}

func (s *GetInstanceCountTrendRequest) SetEndDate(v string) *GetInstanceCountTrendRequest {
	s.EndDate = &v
	return s
}

func (s *GetInstanceCountTrendRequest) SetProjectId(v int64) *GetInstanceCountTrendRequest {
	s.ProjectId = &v
	return s
}

func (s *GetInstanceCountTrendRequest) Validate() error {
	return dara.Validate(s)
}
