// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DeleteCasterComponentRequest
	GetCasterId() *string
	SetComponentId(v string) *DeleteCasterComponentRequest
	GetComponentId() *string
	SetOwnerId(v int64) *DeleteCasterComponentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCasterComponentRequest
	GetRegionId() *string
}

type DeleteCasterComponentRequest struct {
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
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The component ID. If the component was added by calling the [AddCasterComponent](https://help.aliyun.com/document_detail/2848030.html) operation, check the value of the response parameter ComponentId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21926b36-7dd2-4fde-ae25-51b5bc8e****
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId    *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCasterComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterComponentRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterComponentRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DeleteCasterComponentRequest) GetComponentId() *string {
	return s.ComponentId
}

func (s *DeleteCasterComponentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCasterComponentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCasterComponentRequest) SetCasterId(v string) *DeleteCasterComponentRequest {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterComponentRequest) SetComponentId(v string) *DeleteCasterComponentRequest {
	s.ComponentId = &v
	return s
}

func (s *DeleteCasterComponentRequest) SetOwnerId(v int64) *DeleteCasterComponentRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCasterComponentRequest) SetRegionId(v string) *DeleteCasterComponentRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCasterComponentRequest) Validate() error {
	return dara.Validate(s)
}
