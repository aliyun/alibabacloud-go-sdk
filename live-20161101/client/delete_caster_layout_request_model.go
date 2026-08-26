// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterLayoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DeleteCasterLayoutRequest
	GetCasterId() *string
	SetLayoutId(v string) *DeleteCasterLayoutRequest
	GetLayoutId() *string
	SetOwnerId(v int64) *DeleteCasterLayoutRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCasterLayoutRequest
	GetRegionId() *string
}

type DeleteCasterLayoutRequest struct {
	// The ID of the production studio.
	//
	// - If you call the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation to create a production studio, obtain the CasterId from the response.
	//
	// - If you create a production studio in the ApsaraVideo Live console, go to the **Production Studio*	- > **Cloud Production Studio*	- page to view the ID.
	//
	// > The ID of the production studio is its name in the list on the Cloud Production Studio page.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The layout ID. If you call the [AddCasterLayout](https://help.aliyun.com/document_detail/2848025.html) operation to add a layout to the production studio, obtain the LayoutId from the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21926b36-7dd2-4fde-ae25-51b5bc8e****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCasterLayoutRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterLayoutRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterLayoutRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DeleteCasterLayoutRequest) GetLayoutId() *string {
	return s.LayoutId
}

func (s *DeleteCasterLayoutRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCasterLayoutRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCasterLayoutRequest) SetCasterId(v string) *DeleteCasterLayoutRequest {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterLayoutRequest) SetLayoutId(v string) *DeleteCasterLayoutRequest {
	s.LayoutId = &v
	return s
}

func (s *DeleteCasterLayoutRequest) SetOwnerId(v int64) *DeleteCasterLayoutRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCasterLayoutRequest) SetRegionId(v string) *DeleteCasterLayoutRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCasterLayoutRequest) Validate() error {
	return dara.Validate(s)
}
