// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamWatermarkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteLiveStreamWatermarkRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteLiveStreamWatermarkRequest
	GetRegionId() *string
	SetTemplateId(v string) *DeleteLiveStreamWatermarkRequest
	GetTemplateId() *string
}

type DeleteLiveStreamWatermarkRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the watermark template.
	//
	// >  You can obtain the template ID by checking the value of the TemplateId parameter that is returned by the [AddLiveStreamWatermark](https://help.aliyun.com/document_detail/410759.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 445409ec-7eaa-461d-8f29-4bec2eb9****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteLiveStreamWatermarkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamWatermarkRequest) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamWatermarkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteLiveStreamWatermarkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteLiveStreamWatermarkRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteLiveStreamWatermarkRequest) SetOwnerId(v int64) *DeleteLiveStreamWatermarkRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteLiveStreamWatermarkRequest) SetRegionId(v string) *DeleteLiveStreamWatermarkRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteLiveStreamWatermarkRequest) SetTemplateId(v string) *DeleteLiveStreamWatermarkRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteLiveStreamWatermarkRequest) Validate() error {
	return dara.Validate(s)
}
