// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSaslUsersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeSaslUsersRequest
	GetInstanceId() *string
	SetRegionId(v string) *DescribeSaslUsersRequest
	GetRegionId() *string
}

type DescribeSaslUsersRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// alikafka_pre-cn-v0h1cng0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the region.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeSaslUsersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSaslUsersRequest) GoString() string {
	return s.String()
}

func (s *DescribeSaslUsersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSaslUsersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSaslUsersRequest) SetInstanceId(v string) *DescribeSaslUsersRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSaslUsersRequest) SetRegionId(v string) *DescribeSaslUsersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSaslUsersRequest) Validate() error {
	return dara.Validate(s)
}
