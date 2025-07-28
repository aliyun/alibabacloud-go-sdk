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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The resolution line. Default value: default.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The priority of the MX record. You can set priorities for different email servers. Valid values: 1 to 99. A smaller value indicates a higher priority.
	//
	// >  This parameter is required if the type of the DNS record is MX.
	//
	// example:
	//
	// 60
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the DNS record. You can call the DescribeZoneRecords operation to query a list of DNS records.
	//
	// This parameter is required.
	//
	// example:
	//
	// 172223****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The hostname. The hostname is the prefix of the subdomain name for zone. Example: www, @, \\	- (used for wildcard DNS resolution), and mail (used for specifying the mail server that receives emails).
	//
	// For example, if you want to resolve the domain name @.exmaple.com, you must set Rr to @ instead of leaving Rr empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The TTL period. Valid values: 5, 30, 60, 3600, 43200, and 86400. Unit: seconds.
	//
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// The type of the DNS record. Valid values:
	//
	// 	- **A**: An A record maps a domain name to an IPv4 address in the dotted decimal notation format.
	//
	// 	- **AAAA**: An AAAA record maps a domain name to an IPv6 address.
	//
	// 	- **CNAME**: A canonical name (CNAME) record maps a domain name to another domain name.
	//
	// 	- **TXT**: A text (TXT) record usually serves as a Sender Policy Framework (SPF) record to prevent email spam. The record value of the TXT record can be up to 255 characters in length.
	//
	// 	- **MX**: A mail exchanger (MX) record maps a domain name to the domain name of a mail server.
	//
	// 	- **PTR**: A pointer (PTR) record maps an IP address to a domain name.
	//
	// 	- **SRV**: A service (SRV) record specifies a server that hosts a specific service. Enter a record value in the format of Priority Weight Port Destination domain name. Separate these items with spaces.
	//
	// >  Before you add a PTR record, you must configure a reverse lookup zone. For more information, see [Add PTR records](https://help.aliyun.com/document_detail/2592976.html).
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
	// The record value. You need to enter the record value based on the DNS record type.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.16.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight value of the address. You can set a different weight value for each address. This way, addresses are returned based on the weight values for DNS requests. A weight value must be an integer that ranges from 1 to 100.
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
