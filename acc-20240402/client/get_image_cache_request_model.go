// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageCacheRequest interface {
	dara.Model
	String() string
	GoString() string
	SetImageCacheId(v string) *GetImageCacheRequest
	GetImageCacheId() *string
	SetRegionId(v string) *GetImageCacheRequest
	GetRegionId() *string
}

type GetImageCacheRequest struct {
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

func (s GetImageCacheRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageCacheRequest) GoString() string {
	return s.String()
}

func (s *GetImageCacheRequest) GetImageCacheId() *string {
	return s.ImageCacheId
}

func (s *GetImageCacheRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetImageCacheRequest) SetImageCacheId(v string) *GetImageCacheRequest {
	s.ImageCacheId = &v
	return s
}

func (s *GetImageCacheRequest) SetRegionId(v string) *GetImageCacheRequest {
	s.RegionId = &v
	return s
}

func (s *GetImageCacheRequest) Validate() error {
	return dara.Validate(s)
}
