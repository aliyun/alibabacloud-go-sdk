// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineRelatedRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRoutineRelatedRecordResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteRoutineRelatedRecordResponseBody
	GetStatus() *string
}

type DeleteRoutineRelatedRecordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteRoutineRelatedRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineRelatedRecordResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoutineRelatedRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRoutineRelatedRecordResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteRoutineRelatedRecordResponseBody) SetRequestId(v string) *DeleteRoutineRelatedRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRoutineRelatedRecordResponseBody) SetStatus(v string) *DeleteRoutineRelatedRecordResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteRoutineRelatedRecordResponseBody) Validate() error {
	return dara.Validate(s)
}
