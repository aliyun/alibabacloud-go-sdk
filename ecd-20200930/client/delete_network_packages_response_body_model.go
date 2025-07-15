// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNetworkPackagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNetworkPackagesResponseBody
	GetRequestId() *string
}

type DeleteNetworkPackagesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNetworkPackagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNetworkPackagesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNetworkPackagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNetworkPackagesResponseBody) SetRequestId(v string) *DeleteNetworkPackagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkPackagesResponseBody) Validate() error {
	return dara.Validate(s)
}
