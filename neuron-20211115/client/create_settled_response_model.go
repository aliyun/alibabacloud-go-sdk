// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSettledResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSettledResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSettledResponse
	GetStatusCode() *int32
	SetBody(v *Settled) *CreateSettledResponse
	GetBody() *Settled
}

type CreateSettledResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Settled           `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSettledResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSettledResponse) GoString() string {
	return s.String()
}

func (s *CreateSettledResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSettledResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSettledResponse) GetBody() *Settled {
	return s.Body
}

func (s *CreateSettledResponse) SetHeaders(v map[string]*string) *CreateSettledResponse {
	s.Headers = v
	return s
}

func (s *CreateSettledResponse) SetStatusCode(v int32) *CreateSettledResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSettledResponse) SetBody(v *Settled) *CreateSettledResponse {
	s.Body = v
	return s
}

func (s *CreateSettledResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
