// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSlotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSlotResponseBody
	GetRequestId() *string
}

type UpdateSlotResponseBody struct {
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSlotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSlotResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSlotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSlotResponseBody) SetRequestId(v string) *UpdateSlotResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSlotResponseBody) Validate() error {
	return dara.Validate(s)
}
