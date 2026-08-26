// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeShowListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeShowListRequest
	GetCasterId() *string
	SetOwnerId(v int64) *DescribeShowListRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeShowListRequest
	GetRegionId() *string
}

type DescribeShowListRequest struct {
	// The ID of the production studio.
	//
	// - If you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value that is returned in the response.
	//
	// - If you create a production studio in the LIVE console, choose **LIVE Console*	- > **Production Studio*	- > **Cloud Production Studio*	- to find the name of the production studio.
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

func (s DescribeShowListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeShowListRequest) GoString() string {
	return s.String()
}

func (s *DescribeShowListRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeShowListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeShowListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeShowListRequest) SetCasterId(v string) *DescribeShowListRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeShowListRequest) SetOwnerId(v int64) *DescribeShowListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeShowListRequest) SetRegionId(v string) *DescribeShowListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeShowListRequest) Validate() error {
	return dara.Validate(s)
}
