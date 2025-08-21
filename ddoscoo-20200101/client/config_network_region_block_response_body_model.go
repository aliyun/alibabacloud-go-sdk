// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigNetworkRegionBlockResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigNetworkRegionBlockResponseBody
	GetRequestId() *string
}

type ConfigNetworkRegionBlockResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C33EB3D5-AF96-43CA-9C7E-37A81BC06A1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigNetworkRegionBlockResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigNetworkRegionBlockResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigNetworkRegionBlockResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigNetworkRegionBlockResponseBody) SetRequestId(v string) *ConfigNetworkRegionBlockResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigNetworkRegionBlockResponseBody) Validate() error {
	return dara.Validate(s)
}
