// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZoneRecordWeightResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v int64) *UpdateZoneRecordWeightResponseBody
	GetRecordId() *int64
	SetRequestId(v string) *UpdateZoneRecordWeightResponseBody
	GetRequestId() *string
}

type UpdateZoneRecordWeightResponseBody struct {
	// The ID of the DNS record.
	//
	// example:
	//
	// 5808
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39CB16E5-4180-49F2-A060-23C0ECEB80D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateZoneRecordWeightResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateZoneRecordWeightResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateZoneRecordWeightResponseBody) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateZoneRecordWeightResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateZoneRecordWeightResponseBody) SetRecordId(v int64) *UpdateZoneRecordWeightResponseBody {
	s.RecordId = &v
	return s
}

func (s *UpdateZoneRecordWeightResponseBody) SetRequestId(v string) *UpdateZoneRecordWeightResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateZoneRecordWeightResponseBody) Validate() error {
	return dara.Validate(s)
}
