// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamPoolCidrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpamPoolCidrResponseBody
	GetRequestId() *string
}

type DeleteIpamPoolCidrResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F28A239E-F88D-500E-ADE7-FA5E8CA3A170
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpamPoolCidrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamPoolCidrResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolCidrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpamPoolCidrResponseBody) SetRequestId(v string) *DeleteIpamPoolCidrResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpamPoolCidrResponseBody) Validate() error {
	return dara.Validate(s)
}
