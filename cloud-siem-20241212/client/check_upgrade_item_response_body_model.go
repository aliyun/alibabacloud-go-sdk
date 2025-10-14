// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckUpgradeItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckUpgradeItemResponseBody
	GetRequestId() *string
	SetUpgradeItem(v *CheckUpgradeItemResponseBodyUpgradeItem) *CheckUpgradeItemResponseBody
	GetUpgradeItem() *CheckUpgradeItemResponseBodyUpgradeItem
}

type CheckUpgradeItemResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId   *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UpgradeItem *CheckUpgradeItemResponseBodyUpgradeItem `json:"UpgradeItem,omitempty" xml:"UpgradeItem,omitempty" type:"Struct"`
}

func (s CheckUpgradeItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckUpgradeItemResponseBody) GoString() string {
	return s.String()
}

func (s *CheckUpgradeItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckUpgradeItemResponseBody) GetUpgradeItem() *CheckUpgradeItemResponseBodyUpgradeItem {
	return s.UpgradeItem
}

func (s *CheckUpgradeItemResponseBody) SetRequestId(v string) *CheckUpgradeItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckUpgradeItemResponseBody) SetUpgradeItem(v *CheckUpgradeItemResponseBodyUpgradeItem) *CheckUpgradeItemResponseBody {
	s.UpgradeItem = v
	return s
}

func (s *CheckUpgradeItemResponseBody) Validate() error {
	if s.UpgradeItem != nil {
		if err := s.UpgradeItem.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckUpgradeItemResponseBodyUpgradeItem struct {
	// example:
	//
	// OK
	CheckResult *string `json:"CheckResult,omitempty" xml:"CheckResult,omitempty"`
	// example:
	//
	// success
	CheckStatus *string `json:"CheckStatus,omitempty" xml:"CheckStatus,omitempty"`
	// example:
	//
	// incident_upgrade
	UpgradeItemId *string `json:"UpgradeItemId,omitempty" xml:"UpgradeItemId,omitempty"`
}

func (s CheckUpgradeItemResponseBodyUpgradeItem) String() string {
	return dara.Prettify(s)
}

func (s CheckUpgradeItemResponseBodyUpgradeItem) GoString() string {
	return s.String()
}

func (s *CheckUpgradeItemResponseBodyUpgradeItem) GetCheckResult() *string {
	return s.CheckResult
}

func (s *CheckUpgradeItemResponseBodyUpgradeItem) GetCheckStatus() *string {
	return s.CheckStatus
}

func (s *CheckUpgradeItemResponseBodyUpgradeItem) GetUpgradeItemId() *string {
	return s.UpgradeItemId
}

func (s *CheckUpgradeItemResponseBodyUpgradeItem) SetCheckResult(v string) *CheckUpgradeItemResponseBodyUpgradeItem {
	s.CheckResult = &v
	return s
}

func (s *CheckUpgradeItemResponseBodyUpgradeItem) SetCheckStatus(v string) *CheckUpgradeItemResponseBodyUpgradeItem {
	s.CheckStatus = &v
	return s
}

func (s *CheckUpgradeItemResponseBodyUpgradeItem) SetUpgradeItemId(v string) *CheckUpgradeItemResponseBodyUpgradeItem {
	s.UpgradeItemId = &v
	return s
}

func (s *CheckUpgradeItemResponseBodyUpgradeItem) Validate() error {
	return dara.Validate(s)
}
