// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCasterComponentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCasterId(v string) *DescribeCasterComponentsRequest
	GetCasterId() *string
	SetComponentId(v string) *DescribeCasterComponentsRequest
	GetComponentId() *string
	SetOwnerId(v int64) *DescribeCasterComponentsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeCasterComponentsRequest
	GetRegionId() *string
}

type DescribeCasterComponentsRequest struct {
	// The ID of the production studio.
	//
	// - If you created the production studio by calling the [CreateCaster](https://help.aliyun.com/document_detail/2848009.html) operation, obtain the ID from the CasterId parameter in the response.
	//
	// - If you created the production studio in the LIVE console, go to the Cloud Production Studio page by choosing **LIVE Console*	- > **Production Studio*	- > **Cloud Production Studio**.
	//
	// > The name of the production studio in the list on the Cloud Production Studio page is the production studio ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// LIVEPRODUCER_POST-cn-0pp1czt****
	CasterId *string `json:"CasterId,omitempty" xml:"CasterId,omitempty"`
	// The component ID. If you added the component by calling the [AddCasterComponent](https://help.aliyun.com/document_detail/2848030.html) operation, obtain the ID from the ComponentId parameter in the response.
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

func (s DescribeCasterComponentsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCasterComponentsRequest) GoString() string {
	return s.String()
}

func (s *DescribeCasterComponentsRequest) GetCasterId() *string {
	return s.CasterId
}

func (s *DescribeCasterComponentsRequest) GetComponentId() *string {
	return s.ComponentId
}

func (s *DescribeCasterComponentsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCasterComponentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCasterComponentsRequest) SetCasterId(v string) *DescribeCasterComponentsRequest {
	s.CasterId = &v
	return s
}

func (s *DescribeCasterComponentsRequest) SetComponentId(v string) *DescribeCasterComponentsRequest {
	s.ComponentId = &v
	return s
}

func (s *DescribeCasterComponentsRequest) SetOwnerId(v int64) *DescribeCasterComponentsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCasterComponentsRequest) SetRegionId(v string) *DescribeCasterComponentsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCasterComponentsRequest) Validate() error {
	return dara.Validate(s)
}
