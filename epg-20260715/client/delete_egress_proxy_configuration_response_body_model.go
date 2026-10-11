// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyConfigurationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEgressProxyConfigurationResponseBody
	GetRequestId() *string
}

type DeleteEgressProxyConfigurationResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 01A05AB8-CA45-55FC-A2DE-041975EE9AF2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEgressProxyConfigurationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyConfigurationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyConfigurationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEgressProxyConfigurationResponseBody) SetRequestId(v string) *DeleteEgressProxyConfigurationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEgressProxyConfigurationResponseBody) Validate() error {
	return dara.Validate(s)
}
