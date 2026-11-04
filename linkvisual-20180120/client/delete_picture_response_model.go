// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePictureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePictureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePictureResponse
	GetStatusCode() *int32
	SetBody(v *DeletePictureResponseBody) *DeletePictureResponse
	GetBody() *DeletePictureResponseBody
}

type DeletePictureResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePictureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePictureResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePictureResponse) GoString() string {
	return s.String()
}

func (s *DeletePictureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePictureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePictureResponse) GetBody() *DeletePictureResponseBody {
	return s.Body
}

func (s *DeletePictureResponse) SetHeaders(v map[string]*string) *DeletePictureResponse {
	s.Headers = v
	return s
}

func (s *DeletePictureResponse) SetStatusCode(v int32) *DeletePictureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePictureResponse) SetBody(v *DeletePictureResponseBody) *DeletePictureResponse {
	s.Body = v
	return s
}

func (s *DeletePictureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
