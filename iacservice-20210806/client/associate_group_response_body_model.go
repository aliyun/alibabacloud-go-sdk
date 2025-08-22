// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateGroupResponseBody
	GetRequestId() *string
}

type AssociateGroupResponseBody struct {
	// example:
	//
	// B6ED9F71-7FA8-598E-B64D-4606FB3FCCC9
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s AssociateGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateGroupResponseBody) SetRequestId(v string) *AssociateGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
