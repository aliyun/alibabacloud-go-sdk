// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCacheOperateSyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExceptVersion(v int32) *SetCacheOperateSyncRequest
	GetExceptVersion() *int32
	SetExpireSeconds(v int32) *SetCacheOperateSyncRequest
	GetExpireSeconds() *int32
	SetKey(v string) *SetCacheOperateSyncRequest
	GetKey() *string
	SetSetType(v string) *SetCacheOperateSyncRequest
	GetSetType() *string
	SetValueClazz(v string) *SetCacheOperateSyncRequest
	GetValueClazz() *string
	SetValueString(v string) *SetCacheOperateSyncRequest
	GetValueString() *string
}

type SetCacheOperateSyncRequest struct {
	ExceptVersion *int32  `json:"ExceptVersion,omitempty" xml:"ExceptVersion,omitempty"`
	ExpireSeconds *int32  `json:"ExpireSeconds,omitempty" xml:"ExpireSeconds,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	SetType       *string `json:"SetType,omitempty" xml:"SetType,omitempty"`
	ValueClazz    *string `json:"ValueClazz,omitempty" xml:"ValueClazz,omitempty"`
	ValueString   *string `json:"ValueString,omitempty" xml:"ValueString,omitempty"`
}

func (s SetCacheOperateSyncRequest) String() string {
	return dara.Prettify(s)
}

func (s SetCacheOperateSyncRequest) GoString() string {
	return s.String()
}

func (s *SetCacheOperateSyncRequest) GetExceptVersion() *int32 {
	return s.ExceptVersion
}

func (s *SetCacheOperateSyncRequest) GetExpireSeconds() *int32 {
	return s.ExpireSeconds
}

func (s *SetCacheOperateSyncRequest) GetKey() *string {
	return s.Key
}

func (s *SetCacheOperateSyncRequest) GetSetType() *string {
	return s.SetType
}

func (s *SetCacheOperateSyncRequest) GetValueClazz() *string {
	return s.ValueClazz
}

func (s *SetCacheOperateSyncRequest) GetValueString() *string {
	return s.ValueString
}

func (s *SetCacheOperateSyncRequest) SetExceptVersion(v int32) *SetCacheOperateSyncRequest {
	s.ExceptVersion = &v
	return s
}

func (s *SetCacheOperateSyncRequest) SetExpireSeconds(v int32) *SetCacheOperateSyncRequest {
	s.ExpireSeconds = &v
	return s
}

func (s *SetCacheOperateSyncRequest) SetKey(v string) *SetCacheOperateSyncRequest {
	s.Key = &v
	return s
}

func (s *SetCacheOperateSyncRequest) SetSetType(v string) *SetCacheOperateSyncRequest {
	s.SetType = &v
	return s
}

func (s *SetCacheOperateSyncRequest) SetValueClazz(v string) *SetCacheOperateSyncRequest {
	s.ValueClazz = &v
	return s
}

func (s *SetCacheOperateSyncRequest) SetValueString(v string) *SetCacheOperateSyncRequest {
	s.ValueString = &v
	return s
}

func (s *SetCacheOperateSyncRequest) Validate() error {
	return dara.Validate(s)
}
