// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageAlertRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManageAlertRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManageAlertRulesResponse
	GetStatusCode() *int32
	SetBody(v *ServerResponseManageAlertRulesResult) *ManageAlertRulesResponse
	GetBody() *ServerResponseManageAlertRulesResult
}

type ManageAlertRulesResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ServerResponseManageAlertRulesResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageAlertRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ManageAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *ManageAlertRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManageAlertRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManageAlertRulesResponse) GetBody() *ServerResponseManageAlertRulesResult {
	return s.Body
}

func (s *ManageAlertRulesResponse) SetHeaders(v map[string]*string) *ManageAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *ManageAlertRulesResponse) SetStatusCode(v int32) *ManageAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageAlertRulesResponse) SetBody(v *ServerResponseManageAlertRulesResult) *ManageAlertRulesResponse {
	s.Body = v
	return s
}

func (s *ManageAlertRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
