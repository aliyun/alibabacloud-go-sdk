// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDomainRecordInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeDomainRecordInfoRequest
	GetLang() *string
	SetRecordId(v string) *DescribeDomainRecordInfoRequest
	GetRecordId() *string
	SetUserClientIp(v string) *DescribeDomainRecordInfoRequest
	GetUserClientIp() *string
}

type DescribeDomainRecordInfoRequest struct {
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
	// The ID of the DNS record. You can call the [DescribeDomainRecords](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomainrecords?spm=a2c63.p38356.help-menu-search-29697.d_0) to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9999985
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 192.0.2.0
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DescribeDomainRecordInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDomainRecordInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeDomainRecordInfoRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeDomainRecordInfoRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *DescribeDomainRecordInfoRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DescribeDomainRecordInfoRequest) SetLang(v string) *DescribeDomainRecordInfoRequest {
	s.Lang = &v
	return s
}

func (s *DescribeDomainRecordInfoRequest) SetRecordId(v string) *DescribeDomainRecordInfoRequest {
	s.RecordId = &v
	return s
}

func (s *DescribeDomainRecordInfoRequest) SetUserClientIp(v string) *DescribeDomainRecordInfoRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeDomainRecordInfoRequest) Validate() error {
	return dara.Validate(s)
}
