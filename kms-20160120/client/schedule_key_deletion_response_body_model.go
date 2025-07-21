// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScheduleKeyDeletionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ScheduleKeyDeletionResponseBody
	GetRequestId() *string
}

type ScheduleKeyDeletionResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 3da5b8cc-8107-40ac-a170-793cd181d7b7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScheduleKeyDeletionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScheduleKeyDeletionResponseBody) GoString() string {
	return s.String()
}

func (s *ScheduleKeyDeletionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScheduleKeyDeletionResponseBody) SetRequestId(v string) *ScheduleKeyDeletionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScheduleKeyDeletionResponseBody) Validate() error {
	return dara.Validate(s)
}
