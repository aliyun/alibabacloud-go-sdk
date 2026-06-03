// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSyncInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeSyncInfoRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeSyncInfoRequest
	GetRegionId() *string
}

type DescribeSyncInfoRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	RegionId   *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSyncInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSyncInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSyncInfoRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSyncInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSyncInfoRequest) SetInstanceId(v string) *DescribeSyncInfoRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSyncInfoRequest) SetRegionId(v string) *DescribeSyncInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSyncInfoRequest) Validate() error {
	return dara.Validate(s)
}
