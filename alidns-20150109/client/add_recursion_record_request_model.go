// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRecursionRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AddRecursionRecordRequest
	GetClientToken() *string
	SetPriority(v int32) *AddRecursionRecordRequest
	GetPriority() *int32
	SetRequestSource(v string) *AddRecursionRecordRequest
	GetRequestSource() *string
	SetRr(v string) *AddRecursionRecordRequest
	GetRr() *string
	SetTtl(v int32) *AddRecursionRecordRequest
	GetTtl() *int32
	SetType(v string) *AddRecursionRecordRequest
	GetType() *string
	SetUserClientIp(v string) *AddRecursionRecordRequest
	GetUserClientIp() *string
	SetValue(v string) *AddRecursionRecordRequest
	GetValue() *string
	SetWeight(v int32) *AddRecursionRecordRequest
	GetWeight() *int32
	SetZoneId(v string) *AddRecursionRecordRequest
	GetZoneId() *string
}

type AddRecursionRecordRequest struct {
	// example:
	//
	// 6447728c8578e66aacf062d2df4446dc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 1
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// default
	RequestSource *string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty"`
	// example:
	//
	// www
	Rr *string `json:"Rr,omitempty" xml:"Rr,omitempty"`
	// example:
	//
	// 60
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// example:
	//
	// A
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 192.168.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
	// example:
	//
	// 1.1.1.1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 2
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
	// Zone ID。
	//
	// example:
	//
	// 173671468000011
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s AddRecursionRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s AddRecursionRecordRequest) GoString() string {
	return s.String()
}

func (s *AddRecursionRecordRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddRecursionRecordRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *AddRecursionRecordRequest) GetRequestSource() *string {
	return s.RequestSource
}

func (s *AddRecursionRecordRequest) GetRr() *string {
	return s.Rr
}

func (s *AddRecursionRecordRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *AddRecursionRecordRequest) GetType() *string {
	return s.Type
}

func (s *AddRecursionRecordRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *AddRecursionRecordRequest) GetValue() *string {
	return s.Value
}

func (s *AddRecursionRecordRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *AddRecursionRecordRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *AddRecursionRecordRequest) SetClientToken(v string) *AddRecursionRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *AddRecursionRecordRequest) SetPriority(v int32) *AddRecursionRecordRequest {
	s.Priority = &v
	return s
}

func (s *AddRecursionRecordRequest) SetRequestSource(v string) *AddRecursionRecordRequest {
	s.RequestSource = &v
	return s
}

func (s *AddRecursionRecordRequest) SetRr(v string) *AddRecursionRecordRequest {
	s.Rr = &v
	return s
}

func (s *AddRecursionRecordRequest) SetTtl(v int32) *AddRecursionRecordRequest {
	s.Ttl = &v
	return s
}

func (s *AddRecursionRecordRequest) SetType(v string) *AddRecursionRecordRequest {
	s.Type = &v
	return s
}

func (s *AddRecursionRecordRequest) SetUserClientIp(v string) *AddRecursionRecordRequest {
	s.UserClientIp = &v
	return s
}

func (s *AddRecursionRecordRequest) SetValue(v string) *AddRecursionRecordRequest {
	s.Value = &v
	return s
}

func (s *AddRecursionRecordRequest) SetWeight(v int32) *AddRecursionRecordRequest {
	s.Weight = &v
	return s
}

func (s *AddRecursionRecordRequest) SetZoneId(v string) *AddRecursionRecordRequest {
	s.ZoneId = &v
	return s
}

func (s *AddRecursionRecordRequest) Validate() error {
	return dara.Validate(s)
}
