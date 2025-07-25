// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchDnsGtmInstanceStrategyModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchDnsGtmInstanceStrategyModeResponseBody
	GetRequestId() *string
}

type SwitchDnsGtmInstanceStrategyModeResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6856BCF6-11D6-4D7E-AC53-FD579933522B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchDnsGtmInstanceStrategyModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchDnsGtmInstanceStrategyModeResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchDnsGtmInstanceStrategyModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchDnsGtmInstanceStrategyModeResponseBody) SetRequestId(v string) *SwitchDnsGtmInstanceStrategyModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchDnsGtmInstanceStrategyModeResponseBody) Validate() error {
	return dara.Validate(s)
}
