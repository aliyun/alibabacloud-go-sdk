// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserLangResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserLangResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserLangResponse
	GetStatusCode() *int32
	SetBody(v *GetUserLangResponseBody) *GetUserLangResponse
	GetBody() *GetUserLangResponseBody
}

type GetUserLangResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserLangResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserLangResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserLangResponse) GoString() string {
	return s.String()
}

func (s *GetUserLangResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserLangResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserLangResponse) GetBody() *GetUserLangResponseBody {
	return s.Body
}

func (s *GetUserLangResponse) SetHeaders(v map[string]*string) *GetUserLangResponse {
	s.Headers = v
	return s
}

func (s *GetUserLangResponse) SetStatusCode(v int32) *GetUserLangResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserLangResponse) SetBody(v *GetUserLangResponseBody) *GetUserLangResponse {
	s.Body = v
	return s
}

func (s *GetUserLangResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
