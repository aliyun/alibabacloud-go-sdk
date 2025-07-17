// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListIntegrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListIntegrationResponse
	GetStatusCode() *int32
	SetBody(v *ListIntegrationResponseBody) *ListIntegrationResponse
	GetBody() *ListIntegrationResponseBody
}

type ListIntegrationResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIntegrationResponse) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationResponse) GoString() string {
	return s.String()
}

func (s *ListIntegrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListIntegrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListIntegrationResponse) GetBody() *ListIntegrationResponseBody {
	return s.Body
}

func (s *ListIntegrationResponse) SetHeaders(v map[string]*string) *ListIntegrationResponse {
	s.Headers = v
	return s
}

func (s *ListIntegrationResponse) SetStatusCode(v int32) *ListIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIntegrationResponse) SetBody(v *ListIntegrationResponseBody) *ListIntegrationResponse {
	s.Body = v
	return s
}

func (s *ListIntegrationResponse) Validate() error {
	return dara.Validate(s)
}
