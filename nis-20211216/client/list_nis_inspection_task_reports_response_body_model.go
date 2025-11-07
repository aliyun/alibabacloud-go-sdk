// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNisInspectionTaskReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionReportList(v []*ListNisInspectionTaskReportsResponseBodyInspectionReportList) *ListNisInspectionTaskReportsResponseBody
	GetInspectionReportList() []*ListNisInspectionTaskReportsResponseBodyInspectionReportList
	SetMaxResults(v int32) *ListNisInspectionTaskReportsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListNisInspectionTaskReportsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListNisInspectionTaskReportsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListNisInspectionTaskReportsResponseBody
	GetTotalCount() *int32
}

type ListNisInspectionTaskReportsResponseBody struct {
	InspectionReportList []*ListNisInspectionTaskReportsResponseBodyInspectionReportList `json:"InspectionReportList,omitempty" xml:"InspectionReportList,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// hKrS+MVXkuOgztXnvdml194Cz/lMNdmr+DEh0th6dVlNEo/F148UPCh2itDku7Qj
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0D213AF9-7B8A-51A8-B411-2D797A1A447B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 34
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListNisInspectionTaskReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNisInspectionTaskReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNisInspectionTaskReportsResponseBody) GetInspectionReportList() []*ListNisInspectionTaskReportsResponseBodyInspectionReportList {
	return s.InspectionReportList
}

func (s *ListNisInspectionTaskReportsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNisInspectionTaskReportsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNisInspectionTaskReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNisInspectionTaskReportsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListNisInspectionTaskReportsResponseBody) SetInspectionReportList(v []*ListNisInspectionTaskReportsResponseBodyInspectionReportList) *ListNisInspectionTaskReportsResponseBody {
	s.InspectionReportList = v
	return s
}

func (s *ListNisInspectionTaskReportsResponseBody) SetMaxResults(v int32) *ListNisInspectionTaskReportsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNisInspectionTaskReportsResponseBody) SetNextToken(v string) *ListNisInspectionTaskReportsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNisInspectionTaskReportsResponseBody) SetRequestId(v string) *ListNisInspectionTaskReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNisInspectionTaskReportsResponseBody) SetTotalCount(v int32) *ListNisInspectionTaskReportsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListNisInspectionTaskReportsResponseBody) Validate() error {
	if s.InspectionReportList != nil {
		for _, item := range s.InspectionReportList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListNisInspectionTaskReportsResponseBodyInspectionReportList struct {
	// example:
	//
	// nir-7c3dd178738a429abe6d
	InspectionReportId *string `json:"InspectionReportId,omitempty" xml:"InspectionReportId,omitempty"`
}

func (s ListNisInspectionTaskReportsResponseBodyInspectionReportList) String() string {
	return dara.Prettify(s)
}

func (s ListNisInspectionTaskReportsResponseBodyInspectionReportList) GoString() string {
	return s.String()
}

func (s *ListNisInspectionTaskReportsResponseBodyInspectionReportList) GetInspectionReportId() *string {
	return s.InspectionReportId
}

func (s *ListNisInspectionTaskReportsResponseBodyInspectionReportList) SetInspectionReportId(v string) *ListNisInspectionTaskReportsResponseBodyInspectionReportList {
	s.InspectionReportId = &v
	return s
}

func (s *ListNisInspectionTaskReportsResponseBodyInspectionReportList) Validate() error {
	return dara.Validate(s)
}
