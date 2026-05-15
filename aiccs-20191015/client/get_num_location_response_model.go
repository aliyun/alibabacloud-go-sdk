// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNumLocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNumLocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNumLocationResponse
	GetStatusCode() *int32
	SetBody(v *GetNumLocationResponseBody) *GetNumLocationResponse
	GetBody() *GetNumLocationResponseBody
}

type GetNumLocationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNumLocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNumLocationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNumLocationResponse) GoString() string {
	return s.String()
}

func (s *GetNumLocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNumLocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNumLocationResponse) GetBody() *GetNumLocationResponseBody {
	return s.Body
}

func (s *GetNumLocationResponse) SetHeaders(v map[string]*string) *GetNumLocationResponse {
	s.Headers = v
	return s
}

func (s *GetNumLocationResponse) SetStatusCode(v int32) *GetNumLocationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNumLocationResponse) SetBody(v *GetNumLocationResponseBody) *GetNumLocationResponse {
	s.Body = v
	return s
}

func (s *GetNumLocationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
