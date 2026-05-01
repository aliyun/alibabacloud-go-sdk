// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigRuntimeChannelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigRuntimeChannelResponseBody
	GetRequestId() *string
}

type ConfigRuntimeChannelResponseBody struct {
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigRuntimeChannelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigRuntimeChannelResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigRuntimeChannelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigRuntimeChannelResponseBody) SetRequestId(v string) *ConfigRuntimeChannelResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigRuntimeChannelResponseBody) Validate() error {
	return dara.Validate(s)
}
