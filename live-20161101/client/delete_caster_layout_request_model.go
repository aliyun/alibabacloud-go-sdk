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
	// 	- If the production studio was created by calling the [CreateCaster](https://help.aliyun.com/document_detail/69338.html) operation, check the value of the response parameter CasterId to obtain the ID.
	//
	// 	- If the production studio was created by using the ApsaraVideo Live console, obtain the ID on the **Production Studio Management*	- page. To go to the page, log on to the **ApsaraVideo Live console*	- and click **Production Studios*	- in the left-side navigation pane.
	//
	// >  You can find the ID of the production studio in the Instance ID/Name column.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The ID of the layout. If the layout was added by calling the [AddCasterLayout](https://help.aliyun.com/document_detail/60249.html) operation, check the value of the response parameter LayoutId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21926b36-7dd2-4fde-ae25-51b5bc8e****
	LayoutId *string `json:"LayoutId,omitempty" xml:"LayoutId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
