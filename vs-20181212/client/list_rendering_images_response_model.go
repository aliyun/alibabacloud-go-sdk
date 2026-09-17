// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRenderingImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRenderingImagesResponse
	GetStatusCode() *int32
	SetBody(v *ListRenderingImagesResponseBody) *ListRenderingImagesResponse
	GetBody() *ListRenderingImagesResponseBody
}

type ListRenderingImagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRenderingImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRenderingImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingImagesResponse) GoString() string {
	return s.String()
}

func (s *ListRenderingImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRenderingImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRenderingImagesResponse) GetBody() *ListRenderingImagesResponseBody {
	return s.Body
}

func (s *ListRenderingImagesResponse) SetHeaders(v map[string]*string) *ListRenderingImagesResponse {
	s.Headers = v
	return s
}

func (s *ListRenderingImagesResponse) SetStatusCode(v int32) *ListRenderingImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRenderingImagesResponse) SetBody(v *ListRenderingImagesResponseBody) *ListRenderingImagesResponse {
	s.Body = v
	return s
}

func (s *ListRenderingImagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
