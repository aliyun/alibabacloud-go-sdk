// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTriggerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTriggerResponseBody
	GetRequestId() *string
}

type DeleteTriggerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FEDC9B1F-30F2-4C1F-8ED2-B7860187****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTriggerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTriggerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTriggerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTriggerResponseBody) SetRequestId(v string) *DeleteTriggerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTriggerResponseBody) Validate() error {
	return dara.Validate(s)
}
