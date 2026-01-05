// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopSlotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopSlotResponseBody
	GetRequestId() *string
}

type StopSlotResponseBody struct {
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopSlotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopSlotResponseBody) GoString() string {
	return s.String()
}

func (s *StopSlotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopSlotResponseBody) SetRequestId(v string) *StopSlotResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopSlotResponseBody) Validate() error {
	return dara.Validate(s)
}
