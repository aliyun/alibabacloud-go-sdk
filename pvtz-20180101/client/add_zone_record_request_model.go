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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must ensure that it is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
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
	// The DNS request source. Valid values:
	//
	// 	- default: the default resolution line. The default line is equivalent to a global line. We recommend that you configure a default line to ensure that a DNS record can be returned if no intelligent line is matched.
	//
	// 	- Alibaba Cloud lines: indicate that DNS requests are originated from Alibaba Cloud, including Alibaba Cloud public cloud, Alibaba Finance Cloud, and Alibaba Gov Cloud.
	//
	// 	- Custom lines: You can configure custom lines so that Private DNS can return specific IP addresses for DNS requests that are originated from a specific CIDR block.
	//
	// >
	//
	// 	- Only built-in authoritative acceleration zones support custom lines.
	//
	// 	- Set Line to default if you want to choose the default line. Set Line to a specific line code if you want to choose an Alibaba Cloud line or a custom line. Example: aliyun_r_cn-beijing-a.
	//
	// example:
	//
	// default
	Line *string `json:"Line,omitempty" xml:"Line,omitempty"`
	// The priority of the mail exchanger (MX) record. Valid values: **1 to 99**. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The description of the DNS record.
	//
	// example:
	//
	// en
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The hostname. The hostname is the prefix of the subdomain name for the zone. Example: www, @, \\	- (used for wildcard DNS resolution), and mail (used for specifying the mail server that receives emails).
	//
	// For example, if you want to resolve the domain name @.exmaple.com, you must set Rr to @ instead of leaving Rr empty.
	//
	// This parameter is required.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The time to live (TTL) period. Valid values: 5, 30, 60, 3600, 43200, and 86400. Unit: seconds. Default value: 60.
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
	// 114.55.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The weight value of the address. You can set a different weight value for each address. This way, addresses are returned based on the weight values for DNS requests. A weight value must be an integer that ranges from 1 to 100. Default value: 1.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// The zone ID. This ID uniquely identifies the zone.
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
