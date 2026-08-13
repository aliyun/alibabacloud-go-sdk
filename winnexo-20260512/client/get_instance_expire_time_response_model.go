// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceExpireTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetInstanceExpireTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetInstanceExpireTimeResponse
	GetStatusCode() *int32
	SetBody(v *GetInstanceExpireTimeResponseBody) *GetInstanceExpireTimeResponse
	GetBody() *GetInstanceExpireTimeResponseBody
}

type GetInstanceExpireTimeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetInstanceExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetInstanceExpireTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceExpireTimeResponse) GoString() string {
	return s.String()
}

func (s *GetInstanceExpireTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetInstanceExpireTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetInstanceExpireTimeResponse) GetBody() *GetInstanceExpireTimeResponseBody {
	return s.Body
}

func (s *GetInstanceExpireTimeResponse) SetHeaders(v map[string]*string) *GetInstanceExpireTimeResponse {
	s.Headers = v
	return s
}

func (s *GetInstanceExpireTimeResponse) SetStatusCode(v int32) *GetInstanceExpireTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetInstanceExpireTimeResponse) SetBody(v *GetInstanceExpireTimeResponseBody) *GetInstanceExpireTimeResponse {
	s.Body = v
	return s
}

func (s *GetInstanceExpireTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
