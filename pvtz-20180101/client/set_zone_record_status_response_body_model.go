// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetZoneRecordStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v int64) *SetZoneRecordStatusResponseBody
	GetRecordId() *int64
	SetRequestId(v string) *SetZoneRecordStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *SetZoneRecordStatusResponseBody
	GetStatus() *string
}

type SetZoneRecordStatusResponseBody struct {
	// The ID of the DNS record.
	//
	// example:
	//
	// 207541****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 39CB16E5-4180-49F2-A060-23C0ECEB80D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The state of the DNS record. Valid values:
	//
	// 	- ENABLE: The DNS record is enabled.
	//
	// 	- DISABLE: The DNS record is disabled.
	//
	// example:
	//
	// DISABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s SetZoneRecordStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetZoneRecordStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetZoneRecordStatusResponseBody) GetRecordId() *int64 {
	return s.RecordId
}

func (s *SetZoneRecordStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetZoneRecordStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *SetZoneRecordStatusResponseBody) SetRecordId(v int64) *SetZoneRecordStatusResponseBody {
	s.RecordId = &v
	return s
}

func (s *SetZoneRecordStatusResponseBody) SetRequestId(v string) *SetZoneRecordStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetZoneRecordStatusResponseBody) SetStatus(v string) *SetZoneRecordStatusResponseBody {
	s.Status = &v
	return s
}

func (s *SetZoneRecordStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
