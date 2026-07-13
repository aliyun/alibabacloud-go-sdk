// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRecursionRecordRequest
	GetClientToken() *string
	SetPriority(v int32) *UpdateRecursionRecordRequest
	GetPriority() *int32
	SetRecordId(v string) *UpdateRecursionRecordRequest
	GetRecordId() *string
	SetRequestSource(v string) *UpdateRecursionRecordRequest
	GetRequestSource() *string
	SetRr(v string) *UpdateRecursionRecordRequest
	GetRr() *string
	SetTtl(v int32) *UpdateRecursionRecordRequest
	GetTtl() *int32
	SetType(v string) *UpdateRecursionRecordRequest
	GetType() *string
	SetValue(v string) *UpdateRecursionRecordRequest
	GetValue() *string
	SetWeight(v int32) *UpdateRecursionRecordRequest
	GetWeight() *int32
}

type UpdateRecursionRecordRequest struct {
	// A client token that ensures the idempotence of a request. Generate a unique value for this parameter on your client. The value can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The priority of the MX record. A smaller value indicates a higher priority. The value can be an integer from 1 to 99.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the DNS record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9*******
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The resolution line. The default value is **default**. For more information, see:
	//
	// <props="china">
	//
	// [Lines](https://help.aliyun.com/document_detail/29807.html)
	//
	//
	//
	// <props="intl">
	//
	// [Lines](https://www.alibabacloud.com/help/en/doc-detail/29807.htm)
	//
	// example:
	//
	// WebSDK
	RequestSource *string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty"`
	// The host record. This is the prefix of a domain name. Common prefixes are www, @, \\	- for wildcard DNS, and mail for mailboxes.
	//
	// For example, to resolve @.example.com, set the host record to "@". Do not leave it empty.
	//
	// example:
	//
	// test
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The Time to Live (TTL) in seconds. Only the following values are supported: 5, 30, 60, 3600 (1 hour), 43200 (12 hours), and 86400 (24 hours). The default value is 60.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record. The following types are supported: A: An IPv4 record that maps a domain name to an IPv4 address. AAAA: An IPv6 record that maps a domain name to an IPv6 address. CNAME: An alias record that points a domain name to another domain name. MX: A mail exchanger record that points a domain name to a mail server address. TXT: A text record that contains arbitrary human-readable text. SRV: A service record that identifies a server for a specific service. This is common in directory management for Microsoft systems.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The record value. Enter a value that corresponds to the DNS record type.
	//
	// example:
	//
	// 1.1.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight. An integer from 1 to 100, inclusive. The default value is 1. You can set different weights for each address. DNS queries return addresses in proportion to their weights.
	//
	// example:
	//
	// 2
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateRecursionRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRecursionRecordRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateRecursionRecordRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *UpdateRecursionRecordRequest) GetRequestSource() *string {
	return s.RequestSource
}

func (s *UpdateRecursionRecordRequest) GetRr() *string {
	return s.Rr
}

func (s *UpdateRecursionRecordRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateRecursionRecordRequest) GetType() *string {
	return s.Type
}

func (s *UpdateRecursionRecordRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateRecursionRecordRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateRecursionRecordRequest) SetClientToken(v string) *UpdateRecursionRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetPriority(v int32) *UpdateRecursionRecordRequest {
	s.Priority = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetRecordId(v string) *UpdateRecursionRecordRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetRequestSource(v string) *UpdateRecursionRecordRequest {
	s.RequestSource = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetRr(v string) *UpdateRecursionRecordRequest {
	s.Rr = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetTtl(v int32) *UpdateRecursionRecordRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetType(v string) *UpdateRecursionRecordRequest {
	s.Type = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetValue(v string) *UpdateRecursionRecordRequest {
	s.Value = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetWeight(v int32) *UpdateRecursionRecordRequest {
	s.Weight = &v
	return s
}

func (s *UpdateRecursionRecordRequest) Validate() error {
	return dara.Validate(s)
}
