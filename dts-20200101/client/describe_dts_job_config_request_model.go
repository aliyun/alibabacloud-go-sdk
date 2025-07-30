// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDtsJobConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDtsJobId(v string) *DescribeDtsJobConfigRequest
	GetDtsJobId() *string
	SetForAcceleration(v string) *DescribeDtsJobConfigRequest
	GetForAcceleration() *string
	SetModule(v string) *DescribeDtsJobConfigRequest
	GetModule() *string
	SetOwnerId(v string) *DescribeDtsJobConfigRequest
	GetOwnerId() *string
	SetRegionId(v string) *DescribeDtsJobConfigRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDtsJobConfigRequest
	GetResourceGroupId() *string
}

type DescribeDtsJobConfigRequest struct {
	DtsJobId        *string `json:"DtsJobId,omitempty" xml:"DtsJobId,omitempty"`
	ForAcceleration *string `json:"ForAcceleration,omitempty" xml:"ForAcceleration,omitempty"`
	Module          *string `json:"Module,omitempty" xml:"Module,omitempty"`
	OwnerId         *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s DescribeDtsJobConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDtsJobConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeDtsJobConfigRequest) GetDtsJobId() *string {
	return s.DtsJobId
}

func (s *DescribeDtsJobConfigRequest) GetForAcceleration() *string {
	return s.ForAcceleration
}

func (s *DescribeDtsJobConfigRequest) GetModule() *string {
	return s.Module
}

func (s *DescribeDtsJobConfigRequest) GetOwnerId() *string {
	return s.OwnerId
}

func (s *DescribeDtsJobConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDtsJobConfigRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDtsJobConfigRequest) SetDtsJobId(v string) *DescribeDtsJobConfigRequest {
	s.DtsJobId = &v
	return s
}

func (s *DescribeDtsJobConfigRequest) SetForAcceleration(v string) *DescribeDtsJobConfigRequest {
	s.ForAcceleration = &v
	return s
}

func (s *DescribeDtsJobConfigRequest) SetModule(v string) *DescribeDtsJobConfigRequest {
	s.Module = &v
	return s
}

func (s *DescribeDtsJobConfigRequest) SetOwnerId(v string) *DescribeDtsJobConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDtsJobConfigRequest) SetRegionId(v string) *DescribeDtsJobConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDtsJobConfigRequest) SetResourceGroupId(v string) *DescribeDtsJobConfigRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDtsJobConfigRequest) Validate() error {
	return dara.Validate(s)
}
