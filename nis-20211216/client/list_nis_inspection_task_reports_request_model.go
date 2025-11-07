// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNisInspectionTaskReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInspectionTaskId(v string) *ListNisInspectionTaskReportsRequest
	GetInspectionTaskId() *string
	SetMaxResults(v int32) *ListNisInspectionTaskReportsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNisInspectionTaskReportsRequest
	GetNextToken() *string
}

type ListNisInspectionTaskReportsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ni-8svmpe0yso2bhzr7fh79
	InspectionTaskId *string `json:"InspectionTaskId,omitempty" xml:"InspectionTaskId,omitempty"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// hKrS+MVXkuOgztXnvdml194Cz/lMNdmr+DEh0th6dVlNEo/F148UPCh2itDku7Qj
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
}

func (s ListNisInspectionTaskReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNisInspectionTaskReportsRequest) GoString() string {
	return s.String()
}

func (s *ListNisInspectionTaskReportsRequest) GetInspectionTaskId() *string {
	return s.InspectionTaskId
}

func (s *ListNisInspectionTaskReportsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNisInspectionTaskReportsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNisInspectionTaskReportsRequest) SetInspectionTaskId(v string) *ListNisInspectionTaskReportsRequest {
	s.InspectionTaskId = &v
	return s
}

func (s *ListNisInspectionTaskReportsRequest) SetMaxResults(v int32) *ListNisInspectionTaskReportsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNisInspectionTaskReportsRequest) SetNextToken(v string) *ListNisInspectionTaskReportsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNisInspectionTaskReportsRequest) Validate() error {
	return dara.Validate(s)
}
