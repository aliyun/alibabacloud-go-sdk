// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteResourceShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteResourceShrinkRequest
	GetClientToken() *string
	SetFilterShrink(v string) *DeleteResourceShrinkRequest
	GetFilterShrink() *string
	SetRegionId(v string) *DeleteResourceShrinkRequest
	GetRegionId() *string
}

type DeleteResourceShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. If a cloud service supports idempotence, the parameter takes effect.
	//
	// example:
	//
	// 1e810dfe1468721d0664a49b9d9f74f4
	ClientToken  *string `json:"clientToken,omitempty" xml:"clientToken,omitempty"`
	FilterShrink *string `json:"filter,omitempty" xml:"filter,omitempty"`
	// The region. This parameter is required if a cloud service is a regionalized.
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
}

func (s DeleteResourceShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteResourceShrinkRequest) GoString() string {
	return s.String()
}

func (s *DeleteResourceShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteResourceShrinkRequest) GetFilterShrink() *string {
	return s.FilterShrink
}

func (s *DeleteResourceShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteResourceShrinkRequest) SetClientToken(v string) *DeleteResourceShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteResourceShrinkRequest) SetFilterShrink(v string) *DeleteResourceShrinkRequest {
	s.FilterShrink = &v
	return s
}

func (s *DeleteResourceShrinkRequest) SetRegionId(v string) *DeleteResourceShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteResourceShrinkRequest) Validate() error {
	return dara.Validate(s)
}
