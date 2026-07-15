// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportMediaResponse
	GetStatusCode() *int32
	SetBody(v *ImportMediaResponseBody) *ImportMediaResponse
	GetBody() *ImportMediaResponseBody
}

type ImportMediaResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportMediaResponse) GoString() string {
	return s.String()
}

func (s *ImportMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportMediaResponse) GetBody() *ImportMediaResponseBody {
	return s.Body
}

func (s *ImportMediaResponse) SetHeaders(v map[string]*string) *ImportMediaResponse {
	s.Headers = v
	return s
}

func (s *ImportMediaResponse) SetStatusCode(v int32) *ImportMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportMediaResponse) SetBody(v *ImportMediaResponseBody) *ImportMediaResponse {
	s.Body = v
	return s
}

func (s *ImportMediaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
