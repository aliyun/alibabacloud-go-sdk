// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCustomImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCustomImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCustomImagesResponse
	GetStatusCode() *int32
	SetBody(v *ListCustomImagesResponseBody) *ListCustomImagesResponse
	GetBody() *ListCustomImagesResponseBody
}

type ListCustomImagesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCustomImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCustomImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCustomImagesResponse) GoString() string {
	return s.String()
}

func (s *ListCustomImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCustomImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCustomImagesResponse) GetBody() *ListCustomImagesResponseBody {
	return s.Body
}

func (s *ListCustomImagesResponse) SetHeaders(v map[string]*string) *ListCustomImagesResponse {
	s.Headers = v
	return s
}

func (s *ListCustomImagesResponse) SetStatusCode(v int32) *ListCustomImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCustomImagesResponse) SetBody(v *ListCustomImagesResponseBody) *ListCustomImagesResponse {
	s.Body = v
	return s
}

func (s *ListCustomImagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
