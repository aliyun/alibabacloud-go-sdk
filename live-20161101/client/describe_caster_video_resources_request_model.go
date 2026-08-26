// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterVideoResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeCasterVideoResourcesRequest
	GetCasterId() *string
	SetOwnerId(v int64) *DescribeCasterVideoResourcesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCasterVideoResourcesRequest
	GetRegionId() *string
}

type DescribeCasterVideoResourcesRequest struct {
	// The ID of the production studio.
	//
	// - If you create the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the value of CasterId that is returned.
	//
	// - If you create the production studio in the LIVE console, go to the Cloud Production Studio page to view the ID. To go to the page, choose **LIVE Console*	- > **Production Studio*	- > **Cloud Production Studio**.
	//
	// > The name of the production studio in the list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCasterVideoResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterVideoResourcesRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterVideoResourcesRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterVideoResourcesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCasterVideoResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCasterVideoResourcesRequest) SetCasterId(v string) *DescribeCasterVideoResourcesRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterVideoResourcesRequest) SetOwnerId(v int64) *DescribeCasterVideoResourcesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCasterVideoResourcesRequest) SetRegionId(v string) *DescribeCasterVideoResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCasterVideoResourcesRequest) Validate() error {
	return dara.Validate(s)
}
