// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeMinorVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeMinorVersionResponseBody
	GetRequestId() *string
	SetUpgradingComponents(v string) *UpgradeMinorVersionResponseBody
	GetUpgradingComponents() *string
}

type UpgradeMinorVersionResponseBody struct {
	// example:
	//
	// 7B8EC240-BB13-4DBC-B955-F90170E82609
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// HADOOP
	UpgradingComponents *string `json:"UpgradingComponents,omitempty" xml:"UpgradingComponents,omitempty"`
}

func (s UpgradeMinorVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeMinorVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeMinorVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeMinorVersionResponseBody) GetUpgradingComponents() *string {
	return s.UpgradingComponents
}

func (s *UpgradeMinorVersionResponseBody) SetRequestId(v string) *UpgradeMinorVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeMinorVersionResponseBody) SetUpgradingComponents(v string) *UpgradeMinorVersionResponseBody {
	s.UpgradingComponents = &v
	return s
}

func (s *UpgradeMinorVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
