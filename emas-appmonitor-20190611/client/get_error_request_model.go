// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetErrorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppKey(v int64) *GetErrorRequest
	GetAppKey() *int64
	SetBizModule(v string) *GetErrorRequest
	GetBizModule() *string
	SetClientTime(v int64) *GetErrorRequest
	GetClientTime() *int64
	SetDid(v string) *GetErrorRequest
	GetDid() *string
	SetDigestHash(v string) *GetErrorRequest
	GetDigestHash() *string
	SetForce(v bool) *GetErrorRequest
	GetForce() *bool
	SetOs(v string) *GetErrorRequest
	GetOs() *string
	SetUuid(v string) *GetErrorRequest
	GetUuid() *string
}

type GetErrorRequest struct {
	// appKey
	//
	// This parameter is required.
	//
	// example:
	//
	// 233588686
	AppKey *int64 `json:"AppKey,omitempty" xml:"AppKey,omitempty"`
	// example:
	//
	// h5Resource
	BizModule *string `json:"BizModule,omitempty" xml:"BizModule,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1739808000000
	ClientTime *int64 `json:"ClientTime,omitempty" xml:"ClientTime,omitempty"`
	// example:
	//
	// 233588686
	Did        *string `json:"Did,omitempty" xml:"Did,omitempty"`
	DigestHash *string `json:"DigestHash,omitempty" xml:"DigestHash,omitempty"`
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// example:
	//
	// android
	Os *string `json:"Os,omitempty" xml:"Os,omitempty"`
	// example:
	//
	// abcf4a4b-158c-4a0b-b81c-262785d84c4f
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetErrorRequest) String() string {
	return dara.Prettify(s)
}

func (s GetErrorRequest) GoString() string {
	return s.String()
}

func (s *GetErrorRequest) GetAppKey() *int64 {
	return s.AppKey
}

func (s *GetErrorRequest) GetBizModule() *string {
	return s.BizModule
}

func (s *GetErrorRequest) GetClientTime() *int64 {
	return s.ClientTime
}

func (s *GetErrorRequest) GetDid() *string {
	return s.Did
}

func (s *GetErrorRequest) GetDigestHash() *string {
	return s.DigestHash
}

func (s *GetErrorRequest) GetForce() *bool {
	return s.Force
}

func (s *GetErrorRequest) GetOs() *string {
	return s.Os
}

func (s *GetErrorRequest) GetUuid() *string {
	return s.Uuid
}

func (s *GetErrorRequest) SetAppKey(v int64) *GetErrorRequest {
	s.AppKey = &v
	return s
}

func (s *GetErrorRequest) SetBizModule(v string) *GetErrorRequest {
	s.BizModule = &v
	return s
}

func (s *GetErrorRequest) SetClientTime(v int64) *GetErrorRequest {
	s.ClientTime = &v
	return s
}

func (s *GetErrorRequest) SetDid(v string) *GetErrorRequest {
	s.Did = &v
	return s
}

func (s *GetErrorRequest) SetDigestHash(v string) *GetErrorRequest {
	s.DigestHash = &v
	return s
}

func (s *GetErrorRequest) SetForce(v bool) *GetErrorRequest {
	s.Force = &v
	return s
}

func (s *GetErrorRequest) SetOs(v string) *GetErrorRequest {
	s.Os = &v
	return s
}

func (s *GetErrorRequest) SetUuid(v string) *GetErrorRequest {
	s.Uuid = &v
	return s
}

func (s *GetErrorRequest) Validate() error {
	return dara.Validate(s)
}
