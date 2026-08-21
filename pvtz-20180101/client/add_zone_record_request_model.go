// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddZoneRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddZoneRecordRequest
	GetClientToken() *string
	SetLang(v string) *AddZoneRecordRequest
	GetLang() *string
	SetLine(v string) *AddZoneRecordRequest
	GetLine() *string
	SetPriority(v int32) *AddZoneRecordRequest
	GetPriority() *int32
	SetRemark(v string) *AddZoneRecordRequest
	GetRemark() *string
	SetRr(v string) *AddZoneRecordRequest
	GetRr() *string
	SetTtl(v int32) *AddZoneRecordRequest
	GetTtl() *int32
	SetType(v string) *AddZoneRecordRequest
	GetType() *string
	SetUserClientIp(v string) *AddZoneRecordRequest
	GetUserClientIp() *string
	SetValue(v string) *AddZoneRecordRequest
	GetValue() *string
	SetWeight(v int32) *AddZoneRecordRequest
	GetWeight() *int32
	SetZoneId(v string) *AddZoneRecordRequest
	GetZoneId() *string
}

type AddZoneRecordRequest struct {
	// The client token that is used to ensure the idempotence of the request. A client generates this value to ensure that it is unique among different requests. The value can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// Default value: en
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The source of the DNS resolution request. Valid values:
	//
	// - default: The default line. This is equivalent to a global line. Configure a default line to ensure that a DNS record is returned even if no smart line is hit.
	//
	// - Alibaba Cloud line: The DNS resolution request comes from Alibaba Cloud, including Public Cloud, Alibaba Finance Cloud, and Alibaba Gov Cloud.
	//
	// - Custom line: Customize internal domain name resolution to return a specific IP address for DNS query requests from a specific IP address segment.
	//
	// > 	- Only zones in built-in authoritative acceleration regions support adding DNS resolution request source lines.
	//
	// >
	//
	// > 	- To use the default line, enter "default". For Alibaba Cloud lines and custom lines, enter the specified line code. Example: aliyun_r_cn-beijing-a
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The priority of the MX record. A smaller value indicates a higher priority. Valid values: **[1, 99]**.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The remarks.
	//
	// example:
	//
	// en
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The host record. A host record is the prefix of a domain name. Common host records include www, @, \\	- (for wildcard DNS), and mail (for mailboxes).
	//
	// For example, to resolve @.example.com, set the host record to "@", not an empty string.
	//
	// This parameter is required.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The time to live (TTL). The unit is seconds (s). Valid values are 5, 30, 60, 3600 (1 hour), 43200 (12 hours), and 86400 (1 day). The default value is 60.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record. The following types are supported:
	//
	// - **A**: Maps a domain name to an IPv4 address in dotted decimal notation.
	//
	// - **AAAA**: Maps a domain name to an IPv6 address.
	//
	// - **CNAME**: Maps a domain name to another domain name.
	//
	// - **TXT**: A text record. The text can be up to 255 characters in length. TXT records are often used for Sender Policy Framework (SPF) records to prevent spam.
	//
	// - **MX**: Maps a domain name to the domain name of a mail server.
	//
	// - **PTR**: Maps an IP address to a domain name.
	//
	// - **SRV**: Specifies the server for a specific service. The format is: Priority Weight Port Target. Separate each value with a space.
	//
	// > Before adding a PTR record, configure a reverse lookup zone. For more information, see [Reverse DNS lookups and PTR records](https://help.aliyun.com/document_detail/2592976.html).
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
	// 127.0.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The record value. Enter a value based on the DNS record type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 114.55.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight. Valid values are integers from 1 to 100. The default value is 1. Set different weights for each address to return addresses based on the weight ratio for DNS queries.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The ID of the zone. This is the unique identifier of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// df2d03865266bd9842306db586d3****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddZoneRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s AddZoneRecordRequest) GoString() string {
	return s.String()
}

func (s *AddZoneRecordRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddZoneRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *AddZoneRecordRequest) GetLine() *string {
	return s.Line
}

func (s *AddZoneRecordRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *AddZoneRecordRequest) GetRemark() *string {
	return s.Remark
}

func (s *AddZoneRecordRequest) GetRr() *string {
	return s.Rr
}

func (s *AddZoneRecordRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *AddZoneRecordRequest) GetType() *string {
	return s.Type
}

func (s *AddZoneRecordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *AddZoneRecordRequest) GetValue() *string {
	return s.Value
}

func (s *AddZoneRecordRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *AddZoneRecordRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *AddZoneRecordRequest) SetClientToken(v string) *AddZoneRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *AddZoneRecordRequest) SetLang(v string) *AddZoneRecordRequest {
	s.Lang = &v
	return s
}

func (s *AddZoneRecordRequest) SetLine(v string) *AddZoneRecordRequest {
	s.Line = &v
	return s
}

func (s *AddZoneRecordRequest) SetPriority(v int32) *AddZoneRecordRequest {
	s.Priority = &v
	return s
}

func (s *AddZoneRecordRequest) SetRemark(v string) *AddZoneRecordRequest {
	s.Remark = &v
	return s
}

func (s *AddZoneRecordRequest) SetRr(v string) *AddZoneRecordRequest {
	s.Rr = &v
	return s
}

func (s *AddZoneRecordRequest) SetTtl(v int32) *AddZoneRecordRequest {
	s.Ttl = &v
	return s
}

func (s *AddZoneRecordRequest) SetType(v string) *AddZoneRecordRequest {
	s.Type = &v
	return s
}

func (s *AddZoneRecordRequest) SetUserClientIp(v string) *AddZoneRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddZoneRecordRequest) SetValue(v string) *AddZoneRecordRequest {
	s.Value = &v
	return s
}

func (s *AddZoneRecordRequest) SetWeight(v int32) *AddZoneRecordRequest {
	s.Weight = &v
	return s
}

func (s *AddZoneRecordRequest) SetZoneId(v string) *AddZoneRecordRequest {
	s.ZoneId = &v
	return s
}

func (s *AddZoneRecordRequest) Validate() error {
	return dara.Validate(s)
}
