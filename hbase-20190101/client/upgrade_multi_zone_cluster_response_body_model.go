// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMultiZoneClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeMultiZoneClusterResponseBody
	GetRequestId() *string
	SetUpgradingComponents(v string) *UpgradeMultiZoneClusterResponseBody
	GetUpgradingComponents() *string
}

type UpgradeMultiZoneClusterResponseBody struct {
	// example:
	//
	// C532A4D4-9451-4460-BB3E-300FEC852D3F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// LINDORM
	UpgradingComponents *string `json:"UpgradingComponents,omitempty" xml:"UpgradingComponents,omitempty"`
}

func (s UpgradeMultiZoneClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMultiZoneClusterResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMultiZoneClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeMultiZoneClusterResponseBody) GetUpgradingComponents() *string {
	return s.UpgradingComponents
}

func (s *UpgradeMultiZoneClusterResponseBody) SetRequestId(v string) *UpgradeMultiZoneClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeMultiZoneClusterResponseBody) SetUpgradingComponents(v string) *UpgradeMultiZoneClusterResponseBody {
	s.UpgradingComponents = &v
	return s
}

func (s *UpgradeMultiZoneClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
