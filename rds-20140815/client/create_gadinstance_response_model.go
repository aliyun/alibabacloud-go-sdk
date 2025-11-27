// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateGADInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateGADInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateGADInstanceResponse
	GetStatusCode() *int32
	SetBody(v *CreateGADInstanceResponseBody) *CreateGADInstanceResponse
	GetBody() *CreateGADInstanceResponseBody
}

type CreateGADInstanceResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateGADInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGADInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateGADInstanceResponse) GoString() string {
	return s.String()
}

func (s *CreateGADInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateGADInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateGADInstanceResponse) GetBody() *CreateGADInstanceResponseBody {
	return s.Body
}

func (s *CreateGADInstanceResponse) SetHeaders(v map[string]*string) *CreateGADInstanceResponse {
	s.Headers = v
	return s
}

func (s *CreateGADInstanceResponse) SetStatusCode(v int32) *CreateGADInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGADInstanceResponse) SetBody(v *CreateGADInstanceResponseBody) *CreateGADInstanceResponse {
	s.Body = v
	return s
}

func (s *CreateGADInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
