// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConnectorClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteConnectorClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteConnectorClientResponse
	GetStatusCode() *int32
	SetBody(v *DeleteConnectorClientResponseBody) *DeleteConnectorClientResponse
	GetBody() *DeleteConnectorClientResponseBody
}

type DeleteConnectorClientResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteConnectorClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteConnectorClientResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteConnectorClientResponse) GoString() string {
	return s.String()
}

func (s *DeleteConnectorClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteConnectorClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteConnectorClientResponse) GetBody() *DeleteConnectorClientResponseBody {
	return s.Body
}

func (s *DeleteConnectorClientResponse) SetHeaders(v map[string]*string) *DeleteConnectorClientResponse {
	s.Headers = v
	return s
}

func (s *DeleteConnectorClientResponse) SetStatusCode(v int32) *DeleteConnectorClientResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteConnectorClientResponse) SetBody(v *DeleteConnectorClientResponseBody) *DeleteConnectorClientResponse {
	s.Body = v
	return s
}

func (s *DeleteConnectorClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
