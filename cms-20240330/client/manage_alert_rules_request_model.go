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
}

type ManageAlertRulesRequest struct {
	Body *ManageAlertRulesUnifiedActionInput `json:"body,omitempty" xml:"body,omitempty"`
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

func (s *ManageAlertRulesRequest) SetBody(v *ManageAlertRulesUnifiedActionInput) *ManageAlertRulesRequest {
	s.Body = v
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
