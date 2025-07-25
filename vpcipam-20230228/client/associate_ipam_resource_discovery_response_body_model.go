// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateIpamResourceDiscoveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateIpamResourceDiscoveryResponseBody
	GetRequestId() *string
}

type AssociateIpamResourceDiscoveryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// E897D16A-50EB-543F-B002-C5A26AB818FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateIpamResourceDiscoveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateIpamResourceDiscoveryResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateIpamResourceDiscoveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateIpamResourceDiscoveryResponseBody) SetRequestId(v string) *AssociateIpamResourceDiscoveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateIpamResourceDiscoveryResponseBody) Validate() error {
	return dara.Validate(s)
}
