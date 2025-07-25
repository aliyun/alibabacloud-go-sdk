// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamPoolAllocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpamPoolAllocationResponseBody
	GetRequestId() *string
}

type DeleteIpamPoolAllocationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B90776C8-F703-51D5-893A-AD1CA699D535
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpamPoolAllocationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamPoolAllocationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolAllocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpamPoolAllocationResponseBody) SetRequestId(v string) *DeleteIpamPoolAllocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpamPoolAllocationResponseBody) Validate() error {
	return dara.Validate(s)
}
