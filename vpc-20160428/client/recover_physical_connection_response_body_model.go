// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoverPhysicalConnectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RecoverPhysicalConnectionResponseBody
	GetRequestId() *string
}

type RecoverPhysicalConnectionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// CD14EA74-E9C3-59A9-942A-DFEC7E12818D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RecoverPhysicalConnectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RecoverPhysicalConnectionResponseBody) GoString() string {
	return s.String()
}

func (s *RecoverPhysicalConnectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RecoverPhysicalConnectionResponseBody) SetRequestId(v string) *RecoverPhysicalConnectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *RecoverPhysicalConnectionResponseBody) Validate() error {
	return dara.Validate(s)
}
