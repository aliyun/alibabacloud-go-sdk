// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadDetailsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUploadDetailsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUploadDetailsResponse
	GetStatusCode() *int32
	SetBody(v *GetUploadDetailsResponseBody) *GetUploadDetailsResponse
	GetBody() *GetUploadDetailsResponseBody
}

type GetUploadDetailsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadDetailsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadDetailsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUploadDetailsResponse) GoString() string {
	return s.String()
}

func (s *GetUploadDetailsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUploadDetailsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUploadDetailsResponse) GetBody() *GetUploadDetailsResponseBody {
	return s.Body
}

func (s *GetUploadDetailsResponse) SetHeaders(v map[string]*string) *GetUploadDetailsResponse {
	s.Headers = v
	return s
}

func (s *GetUploadDetailsResponse) SetStatusCode(v int32) *GetUploadDetailsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadDetailsResponse) SetBody(v *GetUploadDetailsResponseBody) *GetUploadDetailsResponse {
	s.Body = v
	return s
}

func (s *GetUploadDetailsResponse) Validate() error {
	return dara.Validate(s)
}
