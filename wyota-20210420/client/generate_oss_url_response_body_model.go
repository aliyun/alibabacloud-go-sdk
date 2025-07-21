// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateOssUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GenerateOssUrlResponseBodyData) *GenerateOssUrlResponseBody
	GetData() []*GenerateOssUrlResponseBodyData
	SetRequestId(v string) *GenerateOssUrlResponseBody
	GetRequestId() *string
}

type GenerateOssUrlResponseBody struct {
	Data      []*GenerateOssUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateOssUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateOssUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateOssUrlResponseBody) GetData() []*GenerateOssUrlResponseBodyData {
	return s.Data
}

func (s *GenerateOssUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateOssUrlResponseBody) SetData(v []*GenerateOssUrlResponseBodyData) *GenerateOssUrlResponseBody {
	s.Data = v
	return s
}

func (s *GenerateOssUrlResponseBody) SetRequestId(v string) *GenerateOssUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateOssUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type GenerateOssUrlResponseBodyData struct {
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	ObjectName  *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
}

func (s GenerateOssUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateOssUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateOssUrlResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GenerateOssUrlResponseBodyData) GetObjectName() *string {
	return s.ObjectName
}

func (s *GenerateOssUrlResponseBodyData) SetDownloadUrl(v string) *GenerateOssUrlResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GenerateOssUrlResponseBodyData) SetObjectName(v string) *GenerateOssUrlResponseBodyData {
	s.ObjectName = &v
	return s
}

func (s *GenerateOssUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
