// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyIpv6InternetBandwidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyIpv6InternetBandwidthResponseBody
	GetRequestId() *string
}

type ModifyIpv6InternetBandwidthResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// D560AF68-4CE8-4A5C-B3FE-469F558094D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyIpv6InternetBandwidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyIpv6InternetBandwidthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyIpv6InternetBandwidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyIpv6InternetBandwidthResponseBody) SetRequestId(v string) *ModifyIpv6InternetBandwidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyIpv6InternetBandwidthResponseBody) Validate() error {
	return dara.Validate(s)
}
