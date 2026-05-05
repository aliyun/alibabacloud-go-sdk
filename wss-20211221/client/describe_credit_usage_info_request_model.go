// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCreditUsageInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *DescribeCreditUsageInfoRequest
	GetBizType() *string
	SetInstanceIds(v []*string) *DescribeCreditUsageInfoRequest
	GetInstanceIds() []*string
	SetUsageType(v string) *DescribeCreditUsageInfoRequest
	GetUsageType() *string
}

type DescribeCreditUsageInfoRequest struct {
	BizType     *string   `json:"BizType,omitempty" xml:"BizType,omitempty"`
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// example:
	//
	// User
	UsageType *string `json:"UsageType,omitempty" xml:"UsageType,omitempty"`
}

func (s DescribeCreditUsageInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCreditUsageInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeCreditUsageInfoRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeCreditUsageInfoRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *DescribeCreditUsageInfoRequest) GetUsageType() *string {
	return s.UsageType
}

func (s *DescribeCreditUsageInfoRequest) SetBizType(v string) *DescribeCreditUsageInfoRequest {
	s.BizType = &v
	return s
}

func (s *DescribeCreditUsageInfoRequest) SetInstanceIds(v []*string) *DescribeCreditUsageInfoRequest {
	s.InstanceIds = v
	return s
}

func (s *DescribeCreditUsageInfoRequest) SetUsageType(v string) *DescribeCreditUsageInfoRequest {
	s.UsageType = &v
	return s
}

func (s *DescribeCreditUsageInfoRequest) Validate() error {
	return dara.Validate(s)
}
