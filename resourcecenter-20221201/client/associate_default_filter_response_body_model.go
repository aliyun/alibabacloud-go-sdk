// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateDefaultFilterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateDefaultFilterResponseBody
	GetRequestId() *string
}

type AssociateDefaultFilterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54673B22-2001-556A-B394-B8697AA9670B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateDefaultFilterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateDefaultFilterResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateDefaultFilterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateDefaultFilterResponseBody) SetRequestId(v string) *AssociateDefaultFilterResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateDefaultFilterResponseBody) Validate() error {
	return dara.Validate(s)
}
