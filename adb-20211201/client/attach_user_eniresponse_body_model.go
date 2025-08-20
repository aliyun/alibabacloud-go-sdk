// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachUserENIResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AttachUserENIResponseBody
	GetRequestId() *string
}

type AttachUserENIResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AttachUserENIResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AttachUserENIResponseBody) GoString() string {
	return s.String()
}

func (s *AttachUserENIResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AttachUserENIResponseBody) SetRequestId(v string) *AttachUserENIResponseBody {
	s.RequestId = &v
	return s
}

func (s *AttachUserENIResponseBody) Validate() error {
	return dara.Validate(s)
}
