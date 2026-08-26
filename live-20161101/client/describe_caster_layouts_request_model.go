// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterLayoutsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeCasterLayoutsRequest
	GetCasterId() *string
	SetLayoutId(v string) *DescribeCasterLayoutsRequest
	GetLayoutId() *string
	SetOwnerId(v int64) *DescribeCasterLayoutsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCasterLayoutsRequest
	GetRegionId() *string
}

type DescribeCasterLayoutsRequest struct {
	// The ID of the production studio.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value returned in the response.
	//
	// - If you created the production studio in the LIVE console, find the ID on the **Production Studio*	- > **Cloud Production Studio*	- page.
	//
	// > The name of the production studio in the list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The layout ID.
	//
	// If you do not provide a value for LayoutId, all layouts of the production studio are returned.
	//
	// example:
	//
	// 72d2ec7a-4cd7-4a01-974b-7cd53947****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeCasterLayoutsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterLayoutsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterLayoutsRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterLayoutsRequest) GetLayoutId() *string {
	return s.LayoutId
}

func (s *DescribeCasterLayoutsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCasterLayoutsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCasterLayoutsRequest) SetCasterId(v string) *DescribeCasterLayoutsRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterLayoutsRequest) SetLayoutId(v string) *DescribeCasterLayoutsRequest {
	s.LayoutId = &v
	return s
}

func (s *DescribeCasterLayoutsRequest) SetOwnerId(v int64) *DescribeCasterLayoutsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCasterLayoutsRequest) SetRegionId(v string) *DescribeCasterLayoutsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCasterLayoutsRequest) Validate() error {
	return dara.Validate(s)
}
