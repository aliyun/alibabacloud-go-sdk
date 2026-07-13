// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddDomainRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *AddDomainRecordRequest
	GetDomainName() *string
	SetLang(v string) *AddDomainRecordRequest
	GetLang() *string
	SetLine(v string) *AddDomainRecordRequest
	GetLine() *string
	SetPriority(v int64) *AddDomainRecordRequest
	GetPriority() *int64
	SetRR(v string) *AddDomainRecordRequest
	GetRR() *string
	SetTTL(v int64) *AddDomainRecordRequest
	GetTTL() *int64
	SetType(v string) *AddDomainRecordRequest
	GetType() *string
	SetUserClientIp(v string) *AddDomainRecordRequest
	GetUserClientIp() *string
	SetValue(v string) *AddDomainRecordRequest
	GetValue() *string
}

type AddDomainRecordRequest struct {
	// The domain name. Call the [DescribeDomains](https://www.alibabacloud.com/help/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to query the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the request and response. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English
	//
	//   The default value is **zh**.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The DNS resolution line. The default value is **default**. For more information, see [DNS resolution lines](https://www.alibabacloud.com/help/doc-detail/29807.htm).
	//
	// <props="china">
	//
	// [Resolution line enumeration](https://help.aliyun.com/document_detail/29807.html)
	//
	//
	//
	// <props="intl">
	//
	// [Resolution Line Enumeration](https://www.alibabacloud.com/help/zh/doc-detail/29807.htm)
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The priority of the MX record. Valid values: `[1,50]`.
	//
	// This parameter is required if the record type is MX. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The host record.
	//
	// To resolve example.com, set the host record to "@" instead of leaving it empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// www
	RR *string `json:"RR,omitempty" xml:"RR,omitempty"`
	// The time to live (TTL) value of the Domain Name System (DNS) record. Default value: 600. Unit: seconds. For more information, see the following topic:
	//
	// <props="china">
	//
	// [TTL overview](https://help.aliyun.com/document_detail/29806.html)
	//
	//
	//
	// <props="intl">
	//
	// The time to live (TTL) of the DNS record. The default value is 600 seconds. For more information, see [TTL](https://www.alibabacloud.com/help/doc-detail/29806.htm).
	//
	// example:
	//
	// 600
	TTL *int64 `json:"TTL,omitempty" xml:"TTL,omitempty"`
	// The type of the DNS record. For more information, see
	//
	// <props="china">
	//
	// [DNS record type format](https://help.aliyun.com/document_detail/29805.html)
	//
	//
	//
	// <props="intl">
	//
	// The type of the DNS record. For more information, see [DNS record types](https://www.alibabacloud.com/help/doc-detail/29805.htm).
	//
	// This parameter is required.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 192.0.2.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The record value.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.0.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddDomainRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s AddDomainRecordRequest) GoString() string {
	return s.String()
}

func (s *AddDomainRecordRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *AddDomainRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *AddDomainRecordRequest) GetLine() *string {
	return s.Line
}

func (s *AddDomainRecordRequest) GetPriority() *int64 {
	return s.Priority
}

func (s *AddDomainRecordRequest) GetRR() *string {
	return s.RR
}

func (s *AddDomainRecordRequest) GetTTL() *int64 {
	return s.TTL
}

func (s *AddDomainRecordRequest) GetType() *string {
	return s.Type
}

func (s *AddDomainRecordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *AddDomainRecordRequest) GetValue() *string {
	return s.Value
}

func (s *AddDomainRecordRequest) SetDomainName(v string) *AddDomainRecordRequest {
	s.DomainName = &v
	return s
}

func (s *AddDomainRecordRequest) SetLang(v string) *AddDomainRecordRequest {
	s.Lang = &v
	return s
}

func (s *AddDomainRecordRequest) SetLine(v string) *AddDomainRecordRequest {
	s.Line = &v
	return s
}

func (s *AddDomainRecordRequest) SetPriority(v int64) *AddDomainRecordRequest {
	s.Priority = &v
	return s
}

func (s *AddDomainRecordRequest) SetRR(v string) *AddDomainRecordRequest {
	s.RR = &v
	return s
}

func (s *AddDomainRecordRequest) SetTTL(v int64) *AddDomainRecordRequest {
	s.TTL = &v
	return s
}

func (s *AddDomainRecordRequest) SetType(v string) *AddDomainRecordRequest {
	s.Type = &v
	return s
}

func (s *AddDomainRecordRequest) SetUserClientIp(v string) *AddDomainRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddDomainRecordRequest) SetValue(v string) *AddDomainRecordRequest {
	s.Value = &v
	return s
}

func (s *AddDomainRecordRequest) Validate() error {
	return dara.Validate(s)
}
