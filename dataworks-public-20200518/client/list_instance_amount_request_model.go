// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceAmountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBeginDate(v string) *ListInstanceAmountRequest
	GetBeginDate() *string
	SetEndDate(v string) *ListInstanceAmountRequest
	GetEndDate() *string
	SetProjectId(v int64) *ListInstanceAmountRequest
	GetProjectId() *int64
}

type ListInstanceAmountRequest struct {
	// The beginning of the time range to query, accurate to the day. Specify the time in the ISO 8601 standard in the yyyy-MM-dd\\"T\\"HH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-06-13T00:00:00+0800
	BeginDate *string `json:"BeginDate,omitempty" xml:"BeginDate,omitempty"`
	// The end of the time range to query, accurate to the day. Specify the time in the ISO 8601 standard in the yyyy-MM-dd\\"T\\"HH:mm:ssZ format. The time must be in UTC.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2021-06-16T00:00:00+0800
	EndDate *string `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// The DataWorks workspace ID. You can log on to the DataWorks console and go to the Work space page to obtain the workspace ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s ListInstanceAmountRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceAmountRequest) GoString() string {
	return s.String()
}

func (s *ListInstanceAmountRequest) GetBeginDate() *string {
	return s.BeginDate
}

func (s *ListInstanceAmountRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *ListInstanceAmountRequest) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListInstanceAmountRequest) SetBeginDate(v string) *ListInstanceAmountRequest {
	s.BeginDate = &v
	return s
}

func (s *ListInstanceAmountRequest) SetEndDate(v string) *ListInstanceAmountRequest {
	s.EndDate = &v
	return s
}

func (s *ListInstanceAmountRequest) SetProjectId(v int64) *ListInstanceAmountRequest {
	s.ProjectId = &v
	return s
}

func (s *ListInstanceAmountRequest) Validate() error {
	return dara.Validate(s)
}
