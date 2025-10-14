// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRoutineConfigDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateRoutineConfigDescriptionResponseBody
	GetRequestId() *string
	SetStatus(v string) *UpdateRoutineConfigDescriptionResponseBody
	GetStatus() *string
}

type UpdateRoutineConfigDescriptionResponseBody struct {
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

func (s UpdateRoutineConfigDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateRoutineConfigDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateRoutineConfigDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateRoutineConfigDescriptionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpdateRoutineConfigDescriptionResponseBody) SetRequestId(v string) *UpdateRoutineConfigDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateRoutineConfigDescriptionResponseBody) SetStatus(v string) *UpdateRoutineConfigDescriptionResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateRoutineConfigDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
