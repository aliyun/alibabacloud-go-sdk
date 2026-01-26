// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntegrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteIntegrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteIntegrationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteIntegrationResponseBody) *DeleteIntegrationResponse
	GetBody() *DeleteIntegrationResponseBody
}

type DeleteIntegrationResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteIntegrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteIntegrationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntegrationResponse) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteIntegrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteIntegrationResponse) GetBody() *DeleteIntegrationResponseBody {
	return s.Body
}

func (s *DeleteIntegrationResponse) SetHeaders(v map[string]*string) *DeleteIntegrationResponse {
	s.Headers = v
	return s
}

func (s *DeleteIntegrationResponse) SetStatusCode(v int32) *DeleteIntegrationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteIntegrationResponse) SetBody(v *DeleteIntegrationResponseBody) *DeleteIntegrationResponse {
	s.Body = v
	return s
}

func (s *DeleteIntegrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
