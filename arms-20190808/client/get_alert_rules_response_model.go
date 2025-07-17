// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlertRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAlertRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAlertRulesResponse
	GetStatusCode() *int32
	SetBody(v *GetAlertRulesResponseBody) *GetAlertRulesResponse
	GetBody() *GetAlertRulesResponseBody
}

type GetAlertRulesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAlertRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *GetAlertRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAlertRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAlertRulesResponse) GetBody() *GetAlertRulesResponseBody {
	return s.Body
}

func (s *GetAlertRulesResponse) SetHeaders(v map[string]*string) *GetAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *GetAlertRulesResponse) SetStatusCode(v int32) *GetAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAlertRulesResponse) SetBody(v *GetAlertRulesResponseBody) *GetAlertRulesResponse {
	s.Body = v
	return s
}

func (s *GetAlertRulesResponse) Validate() error {
	return dara.Validate(s)
}
