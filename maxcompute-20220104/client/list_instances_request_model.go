// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *ListInstancesRequest
	GetEndDate() *int64
	SetStartDate(v int64) *ListInstancesRequest
	GetStartDate() *int64
}

type ListInstancesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1759975856382
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1775644926203
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s ListInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesRequest) GoString() string {
	return s.String()
}

func (s *ListInstancesRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *ListInstancesRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *ListInstancesRequest) SetEndDate(v int64) *ListInstancesRequest {
	s.EndDate = &v
	return s
}

func (s *ListInstancesRequest) SetStartDate(v int64) *ListInstancesRequest {
	s.StartDate = &v
	return s
}

func (s *ListInstancesRequest) Validate() error {
	return dara.Validate(s)
}
