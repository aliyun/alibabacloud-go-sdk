// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineRelatedRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecordId(v int64) *CreateRoutineRelatedRecordResponseBody
	GetRecordId() *int64
	SetRequestId(v string) *CreateRoutineRelatedRecordResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateRoutineRelatedRecordResponseBody
	GetStatus() *string
}

type CreateRoutineRelatedRecordResponseBody struct {
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// EDBD3EB3-97DA-5465-AEF5-8DCA5DC5E395
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// OK
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateRoutineRelatedRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineRelatedRecordResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoutineRelatedRecordResponseBody) GetRecordId() *int64 {
	return s.RecordId
}

func (s *CreateRoutineRelatedRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRoutineRelatedRecordResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateRoutineRelatedRecordResponseBody) SetRecordId(v int64) *CreateRoutineRelatedRecordResponseBody {
	s.RecordId = &v
	return s
}

func (s *CreateRoutineRelatedRecordResponseBody) SetRequestId(v string) *CreateRoutineRelatedRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoutineRelatedRecordResponseBody) SetStatus(v string) *CreateRoutineRelatedRecordResponseBody {
	s.Status = &v
	return s
}

func (s *CreateRoutineRelatedRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
