// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPersonalICodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPersonalICodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPersonalICodeResponse
	GetStatusCode() *int32
	SetBody(v *GetPersonalICodeResponseBody) *GetPersonalICodeResponse
	GetBody() *GetPersonalICodeResponseBody
}

type GetPersonalICodeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPersonalICodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPersonalICodeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPersonalICodeResponse) GoString() string {
	return s.String()
}

func (s *GetPersonalICodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPersonalICodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPersonalICodeResponse) GetBody() *GetPersonalICodeResponseBody {
	return s.Body
}

func (s *GetPersonalICodeResponse) SetHeaders(v map[string]*string) *GetPersonalICodeResponse {
	s.Headers = v
	return s
}

func (s *GetPersonalICodeResponse) SetStatusCode(v int32) *GetPersonalICodeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPersonalICodeResponse) SetBody(v *GetPersonalICodeResponseBody) *GetPersonalICodeResponse {
	s.Body = v
	return s
}

func (s *GetPersonalICodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
