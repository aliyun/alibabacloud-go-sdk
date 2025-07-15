// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v *ListExecutorsRequestFilter) *ListExecutorsRequest
	GetFilter() *ListExecutorsRequestFilter
	SetPageNumber(v int32) *ListExecutorsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListExecutorsRequest
	GetPageSize() *int32
}

type ListExecutorsRequest struct {
	Filter *ListExecutorsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListExecutorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutorsRequest) GetFilter() *ListExecutorsRequestFilter {
	return s.Filter
}

func (s *ListExecutorsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListExecutorsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListExecutorsRequest) SetFilter(v *ListExecutorsRequestFilter) *ListExecutorsRequest {
	s.Filter = v
	return s
}

func (s *ListExecutorsRequest) SetPageNumber(v int32) *ListExecutorsRequest {
	s.PageNumber = &v
	return s
}

func (s *ListExecutorsRequest) SetPageSize(v int32) *ListExecutorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListExecutorsRequest) Validate() error {
	return dara.Validate(s)
}

type ListExecutorsRequestFilter struct {
	ExecutorIds []*string `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty" type:"Repeated"`
	Image       *string   `json:"Image,omitempty" xml:"Image,omitempty"`
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// example:
	//
	// testJob
	JobName *string   `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Status  []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// example:
	//
	// 1703819914
	TimeCreatedAfter *int32 `json:"TimeCreatedAfter,omitempty" xml:"TimeCreatedAfter,omitempty"`
	// example:
	//
	// 1703820113
	TimeCreatedBefore *int32  `json:"TimeCreatedBefore,omitempty" xml:"TimeCreatedBefore,omitempty"`
	VswitchId         *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s ListExecutorsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorsRequestFilter) GoString() string {
	return s.String()
}

func (s *ListExecutorsRequestFilter) GetExecutorIds() []*string {
	return s.ExecutorIds
}

func (s *ListExecutorsRequestFilter) GetImage() *string {
	return s.Image
}

func (s *ListExecutorsRequestFilter) GetIpAddresses() []*string {
	return s.IpAddresses
}

func (s *ListExecutorsRequestFilter) GetJobName() *string {
	return s.JobName
}

func (s *ListExecutorsRequestFilter) GetStatus() []*string {
	return s.Status
}

func (s *ListExecutorsRequestFilter) GetTimeCreatedAfter() *int32 {
	return s.TimeCreatedAfter
}

func (s *ListExecutorsRequestFilter) GetTimeCreatedBefore() *int32 {
	return s.TimeCreatedBefore
}

func (s *ListExecutorsRequestFilter) GetVswitchId() *string {
	return s.VswitchId
}

func (s *ListExecutorsRequestFilter) SetExecutorIds(v []*string) *ListExecutorsRequestFilter {
	s.ExecutorIds = v
	return s
}

func (s *ListExecutorsRequestFilter) SetImage(v string) *ListExecutorsRequestFilter {
	s.Image = &v
	return s
}

func (s *ListExecutorsRequestFilter) SetIpAddresses(v []*string) *ListExecutorsRequestFilter {
	s.IpAddresses = v
	return s
}

func (s *ListExecutorsRequestFilter) SetJobName(v string) *ListExecutorsRequestFilter {
	s.JobName = &v
	return s
}

func (s *ListExecutorsRequestFilter) SetStatus(v []*string) *ListExecutorsRequestFilter {
	s.Status = v
	return s
}

func (s *ListExecutorsRequestFilter) SetTimeCreatedAfter(v int32) *ListExecutorsRequestFilter {
	s.TimeCreatedAfter = &v
	return s
}

func (s *ListExecutorsRequestFilter) SetTimeCreatedBefore(v int32) *ListExecutorsRequestFilter {
	s.TimeCreatedBefore = &v
	return s
}

func (s *ListExecutorsRequestFilter) SetVswitchId(v string) *ListExecutorsRequestFilter {
	s.VswitchId = &v
	return s
}

func (s *ListExecutorsRequestFilter) Validate() error {
	return dara.Validate(s)
}
