// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHostCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHostCountResponse
	GetStatusCode() *int32
	SetBody(v *GetHostCountResponseBody) *GetHostCountResponse
	GetBody() *GetHostCountResponseBody
}

type GetHostCountResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHostCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHostCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHostCountResponse) GoString() string {
	return s.String()
}

func (s *GetHostCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHostCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHostCountResponse) GetBody() *GetHostCountResponseBody {
	return s.Body
}

func (s *GetHostCountResponse) SetHeaders(v map[string]*string) *GetHostCountResponse {
	s.Headers = v
	return s
}

func (s *GetHostCountResponse) SetStatusCode(v int32) *GetHostCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHostCountResponse) SetBody(v *GetHostCountResponseBody) *GetHostCountResponse {
	s.Body = v
	return s
}

func (s *GetHostCountResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
