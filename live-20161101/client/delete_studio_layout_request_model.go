// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStudioLayoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DeleteStudioLayoutRequest
	GetCasterId() *string
	SetLayoutId(v string) *DeleteStudioLayoutRequest
	GetLayoutId() *string
	SetOwnerId(v int64) *DeleteStudioLayoutRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteStudioLayoutRequest
	GetRegionId() *string
}

type DeleteStudioLayoutRequest struct {
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
	// 5c6a2a0d-f228-4a64-af62-20e91b96****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the layout. If the layout was added by calling the [AddStudioLayout](https://help.aliyun.com/document_detail/2848062.html) operation, check the value of the response parameter LayoutId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteStudioLayoutRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStudioLayoutRequest) GoString() string {
	return s.String()
}

func (s *DeleteStudioLayoutRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DeleteStudioLayoutRequest) GetLayoutId() *string {
	return s.LayoutId
}

func (s *DeleteStudioLayoutRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteStudioLayoutRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteStudioLayoutRequest) SetCasterId(v string) *DeleteStudioLayoutRequest {
	s.CasterId = &v
	return s
}

func (s *DeleteStudioLayoutRequest) SetLayoutId(v string) *DeleteStudioLayoutRequest {
	s.LayoutId = &v
	return s
}

func (s *DeleteStudioLayoutRequest) SetOwnerId(v int64) *DeleteStudioLayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteStudioLayoutRequest) SetRegionId(v string) *DeleteStudioLayoutRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStudioLayoutRequest) Validate() error {
	return dara.Validate(s)
}
