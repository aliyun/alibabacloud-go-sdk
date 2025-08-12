// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorResourceQuotaAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeMonitorResourceQuotaAttributeRequest
	GetRegionId() *string
	SetShowUsed(v bool) *DescribeMonitorResourceQuotaAttributeRequest
	GetShowUsed() *bool
}

type DescribeMonitorResourceQuotaAttributeRequest struct {
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to return information about used quotas. Valid values:
	//
	// 	- true (default): yes
	//
	// 	- false: no
	//
	// example:
	//
	// true
	ShowUsed *bool `json:"ShowUsed,omitempty" xml:"ShowUsed,omitempty"`
}

func (s DescribeMonitorResourceQuotaAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorResourceQuotaAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeMonitorResourceQuotaAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeMonitorResourceQuotaAttributeRequest) GetShowUsed() *bool {
	return s.ShowUsed
}

func (s *DescribeMonitorResourceQuotaAttributeRequest) SetRegionId(v string) *DescribeMonitorResourceQuotaAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeRequest) SetShowUsed(v bool) *DescribeMonitorResourceQuotaAttributeRequest {
	s.ShowUsed = &v
	return s
}

func (s *DescribeMonitorResourceQuotaAttributeRequest) Validate() error {
	return dara.Validate(s)
}
