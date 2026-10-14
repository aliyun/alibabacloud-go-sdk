// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTrialOrderEligibilityVO interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *TrialOrderEligibilityVO
	GetMessage() *string
	SetValid(v bool) *TrialOrderEligibilityVO
	GetValid() *bool
}

type TrialOrderEligibilityVO struct {
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Valid   *bool   `json:"Valid,omitempty" xml:"Valid,omitempty"`
}

func (s TrialOrderEligibilityVO) String() string {
	return dara.Prettify(s)
}

func (s TrialOrderEligibilityVO) GoString() string {
	return s.String()
}

func (s *TrialOrderEligibilityVO) GetMessage() *string {
	return s.Message
}

func (s *TrialOrderEligibilityVO) GetValid() *bool {
	return s.Valid
}

func (s *TrialOrderEligibilityVO) SetMessage(v string) *TrialOrderEligibilityVO {
	s.Message = &v
	return s
}

func (s *TrialOrderEligibilityVO) SetValid(v bool) *TrialOrderEligibilityVO {
	s.Valid = &v
	return s
}

func (s *TrialOrderEligibilityVO) Validate() error {
	return dara.Validate(s)
}
