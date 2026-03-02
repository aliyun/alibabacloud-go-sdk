// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRuntimeToken interface {
	dara.Model
	String() string
	GoString() string
	SetEnv(v string) *RuntimeToken
	GetEnv() *string
	SetPbcId(v int64) *RuntimeToken
	GetPbcId() *int64
	SetPbcName(v string) *RuntimeToken
	GetPbcName() *string
	SetServiceId(v int64) *RuntimeToken
	GetServiceId() *int64
	SetServiceName(v string) *RuntimeToken
	GetServiceName() *string
	SetToken(v string) *RuntimeToken
	GetToken() *string
}

type RuntimeToken struct {
	// example:
	//
	// TEST
	Env *string `json:"env,omitempty" xml:"env,omitempty"`
	// example:
	//
	// 2
	PbcId *int64 `json:"pbcId,omitempty" xml:"pbcId,omitempty"`
	// example:
	//
	// yunmall-item
	PbcName *string `json:"pbcName,omitempty" xml:"pbcName,omitempty"`
	// example:
	//
	// 1
	ServiceId *int64 `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// example:
	//
	// yunmall-item
	ServiceName *string `json:"serviceName,omitempty" xml:"serviceName,omitempty"`
	// example:
	//
	// dfsdf
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
}

func (s RuntimeToken) String() string {
	return dara.Prettify(s)
}

func (s RuntimeToken) GoString() string {
	return s.String()
}

func (s *RuntimeToken) GetEnv() *string {
	return s.Env
}

func (s *RuntimeToken) GetPbcId() *int64 {
	return s.PbcId
}

func (s *RuntimeToken) GetPbcName() *string {
	return s.PbcName
}

func (s *RuntimeToken) GetServiceId() *int64 {
	return s.ServiceId
}

func (s *RuntimeToken) GetServiceName() *string {
	return s.ServiceName
}

func (s *RuntimeToken) GetToken() *string {
	return s.Token
}

func (s *RuntimeToken) SetEnv(v string) *RuntimeToken {
	s.Env = &v
	return s
}

func (s *RuntimeToken) SetPbcId(v int64) *RuntimeToken {
	s.PbcId = &v
	return s
}

func (s *RuntimeToken) SetPbcName(v string) *RuntimeToken {
	s.PbcName = &v
	return s
}

func (s *RuntimeToken) SetServiceId(v int64) *RuntimeToken {
	s.ServiceId = &v
	return s
}

func (s *RuntimeToken) SetServiceName(v string) *RuntimeToken {
	s.ServiceName = &v
	return s
}

func (s *RuntimeToken) SetToken(v string) *RuntimeToken {
	s.Token = &v
	return s
}

func (s *RuntimeToken) Validate() error {
	return dara.Validate(s)
}
