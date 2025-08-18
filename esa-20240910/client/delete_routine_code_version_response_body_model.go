// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineCodeVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRoutineCodeVersionResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteRoutineCodeVersionResponseBody
	GetStatus() *string
}

type DeleteRoutineCodeVersionResponseBody struct {
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

func (s DeleteRoutineCodeVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineCodeVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoutineCodeVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRoutineCodeVersionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteRoutineCodeVersionResponseBody) SetRequestId(v string) *DeleteRoutineCodeVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRoutineCodeVersionResponseBody) SetStatus(v string) *DeleteRoutineCodeVersionResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteRoutineCodeVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
