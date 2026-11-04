// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceUserPictureResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFaceUserPictureResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFaceUserPictureResponse
	GetStatusCode() *int32
	SetBody(v *AddFaceUserPictureResponseBody) *AddFaceUserPictureResponse
	GetBody() *AddFaceUserPictureResponseBody
}

type AddFaceUserPictureResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceUserPictureResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceUserPictureResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFaceUserPictureResponse) GoString() string {
	return s.String()
}

func (s *AddFaceUserPictureResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFaceUserPictureResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFaceUserPictureResponse) GetBody() *AddFaceUserPictureResponseBody {
	return s.Body
}

func (s *AddFaceUserPictureResponse) SetHeaders(v map[string]*string) *AddFaceUserPictureResponse {
	s.Headers = v
	return s
}

func (s *AddFaceUserPictureResponse) SetStatusCode(v int32) *AddFaceUserPictureResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceUserPictureResponse) SetBody(v *AddFaceUserPictureResponseBody) *AddFaceUserPictureResponse {
	s.Body = v
	return s
}

func (s *AddFaceUserPictureResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
