// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIpamResourceDiscoveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIpamResourceDiscoveryId(v string) *CreateIpamResourceDiscoveryResponseBody
	GetIpamResourceDiscoveryId() *string
	SetRequestId(v string) *CreateIpamResourceDiscoveryResponseBody
	GetRequestId() *string
}

type CreateIpamResourceDiscoveryResponseBody struct {
	// The ID of the instance for resource discovery.
	//
	// example:
	//
	// ipam-res-disco-jt5f2af2u6nk2z321****
	IpamResourceDiscoveryId *string `json:"IpamResourceDiscoveryId,omitempty" xml:"IpamResourceDiscoveryId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BB2C39DE-CEB8-595A-981A-F2EFCBE7324E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIpamResourceDiscoveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIpamResourceDiscoveryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIpamResourceDiscoveryResponseBody) GetIpamResourceDiscoveryId() *string {
	return s.IpamResourceDiscoveryId
}

func (s *CreateIpamResourceDiscoveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIpamResourceDiscoveryResponseBody) SetIpamResourceDiscoveryId(v string) *CreateIpamResourceDiscoveryResponseBody {
	s.IpamResourceDiscoveryId = &v
	return s
}

func (s *CreateIpamResourceDiscoveryResponseBody) SetRequestId(v string) *CreateIpamResourceDiscoveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIpamResourceDiscoveryResponseBody) Validate() error {
	return dara.Validate(s)
}
