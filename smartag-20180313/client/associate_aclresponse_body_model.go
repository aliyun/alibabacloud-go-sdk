// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateACLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateACLResponseBody
	GetRequestId() *string
}

type AssociateACLResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C9A75915-0260-4335-851A-D866A7ED396C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateACLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateACLResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateACLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateACLResponseBody) SetRequestId(v string) *AssociateACLResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateACLResponseBody) Validate() error {
	return dara.Validate(s)
}
