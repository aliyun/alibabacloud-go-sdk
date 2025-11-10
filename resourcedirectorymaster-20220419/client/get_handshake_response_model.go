// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHandshakeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHandshakeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHandshakeResponse
	GetStatusCode() *int32
	SetBody(v *GetHandshakeResponseBody) *GetHandshakeResponse
	GetBody() *GetHandshakeResponseBody
}

type GetHandshakeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHandshakeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHandshakeResponse) GoString() string {
	return s.String()
}

func (s *GetHandshakeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHandshakeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHandshakeResponse) GetBody() *GetHandshakeResponseBody {
	return s.Body
}

func (s *GetHandshakeResponse) SetHeaders(v map[string]*string) *GetHandshakeResponse {
	s.Headers = v
	return s
}

func (s *GetHandshakeResponse) SetStatusCode(v int32) *GetHandshakeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHandshakeResponse) SetBody(v *GetHandshakeResponseBody) *GetHandshakeResponse {
	s.Body = v
	return s
}

func (s *GetHandshakeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
