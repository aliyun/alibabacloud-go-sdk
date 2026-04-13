// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociateMacSecKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisassociateMacSecKeyResponseBody
	GetRequestId() *string
}

type DisassociateMacSecKeyResponseBody struct {
	// example:
	//
	// D32B3C26-6C6C-4988-93E9-D2A6444CE6AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisassociateMacSecKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisassociateMacSecKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DisassociateMacSecKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisassociateMacSecKeyResponseBody) SetRequestId(v string) *DisassociateMacSecKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisassociateMacSecKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
