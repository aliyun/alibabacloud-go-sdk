// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertHdrVideoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ConvertHdrVideoResponseBodyData) *ConvertHdrVideoResponseBody
	GetData() *ConvertHdrVideoResponseBodyData
	SetMessage(v string) *ConvertHdrVideoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConvertHdrVideoResponseBody
	GetRequestId() *string
}

type ConvertHdrVideoResponseBody struct {
	Data    *ConvertHdrVideoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// E0CF495F-E806-4B9C-B204-E1230608239D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConvertHdrVideoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConvertHdrVideoResponseBody) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoResponseBody) GetData() *ConvertHdrVideoResponseBodyData {
	return s.Data
}

func (s *ConvertHdrVideoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConvertHdrVideoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConvertHdrVideoResponseBody) SetData(v *ConvertHdrVideoResponseBodyData) *ConvertHdrVideoResponseBody {
	s.Data = v
	return s
}

func (s *ConvertHdrVideoResponseBody) SetMessage(v string) *ConvertHdrVideoResponseBody {
	s.Message = &v
	return s
}

func (s *ConvertHdrVideoResponseBody) SetRequestId(v string) *ConvertHdrVideoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConvertHdrVideoResponseBody) Validate() error {
	return dara.Validate(s)
}

type ConvertHdrVideoResponseBodyData struct {
	// example:
	//
	// http://vibktprfx-prod-prod-aic-vd-cn-shanghai.oss-cn-shanghai.aliyuncs.com/hdr-enhance/20-12-22/HaKDdTI48i2GQGy7_20-12-22-06-42-45.mp4?Expires=1608621178&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=SWc90T0JHg5eWc64x8GmYHKsvX****
	VideoURL *string `json:"VideoURL,omitempty" xml:"VideoURL,omitempty"`
}

func (s ConvertHdrVideoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ConvertHdrVideoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConvertHdrVideoResponseBodyData) GetVideoURL() *string {
	return s.VideoURL
}

func (s *ConvertHdrVideoResponseBodyData) SetVideoURL(v string) *ConvertHdrVideoResponseBodyData {
	s.VideoURL = &v
	return s
}

func (s *ConvertHdrVideoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
