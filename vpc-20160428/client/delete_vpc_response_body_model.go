// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVpcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVpcResponseBody
	GetRequestId() *string
}

type DeleteVpcResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0ED8D006-F706-4D23-88ED-E11ED28DCAC0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVpcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVpcResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVpcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVpcResponseBody) SetRequestId(v string) *DeleteVpcResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVpcResponseBody) Validate() error {
	return dara.Validate(s)
}
