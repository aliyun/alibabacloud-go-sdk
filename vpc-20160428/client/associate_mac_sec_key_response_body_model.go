// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateMacSecKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateMacSecKeyResponseBody
	GetRequestId() *string
}

type AssociateMacSecKeyResponseBody struct {
	// example:
	//
	// D32B3C26-6C6C-4988-93E9-D2A6444CE6AE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateMacSecKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateMacSecKeyResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateMacSecKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateMacSecKeyResponseBody) SetRequestId(v string) *AssociateMacSecKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateMacSecKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
