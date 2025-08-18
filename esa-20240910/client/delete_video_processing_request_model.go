// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVideoProcessingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteVideoProcessingRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteVideoProcessingRequest
	GetSiteId() *int64
}

type DeleteVideoProcessingRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 35281609698****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1234567890123
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteVideoProcessingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVideoProcessingRequest) GoString() string {
	return s.String()
}

func (s *DeleteVideoProcessingRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteVideoProcessingRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteVideoProcessingRequest) SetConfigId(v int64) *DeleteVideoProcessingRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteVideoProcessingRequest) SetSiteId(v int64) *DeleteVideoProcessingRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteVideoProcessingRequest) Validate() error {
	return dara.Validate(s)
}
