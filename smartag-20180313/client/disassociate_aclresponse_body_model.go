// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateACLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisassociateACLResponseBody
	GetRequestId() *string
}

type DisassociateACLResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// B28A9AB3-05BD-4A88-AB6A-A555A0BF70C0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateACLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociateACLResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateACLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociateACLResponseBody) SetRequestId(v string) *DisassociateACLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateACLResponseBody) Validate() error {
	return dara.Validate(s)
}
