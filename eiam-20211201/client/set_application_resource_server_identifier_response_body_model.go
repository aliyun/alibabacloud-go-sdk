// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetApplicationResourceServerIdentifierResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetApplicationResourceServerIdentifierResponseBody
	GetRequestId() *string
}

type SetApplicationResourceServerIdentifierResponseBody struct {
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetApplicationResourceServerIdentifierResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetApplicationResourceServerIdentifierResponseBody) GoString() string {
	return s.String()
}

func (s *SetApplicationResourceServerIdentifierResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetApplicationResourceServerIdentifierResponseBody) SetRequestId(v string) *SetApplicationResourceServerIdentifierResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetApplicationResourceServerIdentifierResponseBody) Validate() error {
	return dara.Validate(s)
}
