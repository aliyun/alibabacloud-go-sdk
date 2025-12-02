// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImagesFromLibResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteImagesFromLibResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteImagesFromLibResponse
	GetStatusCode() *int32
	SetBody(v *DeleteImagesFromLibResponseBody) *DeleteImagesFromLibResponse
	GetBody() *DeleteImagesFromLibResponseBody
}

type DeleteImagesFromLibResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteImagesFromLibResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteImagesFromLibResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteImagesFromLibResponse) GoString() string {
	return s.String()
}

func (s *DeleteImagesFromLibResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteImagesFromLibResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteImagesFromLibResponse) GetBody() *DeleteImagesFromLibResponseBody {
	return s.Body
}

func (s *DeleteImagesFromLibResponse) SetHeaders(v map[string]*string) *DeleteImagesFromLibResponse {
	s.Headers = v
	return s
}

func (s *DeleteImagesFromLibResponse) SetStatusCode(v int32) *DeleteImagesFromLibResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteImagesFromLibResponse) SetBody(v *DeleteImagesFromLibResponseBody) *DeleteImagesFromLibResponse {
	s.Body = v
	return s
}

func (s *DeleteImagesFromLibResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
