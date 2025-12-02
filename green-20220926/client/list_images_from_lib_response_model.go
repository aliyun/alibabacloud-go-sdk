// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImagesFromLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImagesFromLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImagesFromLibResponse
	GetStatusCode() *int32
	SetBody(v *ListImagesFromLibResponseBody) *ListImagesFromLibResponse
	GetBody() *ListImagesFromLibResponseBody
}

type ListImagesFromLibResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImagesFromLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImagesFromLibResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImagesFromLibResponse) GoString() string {
	return s.String()
}

func (s *ListImagesFromLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImagesFromLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImagesFromLibResponse) GetBody() *ListImagesFromLibResponseBody {
	return s.Body
}

func (s *ListImagesFromLibResponse) SetHeaders(v map[string]*string) *ListImagesFromLibResponse {
	s.Headers = v
	return s
}

func (s *ListImagesFromLibResponse) SetStatusCode(v int32) *ListImagesFromLibResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesFromLibResponse) SetBody(v *ListImagesFromLibResponseBody) *ListImagesFromLibResponse {
	s.Body = v
	return s
}

func (s *ListImagesFromLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
