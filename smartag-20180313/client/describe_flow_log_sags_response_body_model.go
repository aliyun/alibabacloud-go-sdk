// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFlowLogSagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeFlowLogSagsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeFlowLogSagsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeFlowLogSagsResponseBody
	GetRequestId() *string
	SetSags(v *DescribeFlowLogSagsResponseBodySags) *DescribeFlowLogSagsResponseBody
	GetSags() *DescribeFlowLogSagsResponseBodySags
	SetTotalCount(v int32) *DescribeFlowLogSagsResponseBody
	GetTotalCount() *int32
}

type DescribeFlowLogSagsResponseBody struct {
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries on the current page.
	//
	// example:
	//
	// 4
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 8D945945-85F2-4BD7-A144-7DC0E8B5A0DC
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sags      *DescribeFlowLogSagsResponseBodySags `json:"Sags,omitempty" xml:"Sags,omitempty" type:"Struct"`
	// The total number of entries.
	//
	// example:
	//
	// 35
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeFlowLogSagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogSagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogSagsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeFlowLogSagsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeFlowLogSagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFlowLogSagsResponseBody) GetSags() *DescribeFlowLogSagsResponseBodySags {
	return s.Sags
}

func (s *DescribeFlowLogSagsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeFlowLogSagsResponseBody) SetPageNumber(v int32) *DescribeFlowLogSagsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeFlowLogSagsResponseBody) SetPageSize(v int32) *DescribeFlowLogSagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeFlowLogSagsResponseBody) SetRequestId(v string) *DescribeFlowLogSagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFlowLogSagsResponseBody) SetSags(v *DescribeFlowLogSagsResponseBodySags) *DescribeFlowLogSagsResponseBody {
	s.Sags = v
	return s
}

func (s *DescribeFlowLogSagsResponseBody) SetTotalCount(v int32) *DescribeFlowLogSagsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeFlowLogSagsResponseBody) Validate() error {
	if s.Sags != nil {
		if err := s.Sags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeFlowLogSagsResponseBodySags struct {
	Sag []*DescribeFlowLogSagsResponseBodySagsSag `json:"Sag,omitempty" xml:"Sag,omitempty" type:"Repeated"`
}

func (s DescribeFlowLogSagsResponseBodySags) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogSagsResponseBodySags) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogSagsResponseBodySags) GetSag() []*DescribeFlowLogSagsResponseBodySagsSag {
	return s.Sag
}

func (s *DescribeFlowLogSagsResponseBodySags) SetSag(v []*DescribeFlowLogSagsResponseBodySagsSag) *DescribeFlowLogSagsResponseBodySags {
	s.Sag = v
	return s
}

func (s *DescribeFlowLogSagsResponseBodySags) Validate() error {
	if s.Sag != nil {
		for _, item := range s.Sag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFlowLogSagsResponseBodySagsSag struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SmartAGId   *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeFlowLogSagsResponseBodySagsSag) String() string {
	return dara.Prettify(s)
}

func (s DescribeFlowLogSagsResponseBodySagsSag) GoString() string {
	return s.String()
}

func (s *DescribeFlowLogSagsResponseBodySagsSag) GetDescription() *string {
	return s.Description
}

func (s *DescribeFlowLogSagsResponseBodySagsSag) GetName() *string {
	return s.Name
}

func (s *DescribeFlowLogSagsResponseBodySagsSag) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeFlowLogSagsResponseBodySagsSag) SetDescription(v string) *DescribeFlowLogSagsResponseBodySagsSag {
	s.Description = &v
	return s
}

func (s *DescribeFlowLogSagsResponseBodySagsSag) SetName(v string) *DescribeFlowLogSagsResponseBodySagsSag {
	s.Name = &v
	return s
}

func (s *DescribeFlowLogSagsResponseBodySagsSag) SetSmartAGId(v string) *DescribeFlowLogSagsResponseBodySagsSag {
	s.SmartAGId = &v
	return s
}

func (s *DescribeFlowLogSagsResponseBodySagsSag) Validate() error {
	return dara.Validate(s)
}
