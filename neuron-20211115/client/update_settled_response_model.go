// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSettledResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSettledResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSettledResponse
	GetStatusCode() *int32
	SetBody(v *Settled) *UpdateSettledResponse
	GetBody() *Settled
}

type UpdateSettledResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Settled           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSettledResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSettledResponse) GoString() string {
	return s.String()
}

func (s *UpdateSettledResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSettledResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSettledResponse) GetBody() *Settled {
	return s.Body
}

func (s *UpdateSettledResponse) SetHeaders(v map[string]*string) *UpdateSettledResponse {
	s.Headers = v
	return s
}

func (s *UpdateSettledResponse) SetStatusCode(v int32) *UpdateSettledResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSettledResponse) SetBody(v *Settled) *UpdateSettledResponse {
	s.Body = v
	return s
}

func (s *UpdateSettledResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
