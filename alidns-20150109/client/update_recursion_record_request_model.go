// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRecursionRecordRequest
	GetClientToken() *string
	SetPriority(v int32) *UpdateRecursionRecordRequest
	GetPriority() *int32
	SetRecordId(v string) *UpdateRecursionRecordRequest
	GetRecordId() *string
	SetRequestSource(v string) *UpdateRecursionRecordRequest
	GetRequestSource() *string
	SetRr(v string) *UpdateRecursionRecordRequest
	GetRr() *string
	SetTtl(v int32) *UpdateRecursionRecordRequest
	GetTtl() *int32
	SetType(v string) *UpdateRecursionRecordRequest
	GetType() *string
	SetValue(v string) *UpdateRecursionRecordRequest
	GetValue() *string
	SetWeight(v int32) *UpdateRecursionRecordRequest
	GetWeight() *int32
}

type UpdateRecursionRecordRequest struct {
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 5
	Priority *int32 `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// 9*******
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// WebSDK
	RequestSource *string `json:"RequestSource,omitempty" xml:"RequestSource,omitempty"`
	// example:
	//
	// test
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
	// 1.1.XX.XX
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// example:
	//
	// 2
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateRecursionRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRecursionRecordRequest) GetPriority() *int32 {
	return s.Priority
}

func (s *UpdateRecursionRecordRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *UpdateRecursionRecordRequest) GetRequestSource() *string {
	return s.RequestSource
}

func (s *UpdateRecursionRecordRequest) GetRr() *string {
	return s.Rr
}

func (s *UpdateRecursionRecordRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateRecursionRecordRequest) GetType() *string {
	return s.Type
}

func (s *UpdateRecursionRecordRequest) GetValue() *string {
	return s.Value
}

func (s *UpdateRecursionRecordRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateRecursionRecordRequest) SetClientToken(v string) *UpdateRecursionRecordRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetPriority(v int32) *UpdateRecursionRecordRequest {
	s.Priority = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetRecordId(v string) *UpdateRecursionRecordRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetRequestSource(v string) *UpdateRecursionRecordRequest {
	s.RequestSource = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetRr(v string) *UpdateRecursionRecordRequest {
	s.Rr = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetTtl(v int32) *UpdateRecursionRecordRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetType(v string) *UpdateRecursionRecordRequest {
	s.Type = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetValue(v string) *UpdateRecursionRecordRequest {
	s.Value = &v
	return s
}

func (s *UpdateRecursionRecordRequest) SetWeight(v int32) *UpdateRecursionRecordRequest {
	s.Weight = &v
	return s
}

func (s *UpdateRecursionRecordRequest) Validate() error {
	return dara.Validate(s)
}
