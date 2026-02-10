// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPlayChoosenShowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *PlayChoosenShowRequest
	GetCasterId() *string
	SetOwnerId(v int64) *PlayChoosenShowRequest
	GetOwnerId() *int64
	SetRegionId(v string) *PlayChoosenShowRequest
	GetRegionId() *string
	SetShowId(v string) *PlayChoosenShowRequest
	GetShowId() *string
}

type PlayChoosenShowRequest struct {
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
	// The ID of the region.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the episode that you want to switch to.
	//
	// >  You can call the [AddShowIntoShowList](https://help.aliyun.com/document_detail/2848051.html) or [DescribeShowList](https://help.aliyun.com/document_detail/2848054.html) operation and check the value of the response parameter ShowId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	ShowId *string `json:"ShowId,omitempty" xml:"ShowId,omitempty"`
}

func (s PlayChoosenShowRequest) String() string {
	return dara.Prettify(s)
}

func (s PlayChoosenShowRequest) GoString() string {
	return s.String()
}

func (s *PlayChoosenShowRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *PlayChoosenShowRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *PlayChoosenShowRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *PlayChoosenShowRequest) GetShowId() *string {
	return s.ShowId
}

func (s *PlayChoosenShowRequest) SetCasterId(v string) *PlayChoosenShowRequest {
	s.CasterId = &v
	return s
}

func (s *PlayChoosenShowRequest) SetOwnerId(v int64) *PlayChoosenShowRequest {
	s.OwnerId = &v
	return s
}

func (s *PlayChoosenShowRequest) SetRegionId(v string) *PlayChoosenShowRequest {
	s.RegionId = &v
	return s
}

func (s *PlayChoosenShowRequest) SetShowId(v string) *PlayChoosenShowRequest {
	s.ShowId = &v
	return s
}

func (s *PlayChoosenShowRequest) Validate() error {
	return dara.Validate(s)
}
