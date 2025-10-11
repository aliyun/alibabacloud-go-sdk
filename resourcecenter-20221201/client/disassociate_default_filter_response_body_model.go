// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateDefaultFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisassociateDefaultFilterResponseBody
	GetRequestId() *string
}

type DisassociateDefaultFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BCAB07BA-82FA-5DC0-9322-FB7ED726481D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateDefaultFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociateDefaultFilterResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateDefaultFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociateDefaultFilterResponseBody) SetRequestId(v string) *DisassociateDefaultFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateDefaultFilterResponseBody) Validate() error {
	return dara.Validate(s)
}
