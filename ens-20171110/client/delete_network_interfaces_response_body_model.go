// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkInterfacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkInterfacesResponseBody
	GetRequestId() *string
}

type DeleteNetworkInterfacesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 125B04C7-3D0D-4245-AF96-14E3758E3F06
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkInterfacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkInterfacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkInterfacesResponseBody) SetRequestId(v string) *DeleteNetworkInterfacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkInterfacesResponseBody) Validate() error {
	return dara.Validate(s)
}
