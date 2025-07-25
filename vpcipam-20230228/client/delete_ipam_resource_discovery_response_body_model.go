// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamResourceDiscoveryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpamResourceDiscoveryResponseBody
	GetRequestId() *string
}

type DeleteIpamResourceDiscoveryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9F8315CB-560E-5F1E-B069-6E44B440CAF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpamResourceDiscoveryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamResourceDiscoveryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpamResourceDiscoveryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpamResourceDiscoveryResponseBody) SetRequestId(v string) *DeleteIpamResourceDiscoveryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpamResourceDiscoveryResponseBody) Validate() error {
	return dara.Validate(s)
}
