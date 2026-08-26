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
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value returned in the response.
	//
	// - If you created the production studio in the ApsaraVideo Live console, choose **ApsaraVideo Live console*	- > **Production Studio*	- > **Cloud Production Studio*	- to view the ID.
	//
	// > The name of the production studio in the list is its ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5c6a2a0d-f228-4a64-af62-20e91b96****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The layout ID. If you added the layout to the production studio by calling the [AddStudioLayout](https://help.aliyun.com/document_detail/2848062.html) operation, use the LayoutId value returned in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
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
