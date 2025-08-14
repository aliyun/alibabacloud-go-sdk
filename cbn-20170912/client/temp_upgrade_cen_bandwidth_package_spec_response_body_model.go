// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTempUpgradeCenBandwidthPackageSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TempUpgradeCenBandwidthPackageSpecResponseBody
	GetRequestId() *string
}

type TempUpgradeCenBandwidthPackageSpecResponseBody struct {
	// example:
	//
	// DB0A026C-A8E5-40AB-977E-3A87DD78F694
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TempUpgradeCenBandwidthPackageSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TempUpgradeCenBandwidthPackageSpecResponseBody) GoString() string {
	return s.String()
}

func (s *TempUpgradeCenBandwidthPackageSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TempUpgradeCenBandwidthPackageSpecResponseBody) SetRequestId(v string) *TempUpgradeCenBandwidthPackageSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *TempUpgradeCenBandwidthPackageSpecResponseBody) Validate() error {
	return dara.Validate(s)
}
