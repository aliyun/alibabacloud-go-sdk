// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogType(v string) *DescribeLogConfigRequest
	GetLogType() *string
	SetSecurityToken(v string) *DescribeLogConfigRequest
	GetSecurityToken() *string
}

type DescribeLogConfigRequest struct {
	// The log type.
	//
	// Valid values:
	//
	// 	- PROVIDER
	//
	// example:
	//
	// PROVIDER
	LogType       *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DescribeLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeLogConfigRequest) GetLogType() *string {
	return s.LogType
}

func (s *DescribeLogConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DescribeLogConfigRequest) SetLogType(v string) *DescribeLogConfigRequest {
	s.LogType = &v
	return s
}

func (s *DescribeLogConfigRequest) SetSecurityToken(v string) *DescribeLogConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DescribeLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
