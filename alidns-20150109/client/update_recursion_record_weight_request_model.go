// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecursionRecordWeightRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateRecursionRecordWeightRequest
	GetClientToken() *string
	SetRecordId(v string) *UpdateRecursionRecordWeightRequest
	GetRecordId() *string
	SetWeight(v int32) *UpdateRecursionRecordWeightRequest
	GetWeight() *int32
}

type UpdateRecursionRecordWeightRequest struct {
	// example:
	//
	// 21079fa016944979537637959d09bc
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// 17363242424
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// example:
	//
	// 1
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateRecursionRecordWeightRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecursionRecordWeightRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecursionRecordWeightRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateRecursionRecordWeightRequest) GetRecordId() *string {
	return s.RecordId
}

func (s *UpdateRecursionRecordWeightRequest) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateRecursionRecordWeightRequest) SetClientToken(v string) *UpdateRecursionRecordWeightRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateRecursionRecordWeightRequest) SetRecordId(v string) *UpdateRecursionRecordWeightRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateRecursionRecordWeightRequest) SetWeight(v int32) *UpdateRecursionRecordWeightRequest {
	s.Weight = &v
	return s
}

func (s *UpdateRecursionRecordWeightRequest) Validate() error {
	return dara.Validate(s)
}
