// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportAppAlertRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportAppAlertRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportAppAlertRulesResponse
	GetStatusCode() *int32
	SetBody(v *ImportAppAlertRulesResponseBody) *ImportAppAlertRulesResponse
	GetBody() *ImportAppAlertRulesResponseBody
}

type ImportAppAlertRulesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportAppAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportAppAlertRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportAppAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *ImportAppAlertRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportAppAlertRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportAppAlertRulesResponse) GetBody() *ImportAppAlertRulesResponseBody {
	return s.Body
}

func (s *ImportAppAlertRulesResponse) SetHeaders(v map[string]*string) *ImportAppAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *ImportAppAlertRulesResponse) SetStatusCode(v int32) *ImportAppAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportAppAlertRulesResponse) SetBody(v *ImportAppAlertRulesResponseBody) *ImportAppAlertRulesResponse {
	s.Body = v
	return s
}

func (s *ImportAppAlertRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
