// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v int64) *UpdateRecordRemarkResponseBody
	GetRecordId() *int64
	SetRequestId(v string) *UpdateRecordRemarkResponseBody
	GetRequestId() *string
}

type UpdateRecordRemarkResponseBody struct {
	// The ID of the DNS record.
	//
	// example:
	//
	// 202991****
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B7AD377-7E86-44A8-B9A8-53E8666E72FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateRecordRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRecordRemarkResponseBody) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateRecordRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRecordRemarkResponseBody) SetRecordId(v int64) *UpdateRecordRemarkResponseBody {
	s.RecordId = &v
	return s
}

func (s *UpdateRecordRemarkResponseBody) SetRequestId(v string) *UpdateRecordRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRecordRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
