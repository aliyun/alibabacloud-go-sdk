// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageTransformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetImageTransformRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetImageTransformRequest
	GetSiteId() *int64
}

type GetImageTransformRequest struct {
	// This parameter is required.
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetImageTransformRequest) String() string {
	return dara.Prettify(s)
}

func (s GetImageTransformRequest) GoString() string {
	return s.String()
}

func (s *GetImageTransformRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetImageTransformRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetImageTransformRequest) SetConfigId(v int64) *GetImageTransformRequest {
	s.ConfigId = &v
	return s
}

func (s *GetImageTransformRequest) SetSiteId(v int64) *GetImageTransformRequest {
	s.SiteId = &v
	return s
}

func (s *GetImageTransformRequest) Validate() error {
	return dara.Validate(s)
}
