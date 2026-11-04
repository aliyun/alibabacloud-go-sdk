// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceUserPictureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceUserPictureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceUserPictureResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaceUserPictureResponseBody) *DeleteFaceUserPictureResponse
	GetBody() *DeleteFaceUserPictureResponseBody
}

type DeleteFaceUserPictureResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceUserPictureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceUserPictureResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceUserPictureResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserPictureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceUserPictureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceUserPictureResponse) GetBody() *DeleteFaceUserPictureResponseBody {
	return s.Body
}

func (s *DeleteFaceUserPictureResponse) SetHeaders(v map[string]*string) *DeleteFaceUserPictureResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceUserPictureResponse) SetStatusCode(v int32) *DeleteFaceUserPictureResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceUserPictureResponse) SetBody(v *DeleteFaceUserPictureResponseBody) *DeleteFaceUserPictureResponse {
	s.Body = v
	return s
}

func (s *DeleteFaceUserPictureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
