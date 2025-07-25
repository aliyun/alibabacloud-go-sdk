// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDomainRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *UpdateDomainRecordRequest
	GetLang() *string
	SetLine(v string) *UpdateDomainRecordRequest
	GetLine() *string
	SetPriority(v int64) *UpdateDomainRecordRequest
	GetPriority() *int64
	SetRR(v string) *UpdateDomainRecordRequest
	GetRR() *string
	SetRecordId(v string) *UpdateDomainRecordRequest
	GetRecordId() *string
	SetTTL(v int64) *UpdateDomainRecordRequest
	GetTTL() *int64
	SetType(v string) *UpdateDomainRecordRequest
	GetType() *string
	SetUserClientIp(v string) *UpdateDomainRecordRequest
	GetUserClientIp() *string
	SetValue(v string) *UpdateDomainRecordRequest
	GetValue() *string
}

type UpdateDomainRecordRequest struct {
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
	// The resolution line. Default value: **default**.
	//
	// For more information, see
	//
	// [DNS resolution lines](https://www.alibabacloud.com/help/zh/doc-detail/29807.htm).
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The priority of the mail exchanger (MX) record. Valid values: `1 to 50`.
	//
	// This parameter is required if the type of the DNS record is MX.
	//
	// example:
	//
	// 1
	Priority *int64 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The hostname.
	//
	// For example, if you want to resolve @.example.com, you must set RR to an at sign (@) instead of leaving it empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// www
	RR *string `json:"RR,omitempty" xml:"RR,omitempty"`
	// The ID of the Domain Name System (DNS) record.
	//
	// You can call the [DescribeDomainRecords](https://www.alibabacloud.com/help/zh/dns/api-alidns-2015-01-09-describedomainrecords?spm=a2c63.p38356.help-menu-search-29697.d_0) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9999985
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The time to live (TTL) period of the Alibaba Cloud DNS (DNS) record. Default value: 600. Unit: seconds.
	//
	// For more information, see
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

func (s UpdateDomainRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDomainRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateDomainRecordRequest) GetLine() *string {
	return s.Line
}

func (s *UpdateDomainRecordRequest) GetPriority() *int64 {
	return s.Priority
}

func (s *UpdateDomainRecordRequest) GetRR() *string {
	return s.RR
}

func (s *UpdateDomainRecordRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *UpdateDomainRecordRequest) GetTTL() *int64 {
	return s.TTL
}

func (s *UpdateDomainRecordRequest) GetType() *string {
	return s.Type
}

func (s *UpdateDomainRecordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *UpdateDomainRecordRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateDomainRecordRequest) SetLang(v string) *UpdateDomainRecordRequest {
	s.Lang = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetLine(v string) *UpdateDomainRecordRequest {
	s.Line = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetPriority(v int64) *UpdateDomainRecordRequest {
	s.Priority = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetRR(v string) *UpdateDomainRecordRequest {
	s.RR = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetRecordId(v string) *UpdateDomainRecordRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetTTL(v int64) *UpdateDomainRecordRequest {
	s.TTL = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetType(v string) *UpdateDomainRecordRequest {
	s.Type = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetUserClientIp(v string) *UpdateDomainRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateDomainRecordRequest) SetValue(v string) *UpdateDomainRecordRequest {
	s.Value = &v
	return s
}

func (s *UpdateDomainRecordRequest) Validate() error {
	return dara.Validate(s)
}
