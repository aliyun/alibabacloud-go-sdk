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
	// The request body for managing alert rules. This body is shared by the CREATE, UPDATE, PATCH, and BATCH_DELETE actions. Specify the fields based on the action.
	Body *ManageAlertRulesUnifiedActionInput `json:"body,omitempty" xml:"body,omitempty"`
	// The identifier of the call source, which specifies the internal integration channel to which the caller belongs (such as bailian, integrationCenter, or managed_service_for_prometheus). This parameter is used to isolate traffic from different call sources. You do not need to specify this parameter for regular OpenAPI calls.
	//
	// example:
	//
	// bailian
	CallSource *string `json:"callSource,omitempty" xml:"callSource,omitempty"`
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
