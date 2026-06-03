// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSmsSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSmsSignResponse
	GetStatusCode() *int32
	SetBody(v *GetSmsSignResponseBody) *GetSmsSignResponse
	GetBody() *GetSmsSignResponseBody
}

type GetSmsSignResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmsSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmsSignResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSmsSignResponse) GoString() string {
	return s.String()
}

func (s *GetSmsSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSmsSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSmsSignResponse) GetBody() *GetSmsSignResponseBody {
	return s.Body
}

func (s *GetSmsSignResponse) SetHeaders(v map[string]*string) *GetSmsSignResponse {
	s.Headers = v
	return s
}

func (s *GetSmsSignResponse) SetStatusCode(v int32) *GetSmsSignResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmsSignResponse) SetBody(v *GetSmsSignResponseBody) *GetSmsSignResponse {
	s.Body = v
	return s
}

func (s *GetSmsSignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
