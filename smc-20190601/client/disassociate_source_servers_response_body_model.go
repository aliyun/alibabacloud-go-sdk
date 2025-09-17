// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateSourceServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisassociateSourceServersResponseBody
	GetRequestId() *string
}

type DisassociateSourceServersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 3E8B9ABB-289A-44E6-942D-8AA9E493****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateSourceServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociateSourceServersResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateSourceServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociateSourceServersResponseBody) SetRequestId(v string) *DisassociateSourceServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateSourceServersResponseBody) Validate() error {
	return dara.Validate(s)
}
