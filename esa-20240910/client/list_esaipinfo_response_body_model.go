// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListESAIPInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v []*ListESAIPInfoResponseBodyContent) *ListESAIPInfoResponseBody
	GetContent() []*ListESAIPInfoResponseBodyContent
	SetRequestId(v string) *ListESAIPInfoResponseBody
	GetRequestId() *string
}

type ListESAIPInfoResponseBody struct {
	Content []*ListESAIPInfoResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// example:
	//
	// 36af3fcc-43d0-441c-86b1-428951dc8225
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListESAIPInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListESAIPInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListESAIPInfoResponseBody) GetContent() []*ListESAIPInfoResponseBodyContent {
	return s.Content
}

func (s *ListESAIPInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListESAIPInfoResponseBody) SetContent(v []*ListESAIPInfoResponseBodyContent) *ListESAIPInfoResponseBody {
	s.Content = v
	return s
}

func (s *ListESAIPInfoResponseBody) SetRequestId(v string) *ListESAIPInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListESAIPInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListESAIPInfoResponseBodyContent struct {
	// example:
	//
	// true
	CdnIp *string `json:"CdnIp,omitempty" xml:"CdnIp,omitempty"`
	// example:
	//
	// 27.129.167.239
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s ListESAIPInfoResponseBodyContent) String() string {
	return dara.Prettify(s)
}

func (s ListESAIPInfoResponseBodyContent) GoString() string {
	return s.String()
}

func (s *ListESAIPInfoResponseBodyContent) GetCdnIp() *string {
	return s.CdnIp
}

func (s *ListESAIPInfoResponseBodyContent) GetIp() *string {
	return s.Ip
}

func (s *ListESAIPInfoResponseBodyContent) SetCdnIp(v string) *ListESAIPInfoResponseBodyContent {
	s.CdnIp = &v
	return s
}

func (s *ListESAIPInfoResponseBodyContent) SetIp(v string) *ListESAIPInfoResponseBodyContent {
	s.Ip = &v
	return s
}

func (s *ListESAIPInfoResponseBodyContent) Validate() error {
	return dara.Validate(s)
}
