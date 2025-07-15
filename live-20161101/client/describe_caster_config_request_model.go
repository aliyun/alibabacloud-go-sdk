// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeCasterConfigRequest
	GetCasterId() *string
	SetOwnerId(v int64) *DescribeCasterConfigRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCasterConfigRequest
	GetRegionId() *string
}

type DescribeCasterConfigRequest struct {
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

func (s DescribeCasterConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterConfigRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterConfigRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCasterConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCasterConfigRequest) SetCasterId(v string) *DescribeCasterConfigRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterConfigRequest) SetOwnerId(v int64) *DescribeCasterConfigRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCasterConfigRequest) SetRegionId(v string) *DescribeCasterConfigRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCasterConfigRequest) Validate() error {
	return dara.Validate(s)
}
