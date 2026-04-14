// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigSetDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigSetDeleteResponseBody
	GetRequestId() *string
}

type ConfigSetDeleteResponseBody struct {
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigSetDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigSetDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigSetDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigSetDeleteResponseBody) SetRequestId(v string) *ConfigSetDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigSetDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
