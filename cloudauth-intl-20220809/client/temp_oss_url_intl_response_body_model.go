// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempOssUrlIntlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *TempOssUrlIntlResponseBody
	GetCode() *string
	SetData(v *TempOssUrlIntlResponseBodyData) *TempOssUrlIntlResponseBody
	GetData() *TempOssUrlIntlResponseBodyData
	SetMessage(v string) *TempOssUrlIntlResponseBody
	GetMessage() *string
	SetRequestId(v string) *TempOssUrlIntlResponseBody
	GetRequestId() *string
}

type TempOssUrlIntlResponseBody struct {
	// example:
	//
	// Success
	Code *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *TempOssUrlIntlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// EFA11401-C961-5E89-A2D3-BF9883E5CC3D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TempOssUrlIntlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TempOssUrlIntlResponseBody) GoString() string {
	return s.String()
}

func (s *TempOssUrlIntlResponseBody) GetCode() *string {
	return s.Code
}

func (s *TempOssUrlIntlResponseBody) GetData() *TempOssUrlIntlResponseBodyData {
	return s.Data
}

func (s *TempOssUrlIntlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *TempOssUrlIntlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TempOssUrlIntlResponseBody) SetCode(v string) *TempOssUrlIntlResponseBody {
	s.Code = &v
	return s
}

func (s *TempOssUrlIntlResponseBody) SetData(v *TempOssUrlIntlResponseBodyData) *TempOssUrlIntlResponseBody {
	s.Data = v
	return s
}

func (s *TempOssUrlIntlResponseBody) SetMessage(v string) *TempOssUrlIntlResponseBody {
	s.Message = &v
	return s
}

func (s *TempOssUrlIntlResponseBody) SetRequestId(v string) *TempOssUrlIntlResponseBody {
	s.RequestId = &v
	return s
}

func (s *TempOssUrlIntlResponseBody) Validate() error {
	return dara.Validate(s)
}

type TempOssUrlIntlResponseBodyData struct {
	// example:
	//
	// http://bzxh.cdn.weijin365.com/assets/index-55338127.png
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s TempOssUrlIntlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s TempOssUrlIntlResponseBodyData) GoString() string {
	return s.String()
}

func (s *TempOssUrlIntlResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *TempOssUrlIntlResponseBodyData) SetUrl(v string) *TempOssUrlIntlResponseBodyData {
	s.Url = &v
	return s
}

func (s *TempOssUrlIntlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
