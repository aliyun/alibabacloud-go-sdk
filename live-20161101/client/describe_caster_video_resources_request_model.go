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
