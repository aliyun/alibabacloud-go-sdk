// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveContactChannelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ActiveContactChannelRequest
	GetCode() *string
	SetSite(v string) *ActiveContactChannelRequest
	GetSite() *string
}

type ActiveContactChannelRequest struct {
	// The activation verification code. This 32-character string is sent by the system to the mobile phone number or email address of the contact.
	//
	// example:
	//
	// 0b9377d9c56b5c2ea8a41a2b3c4d5e6f
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The site. Valid values: cn_hangzhou (China site) and ap_southeast_1 (International site). Default value: cn_hangzhou.
	//
	// example:
	//
	// cn_hangzhou
	Site *string `json:"site,omitempty" xml:"site,omitempty"`
}

func (s ActiveContactChannelRequest) String() string {
	return dara.Prettify(s)
}

func (s ActiveContactChannelRequest) GoString() string {
	return s.String()
}

func (s *ActiveContactChannelRequest) GetCode() *string {
	return s.Code
}

func (s *ActiveContactChannelRequest) GetSite() *string {
	return s.Site
}

func (s *ActiveContactChannelRequest) SetCode(v string) *ActiveContactChannelRequest {
	s.Code = &v
	return s
}

func (s *ActiveContactChannelRequest) SetSite(v string) *ActiveContactChannelRequest {
	s.Site = &v
	return s
}

func (s *ActiveContactChannelRequest) Validate() error {
	return dara.Validate(s)
}
