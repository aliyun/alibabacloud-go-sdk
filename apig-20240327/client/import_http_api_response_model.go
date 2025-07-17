// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportHttpApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportHttpApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportHttpApiResponse
	GetStatusCode() *int32
	SetBody(v *ImportHttpApiResponseBody) *ImportHttpApiResponse
	GetBody() *ImportHttpApiResponseBody
}

type ImportHttpApiResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportHttpApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportHttpApiResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportHttpApiResponse) GoString() string {
	return s.String()
}

func (s *ImportHttpApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportHttpApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportHttpApiResponse) GetBody() *ImportHttpApiResponseBody {
	return s.Body
}

func (s *ImportHttpApiResponse) SetHeaders(v map[string]*string) *ImportHttpApiResponse {
	s.Headers = v
	return s
}

func (s *ImportHttpApiResponse) SetStatusCode(v int32) *ImportHttpApiResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportHttpApiResponse) SetBody(v *ImportHttpApiResponseBody) *ImportHttpApiResponse {
	s.Body = v
	return s
}

func (s *ImportHttpApiResponse) Validate() error {
	return dara.Validate(s)
}
