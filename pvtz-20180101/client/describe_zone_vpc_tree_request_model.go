// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeZoneVpcTreeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeZoneVpcTreeRequest
	GetLang() *string
	SetUserClientIp(v string) *DescribeZoneVpcTreeRequest
	GetUserClientIp() *string
}

type DescribeZoneVpcTreeRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeZoneVpcTreeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeZoneVpcTreeRequest) GoString() string {
	return s.String()
}

func (s *DescribeZoneVpcTreeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeZoneVpcTreeRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DescribeZoneVpcTreeRequest) SetLang(v string) *DescribeZoneVpcTreeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeZoneVpcTreeRequest) SetUserClientIp(v string) *DescribeZoneVpcTreeRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeZoneVpcTreeRequest) Validate() error {
	return dara.Validate(s)
}
