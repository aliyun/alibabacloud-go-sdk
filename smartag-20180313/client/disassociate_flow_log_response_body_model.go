// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateFlowLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisassociateFlowLogResponseBody
	GetRequestId() *string
}

type DisassociateFlowLogResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// DD864F79-FDAC-4909-AFAF-84D1DCAF30A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateFlowLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociateFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateFlowLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociateFlowLogResponseBody) SetRequestId(v string) *DisassociateFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateFlowLogResponseBody) Validate() error {
	return dara.Validate(s)
}
