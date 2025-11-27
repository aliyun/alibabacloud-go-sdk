// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddExtCodeSignResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddExtCodeSignResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddExtCodeSignResponse
	GetStatusCode() *int32
	SetBody(v *AddExtCodeSignResponseBody) *AddExtCodeSignResponse
	GetBody() *AddExtCodeSignResponseBody
}

type AddExtCodeSignResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddExtCodeSignResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddExtCodeSignResponse) String() string {
	return dara.Prettify(s)
}

func (s AddExtCodeSignResponse) GoString() string {
	return s.String()
}

func (s *AddExtCodeSignResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddExtCodeSignResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddExtCodeSignResponse) GetBody() *AddExtCodeSignResponseBody {
	return s.Body
}

func (s *AddExtCodeSignResponse) SetHeaders(v map[string]*string) *AddExtCodeSignResponse {
	s.Headers = v
	return s
}

func (s *AddExtCodeSignResponse) SetStatusCode(v int32) *AddExtCodeSignResponse {
	s.StatusCode = &v
	return s
}

func (s *AddExtCodeSignResponse) SetBody(v *AddExtCodeSignResponseBody) *AddExtCodeSignResponse {
	s.Body = v
	return s
}

func (s *AddExtCodeSignResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
