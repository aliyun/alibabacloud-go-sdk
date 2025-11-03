// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExchangeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExchangeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExchangeResponse
	GetStatusCode() *int32
	SetBody(v *CreateExchangeResponseBody) *CreateExchangeResponse
	GetBody() *CreateExchangeResponseBody
}

type CreateExchangeResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExchangeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExchangeResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExchangeResponse) GoString() string {
	return s.String()
}

func (s *CreateExchangeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExchangeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExchangeResponse) GetBody() *CreateExchangeResponseBody {
	return s.Body
}

func (s *CreateExchangeResponse) SetHeaders(v map[string]*string) *CreateExchangeResponse {
	s.Headers = v
	return s
}

func (s *CreateExchangeResponse) SetStatusCode(v int32) *CreateExchangeResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExchangeResponse) SetBody(v *CreateExchangeResponseBody) *CreateExchangeResponse {
	s.Body = v
	return s
}

func (s *CreateExchangeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
