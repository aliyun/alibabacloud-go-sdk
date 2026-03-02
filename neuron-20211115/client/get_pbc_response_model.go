// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPbcResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPbcResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPbcResponse
	GetStatusCode() *int32
	SetBody(v *Pbc) *GetPbcResponse
	GetBody() *Pbc
}

type GetPbcResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Pbc               `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPbcResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPbcResponse) GoString() string {
	return s.String()
}

func (s *GetPbcResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPbcResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPbcResponse) GetBody() *Pbc {
	return s.Body
}

func (s *GetPbcResponse) SetHeaders(v map[string]*string) *GetPbcResponse {
	s.Headers = v
	return s
}

func (s *GetPbcResponse) SetStatusCode(v int32) *GetPbcResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPbcResponse) SetBody(v *Pbc) *GetPbcResponse {
	s.Body = v
	return s
}

func (s *GetPbcResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
