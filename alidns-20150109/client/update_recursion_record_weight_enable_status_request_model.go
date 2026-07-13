// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordWeightEnableStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRecursionRecordWeightEnableStatusRequest
	GetClientToken() *string
	SetEnableStatus(v string) *UpdateRecursionRecordWeightEnableStatusRequest
	GetEnableStatus() *string
	SetRequestSource(v string) *UpdateRecursionRecordWeightEnableStatusRequest
	GetRequestSource() *string
	SetRr(v string) *UpdateRecursionRecordWeightEnableStatusRequest
	GetRr() *string
	SetType(v string) *UpdateRecursionRecordWeightEnableStatusRequest
	GetType() *string
	SetZoneId(v string) *UpdateRecursionRecordWeightEnableStatusRequest
	GetZoneId() *string
}

type UpdateRecursionRecordWeightEnableStatusRequest struct {
	// A client token that is used to ensure the idempotence of a request. The client generates the value of this parameter. The value must be unique for each request and can be up to 64 ASCII characters in length.
	//
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable the weight algorithm. Valid values: \\*\\*enable\\*\\	- and \\*\\*disable\\*\\*.
	//
	// example:
	//
	// enable
	EnableStatus *string `json:"EnableStatus,omitempty" xml:"EnableStatus,omitempty"`
	// The DNS resolution line. The default value is **default**. For more information, see [DNS resolution lines](https://help.aliyun.com/document_detail/29807.html).
	//
	// <props="china">
	//
	// [Resolution Line Enumeration](https://help.aliyun.com/document_detail/29807.html)
	//
	//
	//
	// <props="intl">
	//
	// [Enumeration of DNS record lines](https://www.alibabacloud.com/help/zh/doc-detail/29807.htm)
	//
	// example:
	//
	// default
	RequestSource *string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty"`
	// The host record.
	//
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// The type of the DNS record. The following types are supported: \\*\\*A\\*\\*, which maps a domain name to an IPv4 address. \\*\\*AAAA\\*\\*, which maps a domain name to an IPv6 address. \\*\\*CNAME\\*\\*, an alias record that points a domain name to another domain name. \\*\\*MX\\*\\*, a mail exchanger record that points a domain name to a mail server address. \\*\\*TXT\\*\\*, an arbitrary, human-readable text DNS record. \\*\\*SRV\\*\\*, a service record that identifies a server that provides a specific service, commonly used for directory management in Microsoft systems.
	//
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The zone ID for the domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// 176432424234
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s UpdateRecursionRecordWeightEnableStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordWeightEnableStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) GetEnableStatus() *string {
	return s.EnableStatus
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) GetRequestSource() *string {
	return s.RequestSource
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) GetRr() *string {
	return s.Rr
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) GetType() *string {
	return s.Type
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) SetClientToken(v string) *UpdateRecursionRecordWeightEnableStatusRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) SetEnableStatus(v string) *UpdateRecursionRecordWeightEnableStatusRequest {
	s.EnableStatus = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) SetRequestSource(v string) *UpdateRecursionRecordWeightEnableStatusRequest {
	s.RequestSource = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) SetRr(v string) *UpdateRecursionRecordWeightEnableStatusRequest {
	s.Rr = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) SetType(v string) *UpdateRecursionRecordWeightEnableStatusRequest {
	s.Type = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) SetZoneId(v string) *UpdateRecursionRecordWeightEnableStatusRequest {
	s.ZoneId = &v
	return s
}

func (s *UpdateRecursionRecordWeightEnableStatusRequest) Validate() error {
	return dara.Validate(s)
}
