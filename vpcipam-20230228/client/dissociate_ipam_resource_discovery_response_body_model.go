// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateIpamResourceDiscoveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateIpamResourceDiscoveryResponseBody
	GetRequestId() *string
}

type DissociateIpamResourceDiscoveryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 86137597-443F-5B66-B9B6-8514E0C50B8F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateIpamResourceDiscoveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateIpamResourceDiscoveryResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateIpamResourceDiscoveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateIpamResourceDiscoveryResponseBody) SetRequestId(v string) *DissociateIpamResourceDiscoveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateIpamResourceDiscoveryResponseBody) Validate() error {
	return dara.Validate(s)
}
