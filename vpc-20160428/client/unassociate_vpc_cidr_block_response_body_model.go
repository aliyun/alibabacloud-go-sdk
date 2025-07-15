// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnassociateVpcCidrBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnassociateVpcCidrBlockResponseBody
	GetRequestId() *string
}

type UnassociateVpcCidrBlockResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// C1221A1F-2ACD-4592-8F27-474E02883159
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnassociateVpcCidrBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnassociateVpcCidrBlockResponseBody) GoString() string {
	return s.String()
}

func (s *UnassociateVpcCidrBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnassociateVpcCidrBlockResponseBody) SetRequestId(v string) *UnassociateVpcCidrBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnassociateVpcCidrBlockResponseBody) Validate() error {
	return dara.Validate(s)
}
