// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCasterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCasterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCasterResponse
	GetStatusCode() *int32
	SetBody(v *CreateCasterResponseBody) *CreateCasterResponse
	GetBody() *CreateCasterResponseBody
}

type CreateCasterResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCasterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCasterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCasterResponse) GoString() string {
	return s.String()
}

func (s *CreateCasterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCasterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCasterResponse) GetBody() *CreateCasterResponseBody {
	return s.Body
}

func (s *CreateCasterResponse) SetHeaders(v map[string]*string) *CreateCasterResponse {
	s.Headers = v
	return s
}

func (s *CreateCasterResponse) SetStatusCode(v int32) *CreateCasterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCasterResponse) SetBody(v *CreateCasterResponseBody) *CreateCasterResponse {
	s.Body = v
	return s
}

func (s *CreateCasterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
