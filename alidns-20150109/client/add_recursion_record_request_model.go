// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecursionRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddRecursionRecordRequest
	GetClientToken() *string
	SetPriority(v int32) *AddRecursionRecordRequest
	GetPriority() *int32
	SetRequestSource(v string) *AddRecursionRecordRequest
	GetRequestSource() *string
	SetRr(v string) *AddRecursionRecordRequest
	GetRr() *string
	SetTtl(v int32) *AddRecursionRecordRequest
	GetTtl() *int32
	SetType(v string) *AddRecursionRecordRequest
	GetType() *string
	SetUserClientIp(v string) *AddRecursionRecordRequest
	GetUserClientIp() *string
	SetValue(v string) *AddRecursionRecordRequest
	GetValue() *string
	SetWeight(v int32) *AddRecursionRecordRequest
	GetWeight() *int32
	SetZoneId(v string) *AddRecursionRecordRequest
	GetZoneId() *string
}

type AddRecursionRecordRequest struct {
	// A client token to ensure the idempotence of the request. Generate a unique value on your client. The token must be unique for each request. It can contain only ASCII characters and must not exceed 64 characters in length.
	//
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The priority of the MX record. A smaller value indicates a higher priority. Valid values: 1 to 99.
	//
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The DNS resolution line. The default value is **default**. For more information, see:
	//
	// <props="china">
	//
	// [DNS resolution lines](https://help.aliyun.com/document_detail/29807.html)
	//
	//
	//
	// <props="intl">
	//
	// [DNS resolution lines](https://www.alibabacloud.com/help/en/doc-detail/29807.htm)
	//
	// example:
	//
	// default
	RequestSource *string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty"`
	// The host record. The host record is the prefix of a domain name. Common examples include www, @, \\	- (for wildcard DNS), and mail (for mailboxes).
	//
	// For example, to resolve @.example.com, set the host record to "@", not an empty string.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The time to live (TTL) in seconds. This is the duration for which the record is cached. Supported values: 5, 30, 60, 3600 (1 hour), 43200 (12 hours), and 86400 (24 hours). Default value: 60.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record. The following record types are supported: A: An IPv4 record that maps a domain name to an IPv4 address. AAAA: An IPv6 record that maps a domain name to an IPv6 address. CNAME: A canonical name record that points a domain name to another domain name. MX: A mail exchanger record that points a domain name to a mail server address. TXT: A text record that contains any human-readable text. SRV: A service record that identifies a server that provides a specific service. This is common in directory management for Microsoft systems. NS: A name server record that delegates a subdomain to another DNS provider for resolution. CAA: A Certification Authority Authorization record that restricts which certification authorities (CAs) can issue certificates for a domain. URL: A URL record that points a domain name to an existing site. SVCB: A service binding record that is used for service discovery. It provides information about supported protocols and service parameters through a DNS record. HTTPS: A record type specific to HTTPS services. An HTTPS record can define secure HTTPS connection protocols and optimal service endpoint addresses.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The client IP address.
	//
	// example:
	//
	// 192.168.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The record value. Enter a value that corresponds to the specified record type.
	//
	// example:
	//
	// 1.1.1.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight of the record. Valid values are integers from 1 to 100. The default value is 1. Set different weights for each address. DNS queries then return addresses based on the specified weight ratio.
	//
	// example:
	//
	// 2
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The ID of the zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// 173671468000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddRecursionRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRecursionRecordRequest) GoString() string {
	return s.String()
}

func (s *AddRecursionRecordRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddRecursionRecordRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *AddRecursionRecordRequest) GetRequestSource() *string {
	return s.RequestSource
}

func (s *AddRecursionRecordRequest) GetRr() *string {
	return s.Rr
}

func (s *AddRecursionRecordRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *AddRecursionRecordRequest) GetType() *string {
	return s.Type
}

func (s *AddRecursionRecordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *AddRecursionRecordRequest) GetValue() *string {
	return s.Value
}

func (s *AddRecursionRecordRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *AddRecursionRecordRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *AddRecursionRecordRequest) SetClientToken(v string) *AddRecursionRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *AddRecursionRecordRequest) SetPriority(v int32) *AddRecursionRecordRequest {
	s.Priority = &v
	return s
}

func (s *AddRecursionRecordRequest) SetRequestSource(v string) *AddRecursionRecordRequest {
	s.RequestSource = &v
	return s
}

func (s *AddRecursionRecordRequest) SetRr(v string) *AddRecursionRecordRequest {
	s.Rr = &v
	return s
}

func (s *AddRecursionRecordRequest) SetTtl(v int32) *AddRecursionRecordRequest {
	s.Ttl = &v
	return s
}

func (s *AddRecursionRecordRequest) SetType(v string) *AddRecursionRecordRequest {
	s.Type = &v
	return s
}

func (s *AddRecursionRecordRequest) SetUserClientIp(v string) *AddRecursionRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddRecursionRecordRequest) SetValue(v string) *AddRecursionRecordRequest {
	s.Value = &v
	return s
}

func (s *AddRecursionRecordRequest) SetWeight(v int32) *AddRecursionRecordRequest {
	s.Weight = &v
	return s
}

func (s *AddRecursionRecordRequest) SetZoneId(v string) *AddRecursionRecordRequest {
	s.ZoneId = &v
	return s
}

func (s *AddRecursionRecordRequest) Validate() error {
	return dara.Validate(s)
}
