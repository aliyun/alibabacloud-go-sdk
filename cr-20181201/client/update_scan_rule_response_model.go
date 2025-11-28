// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScanRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScanRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScanRuleResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScanRuleResponseBody) *UpdateScanRuleResponse
	GetBody() *UpdateScanRuleResponseBody
}

type UpdateScanRuleResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScanRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScanRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScanRuleResponse) GoString() string {
	return s.String()
}

func (s *UpdateScanRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScanRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScanRuleResponse) GetBody() *UpdateScanRuleResponseBody {
	return s.Body
}

func (s *UpdateScanRuleResponse) SetHeaders(v map[string]*string) *UpdateScanRuleResponse {
	s.Headers = v
	return s
}

func (s *UpdateScanRuleResponse) SetStatusCode(v int32) *UpdateScanRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScanRuleResponse) SetBody(v *UpdateScanRuleResponseBody) *UpdateScanRuleResponse {
	s.Body = v
	return s
}

func (s *UpdateScanRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
