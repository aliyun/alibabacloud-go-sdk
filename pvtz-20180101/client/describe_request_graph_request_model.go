// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRequestGraphRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *DescribeRequestGraphRequest
	GetBizId() *string
	SetBizType(v string) *DescribeRequestGraphRequest
	GetBizType() *string
	SetEndTimestamp(v int64) *DescribeRequestGraphRequest
	GetEndTimestamp() *int64
	SetLang(v string) *DescribeRequestGraphRequest
	GetLang() *string
	SetStartTimestamp(v int64) *DescribeRequestGraphRequest
	GetStartTimestamp() *int64
	SetUserClientIp(v string) *DescribeRequestGraphRequest
	GetUserClientIp() *string
	SetVpcId(v string) *DescribeRequestGraphRequest
	GetVpcId() *string
	SetZoneId(v string) *DescribeRequestGraphRequest
	GetZoneId() *string
}

type DescribeRequestGraphRequest struct {
	// The data ID. This parameter is used with BizType.
	//
	// - If BizType is set to AUTH_ZONE, set this parameter to the zone ID.
	//
	// - If BizType is set to RESOLVER_RULE, set this parameter to the forwarding rule ID.
	//
	// example:
	//
	// b9c93a8954c4098731e863c04302f45a
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The business type. Valid values:
	//
	// - AUTH_ZONE: An authoritative zone.
	//
	// - RESOLVER_RULE: A forwarding rule.
	//
	// example:
	//
	// AUTH_ZONE
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The end of the time range to query. This is a UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1721723624000
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// The language of the response. Valid values:
	//
	// - zh: Chinese.
	//
	// - en: English.
	//
	// Default value: en.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The start of the time range to query. This is a UNIX timestamp.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1721036404000
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
	// The user\\"s IP address.
	//
	// example:
	//
	// 192.168.XX.XX
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// The ID of the Virtual Private Cloud (VPC).
	//
	// example:
	//
	// vpc-f8zvrvr1payllgz38****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
	//
	// > To query the number of requests for a zone, specify either ZoneId or both BizType and BizId.
	//
	// example:
	//
	// 29c752a01cd281a20ddcfa****
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeRequestGraphRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRequestGraphRequest) GoString() string {
	return s.String()
}

func (s *DescribeRequestGraphRequest) GetBizId() *string {
	return s.BizId
}

func (s *DescribeRequestGraphRequest) GetBizType() *string {
	return s.BizType
}

func (s *DescribeRequestGraphRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *DescribeRequestGraphRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeRequestGraphRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribeRequestGraphRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DescribeRequestGraphRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeRequestGraphRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeRequestGraphRequest) SetBizId(v string) *DescribeRequestGraphRequest {
	s.BizId = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetBizType(v string) *DescribeRequestGraphRequest {
	s.BizType = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetEndTimestamp(v int64) *DescribeRequestGraphRequest {
	s.EndTimestamp = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetLang(v string) *DescribeRequestGraphRequest {
	s.Lang = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetStartTimestamp(v int64) *DescribeRequestGraphRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetUserClientIp(v string) *DescribeRequestGraphRequest {
	s.UserClientIp = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetVpcId(v string) *DescribeRequestGraphRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeRequestGraphRequest) SetZoneId(v string) *DescribeRequestGraphRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeRequestGraphRequest) Validate() error {
	return dara.Validate(s)
}
