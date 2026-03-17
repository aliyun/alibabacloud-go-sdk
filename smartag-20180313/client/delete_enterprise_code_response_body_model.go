// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnterpriseCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEnterpriseCodeResponseBody
	GetRequestId() *string
}

type DeleteEnterpriseCodeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9DD3DFB2-A9BF-4BEE-9542-661411A9851E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEnterpriseCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnterpriseCodeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEnterpriseCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEnterpriseCodeResponseBody) SetRequestId(v string) *DeleteEnterpriseCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEnterpriseCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
