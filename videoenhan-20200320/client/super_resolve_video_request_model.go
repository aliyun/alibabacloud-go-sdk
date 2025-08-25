// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuperResolveVideoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBitRate(v int32) *SuperResolveVideoRequest
	GetBitRate() *int32
	SetVideoUrl(v string) *SuperResolveVideoRequest
	GetVideoUrl() *string
}

type SuperResolveVideoRequest struct {
	// example:
	//
	// 5
	BitRate *int32 `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/SuperResolveVideo/SuperResolveVideo2.mp4
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SuperResolveVideoRequest) String() string {
	return dara.Prettify(s)
}

func (s SuperResolveVideoRequest) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoRequest) GetBitRate() *int32 {
	return s.BitRate
}

func (s *SuperResolveVideoRequest) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *SuperResolveVideoRequest) SetBitRate(v int32) *SuperResolveVideoRequest {
	s.BitRate = &v
	return s
}

func (s *SuperResolveVideoRequest) SetVideoUrl(v string) *SuperResolveVideoRequest {
	s.VideoUrl = &v
	return s
}

func (s *SuperResolveVideoRequest) Validate() error {
	return dara.Validate(s)
}
