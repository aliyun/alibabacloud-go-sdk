// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoProcessingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *GetVideoProcessingRequest
	GetConfigId() *int64
	SetSiteId(v int64) *GetVideoProcessingRequest
	GetSiteId() *int64
}

type GetVideoProcessingRequest struct {
	// The configuration ID. You can call the [ListVideoProcessings](~~ListVideoProcessings~~) operation to obtain the configuration ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23321557***
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The website ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 23282348***
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s GetVideoProcessingRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVideoProcessingRequest) GoString() string {
	return s.String()
}

func (s *GetVideoProcessingRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *GetVideoProcessingRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *GetVideoProcessingRequest) SetConfigId(v int64) *GetVideoProcessingRequest {
	s.ConfigId = &v
	return s
}

func (s *GetVideoProcessingRequest) SetSiteId(v int64) *GetVideoProcessingRequest {
	s.SiteId = &v
	return s
}

func (s *GetVideoProcessingRequest) Validate() error {
	return dara.Validate(s)
}
