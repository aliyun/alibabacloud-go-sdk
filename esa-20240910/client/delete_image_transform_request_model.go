// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageTransformRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *DeleteImageTransformRequest
	GetConfigId() *int64
	SetSiteId(v int64) *DeleteImageTransformRequest
	GetSiteId() *int64
}

type DeleteImageTransformRequest struct {
	// Configuration ID. It can be obtained by calling the [ListImageTransforms](https://help.aliyun.com/document_detail/2869056.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Site ID, which can be obtained by calling the [ListSites](https://help.aliyun.com/document_detail/2850189.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123456****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s DeleteImageTransformRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageTransformRequest) GoString() string {
	return s.String()
}

func (s *DeleteImageTransformRequest) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *DeleteImageTransformRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *DeleteImageTransformRequest) SetConfigId(v int64) *DeleteImageTransformRequest {
	s.ConfigId = &v
	return s
}

func (s *DeleteImageTransformRequest) SetSiteId(v int64) *DeleteImageTransformRequest {
	s.SiteId = &v
	return s
}

func (s *DeleteImageTransformRequest) Validate() error {
	return dara.Validate(s)
}
