// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBaselineResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBaselineResponseBody) *UpdateBaselineResponse
	GetBody() *UpdateBaselineResponseBody
}

type UpdateBaselineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBaselineResponse) GoString() string {
	return s.String()
}

func (s *UpdateBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBaselineResponse) GetBody() *UpdateBaselineResponseBody {
	return s.Body
}

func (s *UpdateBaselineResponse) SetHeaders(v map[string]*string) *UpdateBaselineResponse {
	s.Headers = v
	return s
}

func (s *UpdateBaselineResponse) SetStatusCode(v int32) *UpdateBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBaselineResponse) SetBody(v *UpdateBaselineResponseBody) *UpdateBaselineResponse {
	s.Body = v
	return s
}

func (s *UpdateBaselineResponse) Validate() error {
	return dara.Validate(s)
}
