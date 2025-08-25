// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbstractFilmVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLength(v int32) *AbstractFilmVideoRequest
	GetLength() *int32
	SetVideoUrl(v string) *AbstractFilmVideoRequest
	GetVideoUrl() *string
}

type AbstractFilmVideoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 5
	Length *int32 `json:"Length,omitempty" xml:"Length,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/AbstractFilmVideo/AbstractFilmVideo1.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AbstractFilmVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s AbstractFilmVideoRequest) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoRequest) GetLength() *int32 {
	return s.Length
}

func (s *AbstractFilmVideoRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *AbstractFilmVideoRequest) SetLength(v int32) *AbstractFilmVideoRequest {
	s.Length = &v
	return s
}

func (s *AbstractFilmVideoRequest) SetVideoUrl(v string) *AbstractFilmVideoRequest {
	s.VideoUrl = &v
	return s
}

func (s *AbstractFilmVideoRequest) Validate() error {
	return dara.Validate(s)
}
