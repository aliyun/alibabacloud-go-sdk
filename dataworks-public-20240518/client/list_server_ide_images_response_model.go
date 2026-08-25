// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServerIdeImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServerIdeImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServerIdeImagesResponse
	GetStatusCode() *int32
	SetBody(v *ListServerIdeImagesResponseBody) *ListServerIdeImagesResponse
	GetBody() *ListServerIdeImagesResponseBody
}

type ListServerIdeImagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServerIdeImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServerIdeImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServerIdeImagesResponse) GoString() string {
	return s.String()
}

func (s *ListServerIdeImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServerIdeImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServerIdeImagesResponse) GetBody() *ListServerIdeImagesResponseBody {
	return s.Body
}

func (s *ListServerIdeImagesResponse) SetHeaders(v map[string]*string) *ListServerIdeImagesResponse {
	s.Headers = v
	return s
}

func (s *ListServerIdeImagesResponse) SetStatusCode(v int32) *ListServerIdeImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServerIdeImagesResponse) SetBody(v *ListServerIdeImagesResponseBody) *ListServerIdeImagesResponse {
	s.Body = v
	return s
}

func (s *ListServerIdeImagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
