// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TriggerCheckResponseBody
	GetRequestId() *string
}

type TriggerCheckResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 43313389-DED8-5BB7-8CB9-F22CDEB744DE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TriggerCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerCheckResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerCheckResponseBody) SetRequestId(v string) *TriggerCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
