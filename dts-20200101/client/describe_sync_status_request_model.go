// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDirection(v string) *DescribeSyncStatusRequest
	GetDirection() *string
	SetDtsInstanceId(v string) *DescribeSyncStatusRequest
	GetDtsInstanceId() *string
	SetDtsJobId(v string) *DescribeSyncStatusRequest
	GetDtsJobId() *string
	SetRegionId(v string) *DescribeSyncStatusRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSyncStatusRequest
	GetResourceGroupId() *string
}

type DescribeSyncStatusRequest struct {
	Direction       *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	DtsInstanceId   *string `json:"DtsInstanceId,omitempty" xml:"DtsInstanceId,omitempty"`
	DtsJobId        *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeSyncStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeSyncStatusRequest) GetDirection() *string {
	return s.Direction
}

func (s *DescribeSyncStatusRequest) GetDtsInstanceId() *string {
	return s.DtsInstanceId
}

func (s *DescribeSyncStatusRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeSyncStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSyncStatusRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSyncStatusRequest) SetDirection(v string) *DescribeSyncStatusRequest {
	s.Direction = &v
	return s
}

func (s *DescribeSyncStatusRequest) SetDtsInstanceId(v string) *DescribeSyncStatusRequest {
	s.DtsInstanceId = &v
	return s
}

func (s *DescribeSyncStatusRequest) SetDtsJobId(v string) *DescribeSyncStatusRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeSyncStatusRequest) SetRegionId(v string) *DescribeSyncStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSyncStatusRequest) SetResourceGroupId(v string) *DescribeSyncStatusRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSyncStatusRequest) Validate() error {
	return dara.Validate(s)
}
