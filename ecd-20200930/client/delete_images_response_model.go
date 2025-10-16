// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteImagesResponse
	GetStatusCode() *int32
	SetBody(v *DeleteImagesResponseBody) *DeleteImagesResponse
	GetBody() *DeleteImagesResponseBody
}

type DeleteImagesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteImagesResponse) GoString() string {
	return s.String()
}

func (s *DeleteImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteImagesResponse) GetBody() *DeleteImagesResponseBody {
	return s.Body
}

func (s *DeleteImagesResponse) SetHeaders(v map[string]*string) *DeleteImagesResponse {
	s.Headers = v
	return s
}

func (s *DeleteImagesResponse) SetStatusCode(v int32) *DeleteImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImagesResponse) SetBody(v *DeleteImagesResponseBody) *DeleteImagesResponse {
	s.Body = v
	return s
}

func (s *DeleteImagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
