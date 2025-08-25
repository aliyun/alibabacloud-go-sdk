// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iRemoveImageSubtitlesAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBH(v float32) *RemoveImageSubtitlesAdvanceRequest
	GetBH() *float32
	SetBW(v float32) *RemoveImageSubtitlesAdvanceRequest
	GetBW() *float32
	SetBX(v float32) *RemoveImageSubtitlesAdvanceRequest
	GetBX() *float32
	SetBY(v float32) *RemoveImageSubtitlesAdvanceRequest
	GetBY() *float32
	SetImageURLObject(v io.Reader) *RemoveImageSubtitlesAdvanceRequest
	GetImageURLObject() io.Reader
}

type RemoveImageSubtitlesAdvanceRequest struct {
	// example:
	//
	// 0.25
	BH *float32 `json:"BH,omitempty" xml:"BH,omitempty"`
	// example:
	//
	// 1
	BW *float32 `json:"BW,omitempty" xml:"BW,omitempty"`
	// example:
	//
	// 0
	BX *float32 `json:"BX,omitempty" xml:"BX,omitempty"`
	// example:
	//
	// 0.75
	BY *float32 `json:"BY,omitempty" xml:"BY,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/imageenhan/RemoveImageSubtitles/RemoveImageSubtitles1.jpg
	ImageURLObject io.Reader `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RemoveImageSubtitlesAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageSubtitlesAdvanceRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageSubtitlesAdvanceRequest) GetBH() *float32 {
	return s.BH
}

func (s *RemoveImageSubtitlesAdvanceRequest) GetBW() *float32 {
	return s.BW
}

func (s *RemoveImageSubtitlesAdvanceRequest) GetBX() *float32 {
	return s.BX
}

func (s *RemoveImageSubtitlesAdvanceRequest) GetBY() *float32 {
	return s.BY
}

func (s *RemoveImageSubtitlesAdvanceRequest) GetImageURLObject() io.Reader {
	return s.ImageURLObject
}

func (s *RemoveImageSubtitlesAdvanceRequest) SetBH(v float32) *RemoveImageSubtitlesAdvanceRequest {
	s.BH = &v
	return s
}

func (s *RemoveImageSubtitlesAdvanceRequest) SetBW(v float32) *RemoveImageSubtitlesAdvanceRequest {
	s.BW = &v
	return s
}

func (s *RemoveImageSubtitlesAdvanceRequest) SetBX(v float32) *RemoveImageSubtitlesAdvanceRequest {
	s.BX = &v
	return s
}

func (s *RemoveImageSubtitlesAdvanceRequest) SetBY(v float32) *RemoveImageSubtitlesAdvanceRequest {
	s.BY = &v
	return s
}

func (s *RemoveImageSubtitlesAdvanceRequest) SetImageURLObject(v io.Reader) *RemoveImageSubtitlesAdvanceRequest {
	s.ImageURLObject = v
	return s
}

func (s *RemoveImageSubtitlesAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
