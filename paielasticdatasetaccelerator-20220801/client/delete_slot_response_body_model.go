// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSlotResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSlotResponseBody
	GetRequestId() *string
}

type DeleteSlotResponseBody struct {
	// example:
	//
	// A731A84D-55C9-44F7-99BB-E1CF0CF19197
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSlotResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSlotResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSlotResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSlotResponseBody) SetRequestId(v string) *DeleteSlotResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSlotResponseBody) Validate() error {
	return dara.Validate(s)
}
