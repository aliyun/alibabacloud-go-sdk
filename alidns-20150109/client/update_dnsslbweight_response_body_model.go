// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDNSSLBWeightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v string) *UpdateDNSSLBWeightResponseBody
	GetRecordId() *string
	SetRequestId(v string) *UpdateDNSSLBWeightResponseBody
	GetRequestId() *string
	SetWeight(v int32) *UpdateDNSSLBWeightResponseBody
	GetWeight() *int32
}

type UpdateDNSSLBWeightResponseBody struct {
	// The ID of the DNS record.
	//
	// example:
	//
	// 9999985
	RecordId *string `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 536E9CAD-DB30-4647-AC87-AA5CC38C5382
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The updated weight.
	//
	// example:
	//
	// 2
	Weight *int32 `json:"Weight,omitempty" xml:"Weight,omitempty"`
}

func (s UpdateDNSSLBWeightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDNSSLBWeightResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDNSSLBWeightResponseBody) GetRecordId() *string {
	return s.RecordId
}

func (s *UpdateDNSSLBWeightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDNSSLBWeightResponseBody) GetWeight() *int32 {
	return s.Weight
}

func (s *UpdateDNSSLBWeightResponseBody) SetRecordId(v string) *UpdateDNSSLBWeightResponseBody {
	s.RecordId = &v
	return s
}

func (s *UpdateDNSSLBWeightResponseBody) SetRequestId(v string) *UpdateDNSSLBWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDNSSLBWeightResponseBody) SetWeight(v int32) *UpdateDNSSLBWeightResponseBody {
	s.Weight = &v
	return s
}

func (s *UpdateDNSSLBWeightResponseBody) Validate() error {
	return dara.Validate(s)
}
