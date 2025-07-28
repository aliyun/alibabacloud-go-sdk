// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateZoneRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v int64) *UpdateZoneRecordResponseBody
	GetRecordId() *int64
	SetRequestId(v string) *UpdateZoneRecordResponseBody
	GetRequestId() *string
}

type UpdateZoneRecordResponseBody struct {
	// The ID of the DNS record.
	//
	// example:
	//
	// 172223****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 250E2C38-D0AD-4518-851D-1C1055805F82
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateZoneRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateZoneRecordResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateZoneRecordResponseBody) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateZoneRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateZoneRecordResponseBody) SetRecordId(v int64) *UpdateZoneRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *UpdateZoneRecordResponseBody) SetRequestId(v string) *UpdateZoneRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateZoneRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
