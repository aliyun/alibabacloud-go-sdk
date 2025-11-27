// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchOverMajorVersionUpgradeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchOverMajorVersionUpgradeResponseBody
	GetRequestId() *string
}

type SwitchOverMajorVersionUpgradeResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 17F57FEE-EA4F-4337-8D2E-9C23CAA63D74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchOverMajorVersionUpgradeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchOverMajorVersionUpgradeResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchOverMajorVersionUpgradeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchOverMajorVersionUpgradeResponseBody) SetRequestId(v string) *SwitchOverMajorVersionUpgradeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchOverMajorVersionUpgradeResponseBody) Validate() error {
	return dara.Validate(s)
}
