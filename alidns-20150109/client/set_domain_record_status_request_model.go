// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainRecordStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *SetDomainRecordStatusRequest
	GetLang() *string
	SetRecordId(v string) *SetDomainRecordStatusRequest
	GetRecordId() *string
	SetStatus(v string) *SetDomainRecordStatusRequest
	GetStatus() *string
	SetUserClientIp(v string) *SetDomainRecordStatusRequest
	GetUserClientIp() *string
}

type SetDomainRecordStatusRequest struct {
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
	// The ID of the DNS record. You can call the [DescribeDomainRecords](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomainrecords?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9999985
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The state of the DNS record. Valid values:
	//
	// 	- **Enable**: enables the DNS record.
	//
	// 	- **Disable**: disables the DNS record.
	//
	// This parameter is required.
	//
	// example:
	//
	// Disable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 192.0.2.0
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SetDomainRecordStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s SetDomainRecordStatusRequest) GoString() string {
	return s.String()
}

func (s *SetDomainRecordStatusRequest) GetLang() *string {
	return s.Lang
}

func (s *SetDomainRecordStatusRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *SetDomainRecordStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *SetDomainRecordStatusRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SetDomainRecordStatusRequest) SetLang(v string) *SetDomainRecordStatusRequest {
	s.Lang = &v
	return s
}

func (s *SetDomainRecordStatusRequest) SetRecordId(v string) *SetDomainRecordStatusRequest {
	s.RecordId = &v
	return s
}

func (s *SetDomainRecordStatusRequest) SetStatus(v string) *SetDomainRecordStatusRequest {
	s.Status = &v
	return s
}

func (s *SetDomainRecordStatusRequest) SetUserClientIp(v string) *SetDomainRecordStatusRequest {
	s.UserClientIp = &v
	return s
}

func (s *SetDomainRecordStatusRequest) Validate() error {
	return dara.Validate(s)
}
