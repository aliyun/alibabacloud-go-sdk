// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iAbstractFilmVideoAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLength(v int32) *AbstractFilmVideoAdvanceRequest
	GetLength() *int32
	SetVideoUrlObject(v io.Reader) *AbstractFilmVideoAdvanceRequest
	GetVideoUrlObject() io.Reader
}

type AbstractFilmVideoAdvanceRequest struct {
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
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AbstractFilmVideoAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s AbstractFilmVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *AbstractFilmVideoAdvanceRequest) GetLength() *int32 {
	return s.Length
}

func (s *AbstractFilmVideoAdvanceRequest) GetVideoUrlObject() io.Reader {
	return s.VideoUrlObject
}

func (s *AbstractFilmVideoAdvanceRequest) SetLength(v int32) *AbstractFilmVideoAdvanceRequest {
	s.Length = &v
	return s
}

func (s *AbstractFilmVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *AbstractFilmVideoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *AbstractFilmVideoAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
