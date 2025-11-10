// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelHandshakeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelHandshakeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelHandshakeResponse
	GetStatusCode() *int32
	SetBody(v *CancelHandshakeResponseBody) *CancelHandshakeResponse
	GetBody() *CancelHandshakeResponseBody
}

type CancelHandshakeResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelHandshakeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelHandshakeResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelHandshakeResponse) GoString() string {
	return s.String()
}

func (s *CancelHandshakeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelHandshakeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelHandshakeResponse) GetBody() *CancelHandshakeResponseBody {
	return s.Body
}

func (s *CancelHandshakeResponse) SetHeaders(v map[string]*string) *CancelHandshakeResponse {
	s.Headers = v
	return s
}

func (s *CancelHandshakeResponse) SetStatusCode(v int32) *CancelHandshakeResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelHandshakeResponse) SetBody(v *CancelHandshakeResponseBody) *CancelHandshakeResponse {
	s.Body = v
	return s
}

func (s *CancelHandshakeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
