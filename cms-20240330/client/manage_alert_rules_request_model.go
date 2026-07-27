// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageAlertRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *ManageAlertRulesUnifiedActionInput) *ManageAlertRulesRequest
	GetBody() *ManageAlertRulesUnifiedActionInput
	SetCallSource(v string) *ManageAlertRulesRequest
	GetCallSource() *string
}

type ManageAlertRulesRequest struct {
	// The request body for managing alert rules. This body is shared by CREATE, UPDATE, PATCH, and BATCH_DELETE operations. Specify fields based on the action.
	Body       *ManageAlertRulesUnifiedActionInput `json:"body,omitempty" xml:"body,omitempty"`
	CallSource *string                             `json:"callSource,omitempty" xml:"callSource,omitempty"`
}

func (s ManageAlertRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ManageAlertRulesRequest) GoString() string {
	return s.String()
}

func (s *ManageAlertRulesRequest) GetBody() *ManageAlertRulesUnifiedActionInput {
	return s.Body
}

func (s *ManageAlertRulesRequest) GetCallSource() *string {
	return s.CallSource
}

func (s *ManageAlertRulesRequest) SetBody(v *ManageAlertRulesUnifiedActionInput) *ManageAlertRulesRequest {
	s.Body = v
	return s
}

func (s *ManageAlertRulesRequest) SetCallSource(v string) *ManageAlertRulesRequest {
	s.CallSource = &v
	return s
}

func (s *ManageAlertRulesRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
