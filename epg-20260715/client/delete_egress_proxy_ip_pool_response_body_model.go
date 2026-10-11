// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyIpPoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEgressProxyIpPoolResponseBody
	GetRequestId() *string
}

type DeleteEgressProxyIpPoolResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 01A05ABB-C838-53FE-A5BB-D8475117BE6D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEgressProxyIpPoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyIpPoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyIpPoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEgressProxyIpPoolResponseBody) SetRequestId(v string) *DeleteEgressProxyIpPoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEgressProxyIpPoolResponseBody) Validate() error {
	return dara.Validate(s)
}
