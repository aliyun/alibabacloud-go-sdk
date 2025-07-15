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
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
