// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRandomPasswordResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRandomPasswordResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRandomPasswordResponse
	GetStatusCode() *int32
	SetBody(v *GetRandomPasswordResponseBody) *GetRandomPasswordResponse
	GetBody() *GetRandomPasswordResponseBody
}

type GetRandomPasswordResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRandomPasswordResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRandomPasswordResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRandomPasswordResponse) GoString() string {
	return s.String()
}

func (s *GetRandomPasswordResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRandomPasswordResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRandomPasswordResponse) GetBody() *GetRandomPasswordResponseBody {
	return s.Body
}

func (s *GetRandomPasswordResponse) SetHeaders(v map[string]*string) *GetRandomPasswordResponse {
	s.Headers = v
	return s
}

func (s *GetRandomPasswordResponse) SetStatusCode(v int32) *GetRandomPasswordResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRandomPasswordResponse) SetBody(v *GetRandomPasswordResponseBody) *GetRandomPasswordResponse {
	s.Body = v
	return s
}

func (s *GetRandomPasswordResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
