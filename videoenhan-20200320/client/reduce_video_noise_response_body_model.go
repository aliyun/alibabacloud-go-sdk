// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReduceVideoNoiseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ReduceVideoNoiseResponseBodyData) *ReduceVideoNoiseResponseBody
	GetData() *ReduceVideoNoiseResponseBodyData
	SetMessage(v string) *ReduceVideoNoiseResponseBody
	GetMessage() *string
	SetRequestId(v string) *ReduceVideoNoiseResponseBody
	GetRequestId() *string
}

type ReduceVideoNoiseResponseBody struct {
	Data    *ReduceVideoNoiseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// d21a2afa-4d52-4bca-803b-e65028146603
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReduceVideoNoiseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReduceVideoNoiseResponseBody) GoString() string {
	return s.String()
}

func (s *ReduceVideoNoiseResponseBody) GetData() *ReduceVideoNoiseResponseBodyData {
	return s.Data
}

func (s *ReduceVideoNoiseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ReduceVideoNoiseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReduceVideoNoiseResponseBody) SetData(v *ReduceVideoNoiseResponseBodyData) *ReduceVideoNoiseResponseBody {
	s.Data = v
	return s
}

func (s *ReduceVideoNoiseResponseBody) SetMessage(v string) *ReduceVideoNoiseResponseBody {
	s.Message = &v
	return s
}

func (s *ReduceVideoNoiseResponseBody) SetRequestId(v string) *ReduceVideoNoiseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReduceVideoNoiseResponseBody) Validate() error {
	return dara.Validate(s)
}

type ReduceVideoNoiseResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-xstream-cn-shanghai.oss-cn-shanghai.aliyuncs.com/xstream-framework/upload_result_video_2023-02-10_09.45.55.mp4?Expires=1675995564&amp;OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&amp;Signature=aIXTeM4IU4nARjy3SNA3YGhhqj****
	VideoUrl *string `json:"VideoUrl,omitempty" xml:"VideoUrl,omitempty"`
}

func (s ReduceVideoNoiseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ReduceVideoNoiseResponseBodyData) GoString() string {
	return s.String()
}

func (s *ReduceVideoNoiseResponseBodyData) GetVideoUrl() *string {
	return s.VideoUrl
}

func (s *ReduceVideoNoiseResponseBodyData) SetVideoUrl(v string) *ReduceVideoNoiseResponseBodyData {
	s.VideoUrl = &v
	return s
}

func (s *ReduceVideoNoiseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
