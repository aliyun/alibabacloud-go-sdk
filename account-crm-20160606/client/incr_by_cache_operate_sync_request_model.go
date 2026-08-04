// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIncrByCacheOperateSyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultValue(v int32) *IncrByCacheOperateSyncRequest
	GetDefaultValue() *int32
	SetExpireSeconds(v int32) *IncrByCacheOperateSyncRequest
	GetExpireSeconds() *int32
	SetKey(v string) *IncrByCacheOperateSyncRequest
	GetKey() *string
	SetStep(v int32) *IncrByCacheOperateSyncRequest
	GetStep() *int32
}

type IncrByCacheOperateSyncRequest struct {
	DefaultValue  *int32  `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	ExpireSeconds *int32  `json:"ExpireSeconds,omitempty" xml:"ExpireSeconds,omitempty"`
	Key           *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Step          *int32  `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s IncrByCacheOperateSyncRequest) String() string {
	return dara.Prettify(s)
}

func (s IncrByCacheOperateSyncRequest) GoString() string {
	return s.String()
}

func (s *IncrByCacheOperateSyncRequest) GetDefaultValue() *int32 {
	return s.DefaultValue
}

func (s *IncrByCacheOperateSyncRequest) GetExpireSeconds() *int32 {
	return s.ExpireSeconds
}

func (s *IncrByCacheOperateSyncRequest) GetKey() *string {
	return s.Key
}

func (s *IncrByCacheOperateSyncRequest) GetStep() *int32 {
	return s.Step
}

func (s *IncrByCacheOperateSyncRequest) SetDefaultValue(v int32) *IncrByCacheOperateSyncRequest {
	s.DefaultValue = &v
	return s
}

func (s *IncrByCacheOperateSyncRequest) SetExpireSeconds(v int32) *IncrByCacheOperateSyncRequest {
	s.ExpireSeconds = &v
	return s
}

func (s *IncrByCacheOperateSyncRequest) SetKey(v string) *IncrByCacheOperateSyncRequest {
	s.Key = &v
	return s
}

func (s *IncrByCacheOperateSyncRequest) SetStep(v int32) *IncrByCacheOperateSyncRequest {
	s.Step = &v
	return s
}

func (s *IncrByCacheOperateSyncRequest) Validate() error {
	return dara.Validate(s)
}
