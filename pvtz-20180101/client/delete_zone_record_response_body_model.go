// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteZoneRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v int64) *DeleteZoneRecordResponseBody
	GetRecordId() *int64
	SetRequestId(v string) *DeleteZoneRecordResponseBody
	GetRequestId() *string
}

type DeleteZoneRecordResponseBody struct {
	// The ID of the DNS record.
	//
	// example:
	//
	// 306279****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B7AD377-7E86-44A8-B9A8-53E8666E72FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteZoneRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteZoneRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteZoneRecordResponseBody) GetRecordId() *int64 {
	return s.RecordId
}

func (s *DeleteZoneRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteZoneRecordResponseBody) SetRecordId(v int64) *DeleteZoneRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *DeleteZoneRecordResponseBody) SetRequestId(v string) *DeleteZoneRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteZoneRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
