// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomImagesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomImagesResponseBody) *DeleteCustomImagesResponse
	GetBody() *DeleteCustomImagesResponseBody
}

type DeleteCustomImagesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomImagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomImagesResponse) GetBody() *DeleteCustomImagesResponseBody {
	return s.Body
}

func (s *DeleteCustomImagesResponse) SetHeaders(v map[string]*string) *DeleteCustomImagesResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomImagesResponse) SetStatusCode(v int32) *DeleteCustomImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomImagesResponse) SetBody(v *DeleteCustomImagesResponseBody) *DeleteCustomImagesResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomImagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
