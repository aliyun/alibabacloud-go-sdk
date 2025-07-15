// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUserTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DescribeLiveUserTagsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeLiveUserTagsRequest
	GetRegionId() *string
}

type DescribeLiveUserTagsRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeLiveUserTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUserTagsRequest) GoString() string {
	return s.String()
}

func (s *DescribeLiveUserTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLiveUserTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLiveUserTagsRequest) SetOwnerId(v int64) *DescribeLiveUserTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLiveUserTagsRequest) SetRegionId(v string) *DescribeLiveUserTagsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLiveUserTagsRequest) Validate() error {
	return dara.Validate(s)
}
