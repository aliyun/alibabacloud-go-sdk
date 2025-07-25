// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIpamResourceDiscoveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIpamResourceDiscoveryResponseBody
	GetRequestId() *string
}

type UpdateIpamResourceDiscoveryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// BB2C39DE-CEB8-595A-981A-F2EFCBE7324E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateIpamResourceDiscoveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIpamResourceDiscoveryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIpamResourceDiscoveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIpamResourceDiscoveryResponseBody) SetRequestId(v string) *UpdateIpamResourceDiscoveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIpamResourceDiscoveryResponseBody) Validate() error {
	return dara.Validate(s)
}
