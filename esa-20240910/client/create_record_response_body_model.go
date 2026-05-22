// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v int64) *CreateRecordResponseBody
	GetRecordId() *int64
	SetRequestId(v string) *CreateRecordResponseBody
	GetRequestId() *string
}

type CreateRecordResponseBody struct {
	// The record ID.
	//
	// example:
	//
	// 1234567890123
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F61CDR30-E83C-4FDA-BF73-9A94CDD44229
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRecordResponseBody) GetRecordId() *int64 {
	return s.RecordId
}

func (s *CreateRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRecordResponseBody) SetRecordId(v int64) *CreateRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *CreateRecordResponseBody) SetRequestId(v string) *CreateRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
