// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSplitVideoPartsAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxTime(v int32) *SplitVideoPartsAdvanceRequest
	GetMaxTime() *int32
	SetMinTime(v int32) *SplitVideoPartsAdvanceRequest
	GetMinTime() *int32
	SetTemplate(v string) *SplitVideoPartsAdvanceRequest
	GetTemplate() *string
	SetVideoUrlObject(v io.Reader) *SplitVideoPartsAdvanceRequest
	GetVideoUrlObject() io.Reader
}

type SplitVideoPartsAdvanceRequest struct {
	MaxTime  *int32  `json:"MaxTime,omitempty" xml:"MaxTime,omitempty"`
	MinTime  *int32  `json:"MinTime,omitempty" xml:"MinTime,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// https://viapi-test.oss-cn-shanghai.aliyuncs.com/test-team/ocr/xxxx.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SplitVideoPartsAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SplitVideoPartsAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SplitVideoPartsAdvanceRequest) GetMaxTime() *int32 {
	return s.MaxTime
}

func (s *SplitVideoPartsAdvanceRequest) GetMinTime() *int32 {
	return s.MinTime
}

func (s *SplitVideoPartsAdvanceRequest) GetTemplate() *string {
	return s.Template
}

func (s *SplitVideoPartsAdvanceRequest) GetVideoUrlObject() io.Reader {
	return s.VideoUrlObject
}

func (s *SplitVideoPartsAdvanceRequest) SetMaxTime(v int32) *SplitVideoPartsAdvanceRequest {
	s.MaxTime = &v
	return s
}

func (s *SplitVideoPartsAdvanceRequest) SetMinTime(v int32) *SplitVideoPartsAdvanceRequest {
	s.MinTime = &v
	return s
}

func (s *SplitVideoPartsAdvanceRequest) SetTemplate(v string) *SplitVideoPartsAdvanceRequest {
	s.Template = &v
	return s
}

func (s *SplitVideoPartsAdvanceRequest) SetVideoUrlObject(v io.Reader) *SplitVideoPartsAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *SplitVideoPartsAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
