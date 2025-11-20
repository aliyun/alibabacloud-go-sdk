// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuperResolveVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SuperResolveVideoResponseBodyData) *SuperResolveVideoResponseBody
	GetData() *SuperResolveVideoResponseBodyData
	SetMessage(v string) *SuperResolveVideoResponseBody
	GetMessage() *string
	SetRequestId(v string) *SuperResolveVideoResponseBody
	GetRequestId() *string
}

type SuperResolveVideoResponseBody struct {
	Data    *SuperResolveVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 186AC396-0EEC-46F1-AAA1-BF3585227427
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SuperResolveVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuperResolveVideoResponseBody) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoResponseBody) GetData() *SuperResolveVideoResponseBodyData {
	return s.Data
}

func (s *SuperResolveVideoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SuperResolveVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuperResolveVideoResponseBody) SetData(v *SuperResolveVideoResponseBodyData) *SuperResolveVideoResponseBody {
	s.Data = v
	return s
}

func (s *SuperResolveVideoResponseBody) SetMessage(v string) *SuperResolveVideoResponseBody {
	s.Message = &v
	return s
}

func (s *SuperResolveVideoResponseBody) SetRequestId(v string) *SuperResolveVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuperResolveVideoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SuperResolveVideoResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-super-resolution/2020-03-20-12/12%3A11-UlLVELFzIy5EAyEh.mp4?Expires=1584708132&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=0V8yKrCVybC4KIPtRuGKJDJaQT****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s SuperResolveVideoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SuperResolveVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *SuperResolveVideoResponseBodyData) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *SuperResolveVideoResponseBodyData) SetVideoUrl(v string) *SuperResolveVideoResponseBodyData {
	s.VideoUrl = &v
	return s
}

func (s *SuperResolveVideoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
