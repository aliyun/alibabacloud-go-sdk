// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConnectorClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetConnectorClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetConnectorClientResponse
	GetStatusCode() *int32
	SetBody(v *GetConnectorClientResponseBody) *GetConnectorClientResponse
	GetBody() *GetConnectorClientResponseBody
}

type GetConnectorClientResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetConnectorClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetConnectorClientResponse) String() string {
	return dara.Prettify(s)
}

func (s GetConnectorClientResponse) GoString() string {
	return s.String()
}

func (s *GetConnectorClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetConnectorClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetConnectorClientResponse) GetBody() *GetConnectorClientResponseBody {
	return s.Body
}

func (s *GetConnectorClientResponse) SetHeaders(v map[string]*string) *GetConnectorClientResponse {
	s.Headers = v
	return s
}

func (s *GetConnectorClientResponse) SetStatusCode(v int32) *GetConnectorClientResponse {
	s.StatusCode = &v
	return s
}

func (s *GetConnectorClientResponse) SetBody(v *GetConnectorClientResponseBody) *GetConnectorClientResponse {
	s.Body = v
	return s
}

func (s *GetConnectorClientResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
