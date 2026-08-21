// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZoneRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateZoneRecordRequest
	GetClientToken() *string
	SetLang(v string) *UpdateZoneRecordRequest
	GetLang() *string
	SetLine(v string) *UpdateZoneRecordRequest
	GetLine() *string
	SetPriority(v int32) *UpdateZoneRecordRequest
	GetPriority() *int32
	SetRecordId(v int64) *UpdateZoneRecordRequest
	GetRecordId() *int64
	SetRr(v string) *UpdateZoneRecordRequest
	GetRr() *string
	SetTtl(v int32) *UpdateZoneRecordRequest
	GetTtl() *int32
	SetType(v string) *UpdateZoneRecordRequest
	GetType() *string
	SetUserClientIp(v string) *UpdateZoneRecordRequest
	GetUserClientIp() *string
	SetValue(v string) *UpdateZoneRecordRequest
	GetValue() *string
	SetWeight(v int32) *UpdateZoneRecordRequest
	GetWeight() *int32
}

type UpdateZoneRecordRequest struct {
	// A client token that is used to ensure the idempotence of the request. The client generates the value, which must be unique among different requests. The token can contain a maximum of 64 ASCII characters.
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
	// The DNS resolution line. The default value is default.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The priority of the MX record. You can set different priorities for mail servers. Valid values: 1 to 99. A smaller value indicates a higher priority.
	//
	// > This parameter is required if the record type is MX.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the DNS record. To obtain the ID, call the DescribeZoneRecords operation to query a list of DNS records.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172223****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The host record. This is the prefix of a domain name. Common examples include www, @, \\	- (for wildcard DNS), and mail (for mailboxes).
	//
	// For example, to resolve @.example.com, set the host record to "@", not an empty string.
	//
	// This parameter is required.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The time to live (TTL) in seconds (s). Valid values: 5, 30, 60, 3600 (1 hour), 43200 (12 hours), and 86400 (1 day).
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
	// - **SRV**: A service record that specifies the server for a specific service. The format is: Priority Weight Port Target. Each part must be separated by a space.
	//
	// > Before adding a PTR record, configure a reverse lookup zone. For more information, see [Reverse DNS lookup and PTR records](https://help.aliyun.com/document_detail/2592976.html)
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
	// The record value. Enter a value that corresponds to the record type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.16.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record. Valid values are integers from 1 to 100. The default value is 1. You can set different weights for records to return IP addresses in proportion to their weights.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateZoneRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateZoneRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateZoneRecordRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateZoneRecordRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateZoneRecordRequest) GetLine() *string {
	return s.Line
}

func (s *UpdateZoneRecordRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateZoneRecordRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateZoneRecordRequest) GetRr() *string {
	return s.Rr
}

func (s *UpdateZoneRecordRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateZoneRecordRequest) GetType() *string {
	return s.Type
}

func (s *UpdateZoneRecordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *UpdateZoneRecordRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateZoneRecordRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateZoneRecordRequest) SetClientToken(v string) *UpdateZoneRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetLang(v string) *UpdateZoneRecordRequest {
	s.Lang = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetLine(v string) *UpdateZoneRecordRequest {
	s.Line = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetPriority(v int32) *UpdateZoneRecordRequest {
	s.Priority = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetRecordId(v int64) *UpdateZoneRecordRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetRr(v string) *UpdateZoneRecordRequest {
	s.Rr = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetTtl(v int32) *UpdateZoneRecordRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetType(v string) *UpdateZoneRecordRequest {
	s.Type = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetUserClientIp(v string) *UpdateZoneRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetValue(v string) *UpdateZoneRecordRequest {
	s.Value = &v
	return s
}

func (s *UpdateZoneRecordRequest) SetWeight(v int32) *UpdateZoneRecordRequest {
	s.Weight = &v
	return s
}

func (s *UpdateZoneRecordRequest) Validate() error {
	return dara.Validate(s)
}
