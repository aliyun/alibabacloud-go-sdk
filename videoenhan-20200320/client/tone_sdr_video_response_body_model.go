// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iToneSdrVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ToneSdrVideoResponseBodyData) *ToneSdrVideoResponseBody
	GetData() *ToneSdrVideoResponseBodyData
	SetMessage(v string) *ToneSdrVideoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ToneSdrVideoResponseBody
	GetRequestId() *string
}

type ToneSdrVideoResponseBody struct {
	Data    *ToneSdrVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 89B5AFF1-8A64-4F76-B391-56AD7D22DE35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ToneSdrVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ToneSdrVideoResponseBody) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoResponseBody) GetData() *ToneSdrVideoResponseBodyData {
	return s.Data
}

func (s *ToneSdrVideoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ToneSdrVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ToneSdrVideoResponseBody) SetData(v *ToneSdrVideoResponseBodyData) *ToneSdrVideoResponseBody {
	s.Data = v
	return s
}

func (s *ToneSdrVideoResponseBody) SetMessage(v string) *ToneSdrVideoResponseBody {
	s.Message = &v
	return s
}

func (s *ToneSdrVideoResponseBody) SetRequestId(v string) *ToneSdrVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ToneSdrVideoResponseBody) Validate() error {
	return dara.Validate(s)
}

type ToneSdrVideoResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/sdr-color-enhance/20-12-22/SxBKgwBhlObusG20_20-12-22-07-59-45.mp4?Expires=1608625795&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=maoOZ52y7U9ZuL2KqI0IfGq8%2FR****
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s ToneSdrVideoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ToneSdrVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ToneSdrVideoResponseBodyData) GetVideoURL() *string {
	return s.VideoURL
}

func (s *ToneSdrVideoResponseBodyData) SetVideoURL(v string) *ToneSdrVideoResponseBodyData {
	s.VideoURL = &v
	return s
}

func (s *ToneSdrVideoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
