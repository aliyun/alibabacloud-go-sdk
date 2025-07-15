// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigADConnectorTrustResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigADConnectorTrustResponseBody
	GetRequestId() *string
}

type ConfigADConnectorTrustResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigADConnectorTrustResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigADConnectorTrustResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigADConnectorTrustResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigADConnectorTrustResponseBody) SetRequestId(v string) *ConfigADConnectorTrustResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigADConnectorTrustResponseBody) Validate() error {
	return dara.Validate(s)
}
