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
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// >  You can find the ID of the production studio in the Instance ID/Name column.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the layout.
	//
	// If you do not specify a layout ID, all layouts of the production studio are queried.
	//
	// example:
	//
	// 72d2ec7a-4cd7-4a01-974b-7cd53947****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
