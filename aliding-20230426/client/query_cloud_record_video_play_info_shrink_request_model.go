// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCloudRecordVideoPlayInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConferenceId(v string) *QueryCloudRecordVideoPlayInfoShrinkRequest
	GetConferenceId() *string
	SetMediaId(v string) *QueryCloudRecordVideoPlayInfoShrinkRequest
	GetMediaId() *string
	SetRegionId(v string) *QueryCloudRecordVideoPlayInfoShrinkRequest
	GetRegionId() *string
	SetTenantContextShrink(v string) *QueryCloudRecordVideoPlayInfoShrinkRequest
	GetTenantContextShrink() *string
}

type QueryCloudRecordVideoPlayInfoShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 6139b4xxx
	ConferenceId *string `json:"ConferenceId,omitempty" xml:"ConferenceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 44444444
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId            *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
}

func (s QueryCloudRecordVideoPlayInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryCloudRecordVideoPlayInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *QueryCloudRecordVideoPlayInfoShrinkRequest) GetConferenceId() *string {
	return s.ConferenceId
}

func (s *QueryCloudRecordVideoPlayInfoShrinkRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *QueryCloudRecordVideoPlayInfoShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *QueryCloudRecordVideoPlayInfoShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *QueryCloudRecordVideoPlayInfoShrinkRequest) SetConferenceId(v string) *QueryCloudRecordVideoPlayInfoShrinkRequest {
	s.ConferenceId = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoShrinkRequest) SetMediaId(v string) *QueryCloudRecordVideoPlayInfoShrinkRequest {
	s.MediaId = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoShrinkRequest) SetRegionId(v string) *QueryCloudRecordVideoPlayInfoShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoShrinkRequest) SetTenantContextShrink(v string) *QueryCloudRecordVideoPlayInfoShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *QueryCloudRecordVideoPlayInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
