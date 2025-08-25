// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveImageSubtitlesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBH(v float32) *RemoveImageSubtitlesRequest
	GetBH() *float32
	SetBW(v float32) *RemoveImageSubtitlesRequest
	GetBW() *float32
	SetBX(v float32) *RemoveImageSubtitlesRequest
	GetBX() *float32
	SetBY(v float32) *RemoveImageSubtitlesRequest
	GetBY() *float32
	SetImageURL(v string) *RemoveImageSubtitlesRequest
	GetImageURL() *string
}

type RemoveImageSubtitlesRequest struct {
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
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s RemoveImageSubtitlesRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveImageSubtitlesRequest) GoString() string {
	return s.String()
}

func (s *RemoveImageSubtitlesRequest) GetBH() *float32 {
	return s.BH
}

func (s *RemoveImageSubtitlesRequest) GetBW() *float32 {
	return s.BW
}

func (s *RemoveImageSubtitlesRequest) GetBX() *float32 {
	return s.BX
}

func (s *RemoveImageSubtitlesRequest) GetBY() *float32 {
	return s.BY
}

func (s *RemoveImageSubtitlesRequest) GetImageURL() *string {
	return s.ImageURL
}

func (s *RemoveImageSubtitlesRequest) SetBH(v float32) *RemoveImageSubtitlesRequest {
	s.BH = &v
	return s
}

func (s *RemoveImageSubtitlesRequest) SetBW(v float32) *RemoveImageSubtitlesRequest {
	s.BW = &v
	return s
}

func (s *RemoveImageSubtitlesRequest) SetBX(v float32) *RemoveImageSubtitlesRequest {
	s.BX = &v
	return s
}

func (s *RemoveImageSubtitlesRequest) SetBY(v float32) *RemoveImageSubtitlesRequest {
	s.BY = &v
	return s
}

func (s *RemoveImageSubtitlesRequest) SetImageURL(v string) *RemoveImageSubtitlesRequest {
	s.ImageURL = &v
	return s
}

func (s *RemoveImageSubtitlesRequest) Validate() error {
	return dara.Validate(s)
}
