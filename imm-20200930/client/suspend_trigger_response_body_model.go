// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SuspendTriggerResponseBody
	GetRequestId() *string
}

type SuspendTriggerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0BC1F0C9-8E99-46C6-B502-10DED******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SuspendTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendTriggerResponseBody) SetRequestId(v string) *SuspendTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
