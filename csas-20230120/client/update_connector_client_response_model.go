// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectorClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateConnectorClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateConnectorClientResponse
	GetStatusCode() *int32
	SetBody(v *UpdateConnectorClientResponseBody) *UpdateConnectorClientResponse
	GetBody() *UpdateConnectorClientResponseBody
}

type UpdateConnectorClientResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateConnectorClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateConnectorClientResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectorClientResponse) GoString() string {
	return s.String()
}

func (s *UpdateConnectorClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateConnectorClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateConnectorClientResponse) GetBody() *UpdateConnectorClientResponseBody {
	return s.Body
}

func (s *UpdateConnectorClientResponse) SetHeaders(v map[string]*string) *UpdateConnectorClientResponse {
	s.Headers = v
	return s
}

func (s *UpdateConnectorClientResponse) SetStatusCode(v int32) *UpdateConnectorClientResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateConnectorClientResponse) SetBody(v *UpdateConnectorClientResponseBody) *UpdateConnectorClientResponse {
	s.Body = v
	return s
}

func (s *UpdateConnectorClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
