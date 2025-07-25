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
	// The domain name. You can call the [DescribeDomains](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomains?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The language of the content within the request and response. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English Default: **zh**
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The resolution line. Default value: **default**. For more information, see
	//
	// [DNS resolution lines](https://www.alibabacloud.com/help/zh/doc-detail/29807.htm).
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The priority of the mail exchanger (MX) record. Valid values: `1 to 50`.
	//
	// This parameter is required if the type of the DNS record is MX. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The hostname.
	//
	// For example, to resolve @.example.com, you must set this parameter to an at sign (@). You cannot leave this parameter empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// www
	RR *string `json:"RR,omitempty" xml:"RR,omitempty"`
	// The time to live (TTL) period of the Alibaba Cloud DNS (DNS) record. Default value: 600. Unit: seconds. For more information, see
	//
	// [TTL definition](https://www.alibabacloud.com/help/zh/doc-detail/29806.htm).
	//
	// example:
	//
	// 600
	TTL *int64 `json:"TTL,omitempty" xml:"TTL,omitempty"`
	// The type of the DNS record. For more information, see
	//
	// [DNS record types](https://www.alibabacloud.com/help/zh/doc-detail/29805.htm).
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
	// 192.0.2.0
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The value of the DNS record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.0.2.254
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
