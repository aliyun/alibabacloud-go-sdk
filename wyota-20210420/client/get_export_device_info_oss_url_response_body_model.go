// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExportDeviceInfoOssUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetExportDeviceInfoOssUrlResponseBody
	GetCode() *string
	SetData(v *GetExportDeviceInfoOssUrlResponseBodyData) *GetExportDeviceInfoOssUrlResponseBody
	GetData() *GetExportDeviceInfoOssUrlResponseBodyData
	SetMessage(v string) *GetExportDeviceInfoOssUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetExportDeviceInfoOssUrlResponseBody
	GetRequestId() *string
}

type GetExportDeviceInfoOssUrlResponseBody struct {
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetExportDeviceInfoOssUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetExportDeviceInfoOssUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetExportDeviceInfoOssUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetExportDeviceInfoOssUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetExportDeviceInfoOssUrlResponseBody) GetData() *GetExportDeviceInfoOssUrlResponseBodyData {
	return s.Data
}

func (s *GetExportDeviceInfoOssUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetExportDeviceInfoOssUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetExportDeviceInfoOssUrlResponseBody) SetCode(v string) *GetExportDeviceInfoOssUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetExportDeviceInfoOssUrlResponseBody) SetData(v *GetExportDeviceInfoOssUrlResponseBodyData) *GetExportDeviceInfoOssUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetExportDeviceInfoOssUrlResponseBody) SetMessage(v string) *GetExportDeviceInfoOssUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetExportDeviceInfoOssUrlResponseBody) SetRequestId(v string) *GetExportDeviceInfoOssUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetExportDeviceInfoOssUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetExportDeviceInfoOssUrlResponseBodyData struct {
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s GetExportDeviceInfoOssUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetExportDeviceInfoOssUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetExportDeviceInfoOssUrlResponseBodyData) GetUrl() *string {
	return s.Url
}

func (s *GetExportDeviceInfoOssUrlResponseBodyData) SetUrl(v string) *GetExportDeviceInfoOssUrlResponseBodyData {
	s.Url = &v
	return s
}

func (s *GetExportDeviceInfoOssUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
