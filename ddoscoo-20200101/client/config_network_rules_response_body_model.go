// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigNetworkRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ConfigNetworkRulesResponseBody
	GetRequestId() *string
}

type ConfigNetworkRulesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CC042262-15A3-4A49-ADF0-130968EA47BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ConfigNetworkRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigNetworkRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigNetworkRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigNetworkRulesResponseBody) SetRequestId(v string) *ConfigNetworkRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigNetworkRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
