// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortCcAttackTopIPRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIp(v string) *DescribePortCcAttackTopIPRequest
	GetIp() *string
	SetLimit(v int64) *DescribePortCcAttackTopIPRequest
	GetLimit() *int64
	SetPort(v string) *DescribePortCcAttackTopIPRequest
	GetPort() *string
	SetStartTimestamp(v int64) *DescribePortCcAttackTopIPRequest
	GetStartTimestamp() *int64
}

type DescribePortCcAttackTopIPRequest struct {
	// The IP address of the Anti-DDoS Pro or Anti-DDoS Premium instance to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// 203.107.XX.XX
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The maximum number of entries to return.
	//
	// example:
	//
	// 10
	Limit *int64 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The attacked port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6663
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// The beginning of the time range to query. Unit: seconds.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1678017453
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s DescribePortCcAttackTopIPRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortCcAttackTopIPRequest) GoString() string {
	return s.String()
}

func (s *DescribePortCcAttackTopIPRequest) GetIp() *string {
	return s.Ip
}

func (s *DescribePortCcAttackTopIPRequest) GetLimit() *int64 {
	return s.Limit
}

func (s *DescribePortCcAttackTopIPRequest) GetPort() *string {
	return s.Port
}

func (s *DescribePortCcAttackTopIPRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *DescribePortCcAttackTopIPRequest) SetIp(v string) *DescribePortCcAttackTopIPRequest {
	s.Ip = &v
	return s
}

func (s *DescribePortCcAttackTopIPRequest) SetLimit(v int64) *DescribePortCcAttackTopIPRequest {
	s.Limit = &v
	return s
}

func (s *DescribePortCcAttackTopIPRequest) SetPort(v string) *DescribePortCcAttackTopIPRequest {
	s.Port = &v
	return s
}

func (s *DescribePortCcAttackTopIPRequest) SetStartTimestamp(v int64) *DescribePortCcAttackTopIPRequest {
	s.StartTimestamp = &v
	return s
}

func (s *DescribePortCcAttackTopIPRequest) Validate() error {
	return dara.Validate(s)
}
