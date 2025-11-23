// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSLARulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSLARulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSLARulesResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSLARulesResponseBody) *UpdateSLARulesResponse
	GetBody() *UpdateSLARulesResponseBody
}

type UpdateSLARulesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSLARulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSLARulesResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSLARulesResponse) GoString() string {
	return s.String()
}

func (s *UpdateSLARulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSLARulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSLARulesResponse) GetBody() *UpdateSLARulesResponseBody {
	return s.Body
}

func (s *UpdateSLARulesResponse) SetHeaders(v map[string]*string) *UpdateSLARulesResponse {
	s.Headers = v
	return s
}

func (s *UpdateSLARulesResponse) SetStatusCode(v int32) *UpdateSLARulesResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSLARulesResponse) SetBody(v *UpdateSLARulesResponseBody) *UpdateSLARulesResponse {
	s.Body = v
	return s
}

func (s *UpdateSLARulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
