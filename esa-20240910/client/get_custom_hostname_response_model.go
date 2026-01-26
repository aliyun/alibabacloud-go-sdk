// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomHostnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomHostnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomHostnameResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomHostnameResponseBody) *GetCustomHostnameResponse
	GetBody() *GetCustomHostnameResponseBody
}

type GetCustomHostnameResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomHostnameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomHostnameResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomHostnameResponse) GoString() string {
	return s.String()
}

func (s *GetCustomHostnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomHostnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomHostnameResponse) GetBody() *GetCustomHostnameResponseBody {
	return s.Body
}

func (s *GetCustomHostnameResponse) SetHeaders(v map[string]*string) *GetCustomHostnameResponse {
	s.Headers = v
	return s
}

func (s *GetCustomHostnameResponse) SetStatusCode(v int32) *GetCustomHostnameResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomHostnameResponse) SetBody(v *GetCustomHostnameResponseBody) *GetCustomHostnameResponse {
	s.Body = v
	return s
}

func (s *GetCustomHostnameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
