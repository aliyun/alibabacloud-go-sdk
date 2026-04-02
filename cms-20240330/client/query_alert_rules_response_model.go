// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAlertRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAlertRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAlertRulesResponse
	GetStatusCode() *int32
	SetBody(v *ServerResponseQueryAlertRulesResult) *QueryAlertRulesResponse
	GetBody() *ServerResponseQueryAlertRulesResult
}

type QueryAlertRulesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ServerResponseQueryAlertRulesResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAlertRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *QueryAlertRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAlertRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAlertRulesResponse) GetBody() *ServerResponseQueryAlertRulesResult {
	return s.Body
}

func (s *QueryAlertRulesResponse) SetHeaders(v map[string]*string) *QueryAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *QueryAlertRulesResponse) SetStatusCode(v int32) *QueryAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAlertRulesResponse) SetBody(v *ServerResponseQueryAlertRulesResult) *QueryAlertRulesResponse {
	s.Body = v
	return s
}

func (s *QueryAlertRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
