// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterVideoResourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DeleteCasterVideoResourceRequest
	GetCasterId() *string
	SetOwnerId(v int64) *DeleteCasterVideoResourceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCasterVideoResourceRequest
	GetRegionId() *string
	SetResourceId(v string) *DeleteCasterVideoResourceRequest
	GetResourceId() *string
}

type DeleteCasterVideoResourceRequest struct {
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
	// 80787064-1c94-4dc1-85ce-9409960a****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource ID. If the input source was added by calling the AddCasterVideoResource operation, check the value of the response parameter ResourceId to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 05ab713c-676e-49c0-96ce-cc408da1****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s DeleteCasterVideoResourceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterVideoResourceRequest) GoString() string {
	return s.String()
}

func (s *DeleteCasterVideoResourceRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DeleteCasterVideoResourceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCasterVideoResourceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCasterVideoResourceRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *DeleteCasterVideoResourceRequest) SetCasterId(v string) *DeleteCasterVideoResourceRequest {
	s.CasterId = &v
	return s
}

func (s *DeleteCasterVideoResourceRequest) SetOwnerId(v int64) *DeleteCasterVideoResourceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCasterVideoResourceRequest) SetRegionId(v string) *DeleteCasterVideoResourceRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCasterVideoResourceRequest) SetResourceId(v string) *DeleteCasterVideoResourceRequest {
	s.ResourceId = &v
	return s
}

func (s *DeleteCasterVideoResourceRequest) Validate() error {
	return dara.Validate(s)
}
