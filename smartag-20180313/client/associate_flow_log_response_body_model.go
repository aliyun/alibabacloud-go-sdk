// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateFlowLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateFlowLogResponseBody
	GetRequestId() *string
}

type AssociateFlowLogResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// A6E1680B-B34F-4BB7-B504-F8ED675E721C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateFlowLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateFlowLogResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateFlowLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateFlowLogResponseBody) SetRequestId(v string) *AssociateFlowLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateFlowLogResponseBody) Validate() error {
	return dara.Validate(s)
}
