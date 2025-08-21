// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigWebIpSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigWebIpSetResponseBody
	GetRequestId() *string
}

type ConfigWebIpSetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigWebIpSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigWebIpSetResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigWebIpSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigWebIpSetResponseBody) SetRequestId(v string) *ConfigWebIpSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigWebIpSetResponseBody) Validate() error {
	return dara.Validate(s)
}
