// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
	"io"
)

type iSuperResolveVideoAdvanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBitRate(v int32) *SuperResolveVideoAdvanceRequest
	GetBitRate() *int32
	SetVideoUrlObject(v io.Reader) *SuperResolveVideoAdvanceRequest
	GetVideoUrlObject() io.Reader
}

type SuperResolveVideoAdvanceRequest struct {
	// example:
	//
	// 5
	BitRate *int32 `json:"BitRate,omitempty" xml:"BitRate,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// http://viapi-test.oss-cn-shanghai.aliyuncs.com/viapi-3.0domepic/videoenhan/SuperResolveVideo/SuperResolveVideo2.mp4
	VideoUrlObject io.Reader `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SuperResolveVideoAdvanceRequest) String() string {
	return dara.Prettify(s)
}

func (s SuperResolveVideoAdvanceRequest) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoAdvanceRequest) GetBitRate() *int32 {
	return s.BitRate
}

func (s *SuperResolveVideoAdvanceRequest) GetVideoUrlObject() io.Reader {
	return s.VideoUrlObject
}

func (s *SuperResolveVideoAdvanceRequest) SetBitRate(v int32) *SuperResolveVideoAdvanceRequest {
	s.BitRate = &v
	return s
}

func (s *SuperResolveVideoAdvanceRequest) SetVideoUrlObject(v io.Reader) *SuperResolveVideoAdvanceRequest {
	s.VideoUrlObject = v
	return s
}

func (s *SuperResolveVideoAdvanceRequest) Validate() error {
	return dara.Validate(s)
}
