// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChangeUserLangResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeUserLangResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeUserLangResponse
	GetStatusCode() *int32
	SetBody(v *ChangeUserLangResponseBody) *ChangeUserLangResponse
	GetBody() *ChangeUserLangResponseBody
}

type ChangeUserLangResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeUserLangResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeUserLangResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeUserLangResponse) GoString() string {
	return s.String()
}

func (s *ChangeUserLangResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeUserLangResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeUserLangResponse) GetBody() *ChangeUserLangResponseBody {
	return s.Body
}

func (s *ChangeUserLangResponse) SetHeaders(v map[string]*string) *ChangeUserLangResponse {
	s.Headers = v
	return s
}

func (s *ChangeUserLangResponse) SetStatusCode(v int32) *ChangeUserLangResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeUserLangResponse) SetBody(v *ChangeUserLangResponseBody) *ChangeUserLangResponse {
	s.Body = v
	return s
}

func (s *ChangeUserLangResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
