// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAvailableImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAvailableImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAvailableImagesResponse
	GetStatusCode() *int32
	SetBody(v *ListAvailableImagesResponseBody) *ListAvailableImagesResponse
	GetBody() *ListAvailableImagesResponseBody
}

type ListAvailableImagesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAvailableImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAvailableImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAvailableImagesResponse) GoString() string {
	return s.String()
}

func (s *ListAvailableImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAvailableImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAvailableImagesResponse) GetBody() *ListAvailableImagesResponseBody {
	return s.Body
}

func (s *ListAvailableImagesResponse) SetHeaders(v map[string]*string) *ListAvailableImagesResponse {
	s.Headers = v
	return s
}

func (s *ListAvailableImagesResponse) SetStatusCode(v int32) *ListAvailableImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAvailableImagesResponse) SetBody(v *ListAvailableImagesResponseBody) *ListAvailableImagesResponse {
	s.Body = v
	return s
}

func (s *ListAvailableImagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
