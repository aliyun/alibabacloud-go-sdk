// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRoutineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteRoutineResponseBody
	GetRequestId() *string
	SetStatus(v string) *DeleteRoutineResponseBody
	GetStatus() *string
}

type DeleteRoutineResponseBody struct {
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

func (s DeleteRoutineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRoutineResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRoutineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRoutineResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DeleteRoutineResponseBody) SetRequestId(v string) *DeleteRoutineResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRoutineResponseBody) SetStatus(v string) *DeleteRoutineResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteRoutineResponseBody) Validate() error {
	return dara.Validate(s)
}
