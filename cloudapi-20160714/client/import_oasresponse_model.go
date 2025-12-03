// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportOASResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportOASResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportOASResponse
	GetStatusCode() *int32
	SetBody(v *ImportOASResponseBody) *ImportOASResponse
	GetBody() *ImportOASResponseBody
}

type ImportOASResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportOASResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportOASResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportOASResponse) GoString() string {
	return s.String()
}

func (s *ImportOASResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportOASResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportOASResponse) GetBody() *ImportOASResponseBody {
	return s.Body
}

func (s *ImportOASResponse) SetHeaders(v map[string]*string) *ImportOASResponse {
	s.Headers = v
	return s
}

func (s *ImportOASResponse) SetStatusCode(v int32) *ImportOASResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportOASResponse) SetBody(v *ImportOASResponseBody) *ImportOASResponse {
	s.Body = v
	return s
}

func (s *ImportOASResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
