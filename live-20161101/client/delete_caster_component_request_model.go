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
	// - If you create a production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, use the CasterId value that is returned in the response.
	//
	// - If you create a production studio in the ApsaraVideo Live console, find the ID on the **Cloud Production Studio*	- page. In the ApsaraVideo Live console, choose **Production Studio*	- > **Cloud Production Studio**.
	//
	// > The name of the production studio in the list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// a2b8e671-2fe5-4642-a2ec-bf93880e****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The component ID. If you add a component to the production studio by calling the [AddCasterComponent](https://help.aliyun.com/document_detail/2848030.html) operation, use the ComponentId value that is returned in the response.
	//
	// This parameter is required.
	//
	// example:
	//
	// 21926b36-7dd2-4fde-ae25-51b5bc8e****
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
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
