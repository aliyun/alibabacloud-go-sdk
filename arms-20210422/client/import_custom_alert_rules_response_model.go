// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCustomAlertRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportCustomAlertRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportCustomAlertRulesResponse
	GetStatusCode() *int32
	SetBody(v *ImportCustomAlertRulesResponseBody) *ImportCustomAlertRulesResponse
	GetBody() *ImportCustomAlertRulesResponseBody
}

type ImportCustomAlertRulesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportCustomAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportCustomAlertRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportCustomAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *ImportCustomAlertRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportCustomAlertRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportCustomAlertRulesResponse) GetBody() *ImportCustomAlertRulesResponseBody {
	return s.Body
}

func (s *ImportCustomAlertRulesResponse) SetHeaders(v map[string]*string) *ImportCustomAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *ImportCustomAlertRulesResponse) SetStatusCode(v int32) *ImportCustomAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportCustomAlertRulesResponse) SetBody(v *ImportCustomAlertRulesResponseBody) *ImportCustomAlertRulesResponse {
	s.Body = v
	return s
}

func (s *ImportCustomAlertRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
