// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRoutineResponseBody
	GetRequestId() *string
	SetStatus(v string) *CreateRoutineResponseBody
	GetStatus() *string
}

type CreateRoutineResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateRoutineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRoutineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRoutineResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateRoutineResponseBody) SetRequestId(v string) *CreateRoutineResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRoutineResponseBody) SetStatus(v string) *CreateRoutineResponseBody {
	s.Status = &v
	return s
}

func (s *CreateRoutineResponseBody) Validate() error {
	return dara.Validate(s)
}
