// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSplitVideoPartsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxTime(v int32) *SplitVideoPartsRequest
	GetMaxTime() *int32
	SetMinTime(v int32) *SplitVideoPartsRequest
	GetMinTime() *int32
	SetTemplate(v string) *SplitVideoPartsRequest
	GetTemplate() *string
	SetVideoUrl(v string) *SplitVideoPartsRequest
	GetVideoUrl() *string
}

type SplitVideoPartsRequest struct {
	MaxTime  *int32  `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	MinTime  *int32  `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/ocr/xxxx.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SplitVideoPartsRequest) String() string {
	return dara.Prettify(s)
}

func (s SplitVideoPartsRequest) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsRequest) GetMaxTime() *int32 {
	return s.MaxTime
}

func (s *SplitVideoPartsRequest) GetMinTime() *int32 {
	return s.MinTime
}

func (s *SplitVideoPartsRequest) GetTemplate() *string {
	return s.Template
}

func (s *SplitVideoPartsRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *SplitVideoPartsRequest) SetMaxTime(v int32) *SplitVideoPartsRequest {
	s.MaxTime = &v
	return s
}

func (s *SplitVideoPartsRequest) SetMinTime(v int32) *SplitVideoPartsRequest {
	s.MinTime = &v
	return s
}

func (s *SplitVideoPartsRequest) SetTemplate(v string) *SplitVideoPartsRequest {
	s.Template = &v
	return s
}

func (s *SplitVideoPartsRequest) SetVideoUrl(v string) *SplitVideoPartsRequest {
	s.VideoUrl = &v
	return s
}

func (s *SplitVideoPartsRequest) Validate() error {
	return dara.Validate(s)
}
