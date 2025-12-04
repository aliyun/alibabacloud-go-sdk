// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVpdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVpdResponse
	GetStatusCode() *int32
	SetBody(v *GetVpdResponseBody) *GetVpdResponse
	GetBody() *GetVpdResponseBody
}

type GetVpdResponse struct {
	Headers    map[string]*string  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVpdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVpdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVpdResponse) GoString() string {
	return s.String()
}

func (s *GetVpdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVpdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVpdResponse) GetBody() *GetVpdResponseBody {
	return s.Body
}

func (s *GetVpdResponse) SetHeaders(v map[string]*string) *GetVpdResponse {
	s.Headers = v
	return s
}

func (s *GetVpdResponse) SetStatusCode(v int32) *GetVpdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVpdResponse) SetBody(v *GetVpdResponseBody) *GetVpdResponse {
	s.Body = v
	return s
}

func (s *GetVpdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
