// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAdjustVideoColorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *AdjustVideoColorResponseBodyData) *AdjustVideoColorResponseBody
	GetData() *AdjustVideoColorResponseBodyData
	SetMessage(v string) *AdjustVideoColorResponseBody
	GetMessage() *string
	SetRequestId(v string) *AdjustVideoColorResponseBody
	GetRequestId() *string
}

type AdjustVideoColorResponseBody struct {
	Data    *AdjustVideoColorResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C4EB5E0B-0718-42CC-9B2C-1FB149256874
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AdjustVideoColorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AdjustVideoColorResponseBody) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorResponseBody) GetData() *AdjustVideoColorResponseBodyData {
	return s.Data
}

func (s *AdjustVideoColorResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AdjustVideoColorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AdjustVideoColorResponseBody) SetData(v *AdjustVideoColorResponseBodyData) *AdjustVideoColorResponseBody {
	s.Data = v
	return s
}

func (s *AdjustVideoColorResponseBody) SetMessage(v string) *AdjustVideoColorResponseBody {
	s.Message = &v
	return s
}

func (s *AdjustVideoColorResponseBody) SetRequestId(v string) *AdjustVideoColorResponseBody {
	s.RequestId = &v
	return s
}

func (s *AdjustVideoColorResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AdjustVideoColorResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vd-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/video-recolor/2021-01-21-07/46%3A05-test.mov?Expires=1611216966&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=icKn5gEQ6rNlSHmCi2zAf2tC0L****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s AdjustVideoColorResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AdjustVideoColorResponseBodyData) GoString() string {
	return s.String()
}

func (s *AdjustVideoColorResponseBodyData) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *AdjustVideoColorResponseBodyData) SetVideoUrl(v string) *AdjustVideoColorResponseBodyData {
	s.VideoUrl = &v
	return s
}

func (s *AdjustVideoColorResponseBodyData) Validate() error {
	return dara.Validate(s)
}
