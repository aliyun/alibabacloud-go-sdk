// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSupportLinesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeSupportLinesRequest
	GetDomainName() *string
	SetLang(v string) *DescribeSupportLinesRequest
	GetLang() *string
	SetUserClientIp(v string) *DescribeSupportLinesRequest
	GetUserClientIp() *string
}

type DescribeSupportLinesRequest struct {
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// example:
	//
	// 1.1.*.*
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeSupportLinesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSupportLinesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSupportLinesRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeSupportLinesRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeSupportLinesRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DescribeSupportLinesRequest) SetDomainName(v string) *DescribeSupportLinesRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeSupportLinesRequest) SetLang(v string) *DescribeSupportLinesRequest {
	s.Lang = &v
	return s
}

func (s *DescribeSupportLinesRequest) SetUserClientIp(v string) *DescribeSupportLinesRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeSupportLinesRequest) Validate() error {
	return dara.Validate(s)
}
