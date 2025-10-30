// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRecoveryContent interface {
	dara.Model
	String() string
	GoString() string
	SetAuthReportInterval(v *AuthReportInterval) *RecoveryContent
	GetAuthReportInterval() *AuthReportInterval
	SetRecoveryActions(v []*string) *RecoveryContent
	GetRecoveryActions() []*string
}

type RecoveryContent struct {
	AuthReportInterval *AuthReportInterval `json:"AuthReportInterval,omitempty" xml:"AuthReportInterval,omitempty"`
	// This parameter is required.
	RecoveryActions []*string `json:"RecoveryActions,omitempty" xml:"RecoveryActions,omitempty" type:"Repeated"`
}

func (s RecoveryContent) String() string {
	return dara.Prettify(s)
}

func (s RecoveryContent) GoString() string {
	return s.String()
}

func (s *RecoveryContent) GetAuthReportInterval() *AuthReportInterval {
	return s.AuthReportInterval
}

func (s *RecoveryContent) GetRecoveryActions() []*string {
	return s.RecoveryActions
}

func (s *RecoveryContent) SetAuthReportInterval(v *AuthReportInterval) *RecoveryContent {
	s.AuthReportInterval = v
	return s
}

func (s *RecoveryContent) SetRecoveryActions(v []*string) *RecoveryContent {
	s.RecoveryActions = v
	return s
}

func (s *RecoveryContent) Validate() error {
	if s.AuthReportInterval != nil {
		if err := s.AuthReportInterval.Validate(); err != nil {
			return err
		}
	}
	return nil
}
