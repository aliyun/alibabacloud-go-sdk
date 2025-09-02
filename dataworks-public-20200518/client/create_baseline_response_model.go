// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBaselineResponse
	GetStatusCode() *int32
	SetBody(v *CreateBaselineResponseBody) *CreateBaselineResponse
	GetBody() *CreateBaselineResponseBody
}

type CreateBaselineResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBaselineResponse) GoString() string {
	return s.String()
}

func (s *CreateBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBaselineResponse) GetBody() *CreateBaselineResponseBody {
	return s.Body
}

func (s *CreateBaselineResponse) SetHeaders(v map[string]*string) *CreateBaselineResponse {
	s.Headers = v
	return s
}

func (s *CreateBaselineResponse) SetStatusCode(v int32) *CreateBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBaselineResponse) SetBody(v *CreateBaselineResponseBody) *CreateBaselineResponse {
	s.Body = v
	return s
}

func (s *CreateBaselineResponse) Validate() error {
	return dara.Validate(s)
}
