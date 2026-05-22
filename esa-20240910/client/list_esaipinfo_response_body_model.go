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
	// The objects that are returned.
	Content []*ListESAIPInfoResponseBodyContent `json:"Content,omitempty" xml:"Content,omitempty" type:"Repeated"`
	// The request ID.
	//
	// Example D03F9502-6653-127C-8A5F-0647197\\*\\*\\*\\*\\*
	//
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
	if s.Content != nil {
		for _, item := range s.Content {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListESAIPInfoResponseBodyContent struct {
	// Whether the IP address in the parameter belongs to ESA POPs.
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CdnIp *string `json:"CdnIp,omitempty" xml:"CdnIp,omitempty"`
	// The IP addresses.
	//
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
