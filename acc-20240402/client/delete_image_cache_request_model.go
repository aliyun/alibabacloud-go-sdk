// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteImageCacheRequest
	GetForce() *bool
	SetImageCacheId(v string) *DeleteImageCacheRequest
	GetImageCacheId() *string
	SetRegionId(v string) *DeleteImageCacheRequest
	GetRegionId() *string
}

type DeleteImageCacheRequest struct {
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// imc-bp1dj*****
	ImageCacheId *string `json:"ImageCacheId,omitempty" xml:"ImageCacheId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteImageCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageCacheRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageCacheRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteImageCacheRequest) GetImageCacheId() *string {
	return s.ImageCacheId
}

func (s *DeleteImageCacheRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteImageCacheRequest) SetForce(v bool) *DeleteImageCacheRequest {
	s.Force = &v
	return s
}

func (s *DeleteImageCacheRequest) SetImageCacheId(v string) *DeleteImageCacheRequest {
	s.ImageCacheId = &v
	return s
}

func (s *DeleteImageCacheRequest) SetRegionId(v string) *DeleteImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteImageCacheRequest) Validate() error {
	return dara.Validate(s)
}
