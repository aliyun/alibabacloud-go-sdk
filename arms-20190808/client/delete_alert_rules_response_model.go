// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAlertRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAlertRulesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAlertRulesResponseBody) *DeleteAlertRulesResponse
	GetBody() *DeleteAlertRulesResponseBody
}

type DeleteAlertRulesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAlertRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAlertRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertRulesResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlertRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAlertRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAlertRulesResponse) GetBody() *DeleteAlertRulesResponseBody {
	return s.Body
}

func (s *DeleteAlertRulesResponse) SetHeaders(v map[string]*string) *DeleteAlertRulesResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlertRulesResponse) SetStatusCode(v int32) *DeleteAlertRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAlertRulesResponse) SetBody(v *DeleteAlertRulesResponseBody) *DeleteAlertRulesResponse {
	s.Body = v
	return s
}

func (s *DeleteAlertRulesResponse) Validate() error {
	return dara.Validate(s)
}
