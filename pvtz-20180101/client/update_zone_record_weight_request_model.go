// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZoneRecordWeightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateZoneRecordWeightRequest
	GetClientToken() *string
	SetLang(v string) *UpdateZoneRecordWeightRequest
	GetLang() *string
	SetRecordId(v int64) *UpdateZoneRecordWeightRequest
	GetRecordId() *int64
	SetWeight(v int32) *UpdateZoneRecordWeightRequest
	GetWeight() *int32
}

type UpdateZoneRecordWeightRequest struct {
	// The client token that is used to ensure the idempotence of the request. If you do not specify this parameter, the system automatically generates a value. To ensure uniqueness across different requests, the value cannot exceed 64 ASCII characters.
	//
	// example:
	//
	// 210bc45716943908285687176dcf0a
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The language.
	//
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the DNS record.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5808
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The updated weight value `[0,100]`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateZoneRecordWeightRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateZoneRecordWeightRequest) GoString() string {
	return s.String()
}

func (s *UpdateZoneRecordWeightRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateZoneRecordWeightRequest) GetLang() *string {
	return s.Lang
}

func (s *UpdateZoneRecordWeightRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateZoneRecordWeightRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateZoneRecordWeightRequest) SetClientToken(v string) *UpdateZoneRecordWeightRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateZoneRecordWeightRequest) SetLang(v string) *UpdateZoneRecordWeightRequest {
	s.Lang = &v
	return s
}

func (s *UpdateZoneRecordWeightRequest) SetRecordId(v int64) *UpdateZoneRecordWeightRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateZoneRecordWeightRequest) SetWeight(v int32) *UpdateZoneRecordWeightRequest {
	s.Weight = &v
	return s
}

func (s *UpdateZoneRecordWeightRequest) Validate() error {
	return dara.Validate(s)
}
