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
	// The filter conditions for querying executors.
	Filter *ListExecutorsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Struct"`
	// The page number. Starting value: 1. Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page in the paging Settings for a paged query. Default value: 50. Maximum value: 100.
	//
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
	if s.Filter != nil {
		if err := s.Filter.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExecutorsRequestFilter struct {
	// The list of executor IDs. A maximum of 100 IDs are supported.
	ExecutorIds []*string `json:"ExecutorIds,omitempty" xml:"ExecutorIds,omitempty" type:"Repeated"`
	// The executor image.
	//
	// example:
	//
	// m-f8z0dfa96luxxxxx
	Image *string `json:"Image,omitempty" xml:"Image,omitempty"`
	// The list of internal IP addresses. A maximum of 100 IP addresses are supported.
	IpAddresses []*string `json:"IpAddresses,omitempty" xml:"IpAddresses,omitempty" type:"Repeated"`
	// The job name. Fuzzy match is supported.
	//
	// example:
	//
	// testJob
	JobName *string `json:"JobName,omitempty" xml:"JobName,omitempty"`
	// The job template ID.
	//
	// example:
	//
	// jt-xxxx
	JobTemplateId *string `json:"JobTemplateId,omitempty" xml:"JobTemplateId,omitempty"`
	// The resource pool name.
	//
	// example:
	//
	// MyPool
	PoolName *string `json:"PoolName,omitempty" xml:"PoolName,omitempty"`
	// The list of executor statuses.
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The time after which the job was submitted. The time is converted to a UNIX timestamp based on the time zone of the region. For regions in the Chinese mainland, the time zone is UTC+8.
	//
	// example:
	//
	// 1703819914
	TimeCreatedAfter *int32 `json:"TimeCreatedAfter,omitempty" xml:"TimeCreatedAfter,omitempty"`
	// The time before which the job was submitted. The time is converted to a UNIX timestamp based on the time zone of the region. For regions in the Chinese mainland, the time zone is UTC+8.
	//
	// example:
	//
	// 1703820113
	TimeCreatedBefore *int32 `json:"TimeCreatedBefore,omitempty" xml:"TimeCreatedBefore,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-xxx
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-xxx
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
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

func (s *ListExecutorsRequestFilter) GetJobTemplateId() *string {
	return s.JobTemplateId
}

func (s *ListExecutorsRequestFilter) GetPoolName() *string {
	return s.PoolName
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

func (s *ListExecutorsRequestFilter) GetVpcId() *string {
	return s.VpcId
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

func (s *ListExecutorsRequestFilter) SetJobTemplateId(v string) *ListExecutorsRequestFilter {
	s.JobTemplateId = &v
	return s
}

func (s *ListExecutorsRequestFilter) SetPoolName(v string) *ListExecutorsRequestFilter {
	s.PoolName = &v
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

func (s *ListExecutorsRequestFilter) SetVpcId(v string) *ListExecutorsRequestFilter {
	s.VpcId = &v
	return s
}

func (s *ListExecutorsRequestFilter) SetVswitchId(v string) *ListExecutorsRequestFilter {
	s.VswitchId = &v
	return s
}

func (s *ListExecutorsRequestFilter) Validate() error {
	return dara.Validate(s)
}
