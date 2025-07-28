// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddZoneRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v int64) *AddZoneRecordResponseBody
	GetRecordId() *int64
	SetRequestId(v string) *AddZoneRecordResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddZoneRecordResponseBody
	GetSuccess() *bool
}

type AddZoneRecordResponseBody struct {
	// The ID of the DNS record.
	//
	// example:
	//
	// 429570****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B7AD377-7E86-44A8-B9A8-53E8666E72FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddZoneRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddZoneRecordResponseBody) GoString() string {
	return s.String()
}

func (s *AddZoneRecordResponseBody) GetRecordId() *int64 {
	return s.RecordId
}

func (s *AddZoneRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddZoneRecordResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddZoneRecordResponseBody) SetRecordId(v int64) *AddZoneRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *AddZoneRecordResponseBody) SetRequestId(v string) *AddZoneRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddZoneRecordResponseBody) SetSuccess(v bool) *AddZoneRecordResponseBody {
	s.Success = &v
	return s
}

func (s *AddZoneRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
