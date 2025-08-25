// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iColorizeImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ColorizeImageResponseBodyData) *ColorizeImageResponseBody
	GetData() *ColorizeImageResponseBodyData
	SetRequestId(v string) *ColorizeImageResponseBody
	GetRequestId() *string
}

type ColorizeImageResponseBody struct {
	Data *ColorizeImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 124A4B09-68EF-4178-B98D-319089D4268B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ColorizeImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ColorizeImageResponseBody) GoString() string {
	return s.String()
}

func (s *ColorizeImageResponseBody) GetData() *ColorizeImageResponseBodyData {
	return s.Data
}

func (s *ColorizeImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ColorizeImageResponseBody) SetData(v *ColorizeImageResponseBodyData) *ColorizeImageResponseBody {
	s.Data = v
	return s
}

func (s *ColorizeImageResponseBody) SetRequestId(v string) *ColorizeImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ColorizeImageResponseBody) Validate() error {
	return dara.Validate(s)
}

type ColorizeImageResponseBodyData struct {
	// example:
	//
	// http://algo-app-aic-vc-cn-shanghai-prod.oss-cn-shanghai.aliyuncs.com/face-enhancement/2020_11_26/20201126_182812286268_079260.jpg?Expires=1606388292&OSSAccessKeyId=LTAI4FoLmvQ9urWXgSRp****&Signature=f71Bx37g%2BGhM%2B6FOXM0EbNL8W4****
	ImageURL *string `json:"ImageURL,omitempty" xml:"ImageURL,omitempty"`
}

func (s ColorizeImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ColorizeImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *ColorizeImageResponseBodyData) GetImageURL() *string {
	return s.ImageURL
}

func (s *ColorizeImageResponseBodyData) SetImageURL(v string) *ColorizeImageResponseBodyData {
	s.ImageURL = &v
	return s
}

func (s *ColorizeImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
