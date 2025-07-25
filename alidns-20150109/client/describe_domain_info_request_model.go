// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeDomainInfoRequest
	GetDomainName() *string
	SetLang(v string) *DescribeDomainInfoRequest
	GetLang() *string
	SetNeedDetailAttributes(v bool) *DescribeDomainInfoRequest
	GetNeedDetailAttributes() *bool
}

type DescribeDomainInfoRequest struct {
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// dns-example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language type.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Specifies whether detailed attributes are required. Default value: **false**, which indicates that detailed attributes are not returned.
	//
	// If you set this parameter to **true**, the values of the following parameters are returned: LineType, MinTtl, RecordLineTreeJson, RecordLines, LineCode, LineDisplayName, LineName, RegionLines, and SlaveDns.
	//
	// example:
	//
	// true
	NeedDetailAttributes *bool `json:"NeedDetailAttributes,omitempty" xml:"NeedDetailAttributes,omitempty"`
}

func (s DescribeDomainInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainInfoRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeDomainInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainInfoRequest) GetNeedDetailAttributes() *bool {
	return s.NeedDetailAttributes
}

func (s *DescribeDomainInfoRequest) SetDomainName(v string) *DescribeDomainInfoRequest {
	s.DomainName = &v
	return s
}

func (s *DescribeDomainInfoRequest) SetLang(v string) *DescribeDomainInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainInfoRequest) SetNeedDetailAttributes(v bool) *DescribeDomainInfoRequest {
	s.NeedDetailAttributes = &v
	return s
}

func (s *DescribeDomainInfoRequest) Validate() error {
	return dara.Validate(s)
}
