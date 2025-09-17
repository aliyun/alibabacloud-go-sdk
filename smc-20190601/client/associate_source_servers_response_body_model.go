// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateSourceServersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateSourceServersResponseBody
	GetRequestId() *string
}

type AssociateSourceServersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C8B26B44-0189-443E-9816-D951F59623A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateSourceServersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateSourceServersResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateSourceServersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateSourceServersResponseBody) SetRequestId(v string) *AssociateSourceServersResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateSourceServersResponseBody) Validate() error {
	return dara.Validate(s)
}
