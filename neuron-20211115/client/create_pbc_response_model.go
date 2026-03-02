// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePbcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePbcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePbcResponse
	GetStatusCode() *int32
	SetBody(v *Pbc) *CreatePbcResponse
	GetBody() *Pbc
}

type CreatePbcResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Pbc               `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePbcResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePbcResponse) GoString() string {
	return s.String()
}

func (s *CreatePbcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePbcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePbcResponse) GetBody() *Pbc {
	return s.Body
}

func (s *CreatePbcResponse) SetHeaders(v map[string]*string) *CreatePbcResponse {
	s.Headers = v
	return s
}

func (s *CreatePbcResponse) SetStatusCode(v int32) *CreatePbcResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePbcResponse) SetBody(v *Pbc) *CreatePbcResponse {
	s.Body = v
	return s
}

func (s *CreatePbcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
