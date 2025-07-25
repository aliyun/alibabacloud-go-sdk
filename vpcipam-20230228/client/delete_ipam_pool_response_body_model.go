// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpamPoolResponseBody
	GetRequestId() *string
}

type DeleteIpamPoolResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 57B7DCCA-F192-5528-8AF3-2FE1413228C9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpamPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpamPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpamPoolResponseBody) SetRequestId(v string) *DeleteIpamPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpamPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
