// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachUserENIResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DetachUserENIResponseBody
	GetRequestId() *string
}

type DetachUserENIResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DetachUserENIResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DetachUserENIResponseBody) GoString() string {
	return s.String()
}

func (s *DetachUserENIResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DetachUserENIResponseBody) SetRequestId(v string) *DetachUserENIResponseBody {
	s.RequestId = &v
	return s
}

func (s *DetachUserENIResponseBody) Validate() error {
	return dara.Validate(s)
}
