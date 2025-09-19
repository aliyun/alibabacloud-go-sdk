// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckTrialFixCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanFix(v bool) *CheckTrialFixCountResponseBody
	GetCanFix() *bool
	SetExpendCount(v int32) *CheckTrialFixCountResponseBody
	GetExpendCount() *int32
	SetRemainCount(v int32) *CheckTrialFixCountResponseBody
	GetRemainCount() *int32
	SetRepairedCount(v int32) *CheckTrialFixCountResponseBody
	GetRepairedCount() *int32
	SetRequestId(v string) *CheckTrialFixCountResponseBody
	GetRequestId() *string
	SetIsTrial(v bool) *CheckTrialFixCountResponseBody
	GetIsTrial() *bool
}

type CheckTrialFixCountResponseBody struct {
	// Indicates whether the vulnerability can be fixed. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	CanFix *bool `json:"CanFix,omitempty" xml:"CanFix,omitempty"`
	// The quota usage required for the current fix operation.
	//
	// example:
	//
	// 10
	ExpendCount *int32 `json:"ExpendCount,omitempty" xml:"ExpendCount,omitempty"`
	// The quota that remains after the current fix operation is complete.
	//
	// example:
	//
	// 0
	RemainCount *int32 `json:"RemainCount,omitempty" xml:"RemainCount,omitempty"`
	// The number of the vulnerabilities that are fixed.
	//
	// example:
	//
	// 10
	RepairedCount *int32 `json:"RepairedCount,omitempty" xml:"RepairedCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether Security Center is in free trial. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	IsTrial *bool `json:"isTrial,omitempty" xml:"isTrial,omitempty"`
}

func (s CheckTrialFixCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckTrialFixCountResponseBody) GoString() string {
	return s.String()
}

func (s *CheckTrialFixCountResponseBody) GetCanFix() *bool {
	return s.CanFix
}

func (s *CheckTrialFixCountResponseBody) GetExpendCount() *int32 {
	return s.ExpendCount
}

func (s *CheckTrialFixCountResponseBody) GetRemainCount() *int32 {
	return s.RemainCount
}

func (s *CheckTrialFixCountResponseBody) GetRepairedCount() *int32 {
	return s.RepairedCount
}

func (s *CheckTrialFixCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckTrialFixCountResponseBody) GetIsTrial() *bool {
	return s.IsTrial
}

func (s *CheckTrialFixCountResponseBody) SetCanFix(v bool) *CheckTrialFixCountResponseBody {
	s.CanFix = &v
	return s
}

func (s *CheckTrialFixCountResponseBody) SetExpendCount(v int32) *CheckTrialFixCountResponseBody {
	s.ExpendCount = &v
	return s
}

func (s *CheckTrialFixCountResponseBody) SetRemainCount(v int32) *CheckTrialFixCountResponseBody {
	s.RemainCount = &v
	return s
}

func (s *CheckTrialFixCountResponseBody) SetRepairedCount(v int32) *CheckTrialFixCountResponseBody {
	s.RepairedCount = &v
	return s
}

func (s *CheckTrialFixCountResponseBody) SetRequestId(v string) *CheckTrialFixCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckTrialFixCountResponseBody) SetIsTrial(v bool) *CheckTrialFixCountResponseBody {
	s.IsTrial = &v
	return s
}

func (s *CheckTrialFixCountResponseBody) Validate() error {
	return dara.Validate(s)
}
