// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigInstanceNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigInstanceNetworkResponseBody
	GetRequestId() *string
}

type ConfigInstanceNetworkResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigInstanceNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigInstanceNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigInstanceNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigInstanceNetworkResponseBody) SetRequestId(v string) *ConfigInstanceNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigInstanceNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
