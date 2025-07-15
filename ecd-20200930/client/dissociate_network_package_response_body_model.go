// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateNetworkPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateNetworkPackageResponseBody
	GetRequestId() *string
}

type DissociateNetworkPackageResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateNetworkPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateNetworkPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateNetworkPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateNetworkPackageResponseBody) SetRequestId(v string) *DissociateNetworkPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateNetworkPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
