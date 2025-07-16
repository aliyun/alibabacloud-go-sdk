// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchDescribeCdnIpInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIpAddrList(v string) *BatchDescribeCdnIpInfoRequest
	GetIpAddrList() *string
	SetLanguage(v string) *BatchDescribeCdnIpInfoRequest
	GetLanguage() *string
}

type BatchDescribeCdnIpInfoRequest struct {
	// The list of IP addresses to query. Separate IP addresses with commas (,). You can specify up to 20 IP addresses at a time.
	//
	// > 	- Example of an IPv4 address: 192.0.2.1
	//
	// >	- Example of an IPv6 address: 2001:db8:ffff:ffff:ffff:\\*\\*\\*\\*:ffff.
	//
	// This parameter is required.
	//
	// example:
	//
	// 111.XXX.XXX.230,47.XXX.XXX.243
	IpAddrList *string `json:"IpAddrList,omitempty" xml:"IpAddrList,omitempty"`
	// The language of the query results. Valid values:
	//
	// 	- **zh*	- (default): Simplified Chinese.
	//
	// 	- **en**: English.
	//
	// example:
	//
	// en
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s BatchDescribeCdnIpInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchDescribeCdnIpInfoRequest) GoString() string {
	return s.String()
}

func (s *BatchDescribeCdnIpInfoRequest) GetIpAddrList() *string {
	return s.IpAddrList
}

func (s *BatchDescribeCdnIpInfoRequest) GetLanguage() *string {
	return s.Language
}

func (s *BatchDescribeCdnIpInfoRequest) SetIpAddrList(v string) *BatchDescribeCdnIpInfoRequest {
	s.IpAddrList = &v
	return s
}

func (s *BatchDescribeCdnIpInfoRequest) SetLanguage(v string) *BatchDescribeCdnIpInfoRequest {
	s.Language = &v
	return s
}

func (s *BatchDescribeCdnIpInfoRequest) Validate() error {
	return dara.Validate(s)
}
