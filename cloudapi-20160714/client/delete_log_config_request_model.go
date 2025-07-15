// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLogType(v string) *DeleteLogConfigRequest
	GetLogType() *string
	SetSecurityToken(v string) *DeleteLogConfigRequest
	GetSecurityToken() *string
}

type DeleteLogConfigRequest struct {
	// The log type. Valid values: **log*	- and **survey**.
	//
	// example:
	//
	// log
	LogType       *string `json:"LogType,omitempty" xml:"LogType,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s DeleteLogConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteLogConfigRequest) GetLogType() *string {
	return s.LogType
}

func (s *DeleteLogConfigRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *DeleteLogConfigRequest) SetLogType(v string) *DeleteLogConfigRequest {
	s.LogType = &v
	return s
}

func (s *DeleteLogConfigRequest) SetSecurityToken(v string) *DeleteLogConfigRequest {
	s.SecurityToken = &v
	return s
}

func (s *DeleteLogConfigRequest) Validate() error {
	return dara.Validate(s)
}
