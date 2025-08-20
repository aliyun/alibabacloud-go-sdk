// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelUpdateStackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CancelUpdateStackResponseBody
	GetRequestId() *string
}

type CancelUpdateStackResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B288A0BE-D927-4888-B0F7-B35EF84B6E6F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelUpdateStackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelUpdateStackResponseBody) GoString() string {
	return s.String()
}

func (s *CancelUpdateStackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelUpdateStackResponseBody) SetRequestId(v string) *CancelUpdateStackResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelUpdateStackResponseBody) Validate() error {
	return dara.Validate(s)
}
