// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOpsNoticesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*ListOpsNoticesRequestFilter) *ListOpsNoticesRequest
	GetFilter() []*ListOpsNoticesRequestFilter
	SetMaxResults(v int32) *ListOpsNoticesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListOpsNoticesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListOpsNoticesRequest
	GetRegionId() *string
}

type ListOpsNoticesRequest struct {
	Filter []*ListOpsNoticesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAfu+XtuBE55iRLHEYYuojI4=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListOpsNoticesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListOpsNoticesRequest) GoString() string {
	return s.String()
}

func (s *ListOpsNoticesRequest) GetFilter() []*ListOpsNoticesRequestFilter {
	return s.Filter
}

func (s *ListOpsNoticesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListOpsNoticesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListOpsNoticesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListOpsNoticesRequest) SetFilter(v []*ListOpsNoticesRequestFilter) *ListOpsNoticesRequest {
	s.Filter = v
	return s
}

func (s *ListOpsNoticesRequest) SetMaxResults(v int32) *ListOpsNoticesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListOpsNoticesRequest) SetNextToken(v string) *ListOpsNoticesRequest {
	s.NextToken = &v
	return s
}

func (s *ListOpsNoticesRequest) SetRegionId(v string) *ListOpsNoticesRequest {
	s.RegionId = &v
	return s
}

func (s *ListOpsNoticesRequest) Validate() error {
	return dara.Validate(s)
}

type ListOpsNoticesRequestFilter struct {
	// example:
	//
	// ServiceId
	Name  *string   `json:"Name,omitempty" xml:"Name,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s ListOpsNoticesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListOpsNoticesRequestFilter) GoString() string {
	return s.String()
}

func (s *ListOpsNoticesRequestFilter) GetName() *string {
	return s.Name
}

func (s *ListOpsNoticesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *ListOpsNoticesRequestFilter) SetName(v string) *ListOpsNoticesRequestFilter {
	s.Name = &v
	return s
}

func (s *ListOpsNoticesRequestFilter) SetValue(v []*string) *ListOpsNoticesRequestFilter {
	s.Value = v
	return s
}

func (s *ListOpsNoticesRequestFilter) Validate() error {
	return dara.Validate(s)
}
