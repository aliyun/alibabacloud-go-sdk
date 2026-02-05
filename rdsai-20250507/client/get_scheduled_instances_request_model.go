// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduledInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *GetScheduledInstancesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *GetScheduledInstancesRequest
	GetPageSize() *int64
	SetScheduledId(v string) *GetScheduledInstancesRequest
	GetScheduledId() *string
}

type GetScheduledInstancesRequest struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 847268a4-196f-416b-aa12-bfe0c115****
	ScheduledId *string `json:"ScheduledId,omitempty" xml:"ScheduledId,omitempty"`
}

func (s GetScheduledInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetScheduledInstancesRequest) GoString() string {
	return s.String()
}

func (s *GetScheduledInstancesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *GetScheduledInstancesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetScheduledInstancesRequest) GetScheduledId() *string {
	return s.ScheduledId
}

func (s *GetScheduledInstancesRequest) SetPageNumber(v int64) *GetScheduledInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *GetScheduledInstancesRequest) SetPageSize(v int64) *GetScheduledInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *GetScheduledInstancesRequest) SetScheduledId(v string) *GetScheduledInstancesRequest {
	s.ScheduledId = &v
	return s
}

func (s *GetScheduledInstancesRequest) Validate() error {
	return dara.Validate(s)
}
