// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceUserResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaceUserResponseBody) *DeleteFaceUserResponse
	GetBody() *DeleteFaceUserResponseBody
}

type DeleteFaceUserResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceUserResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceUserResponse) GetBody() *DeleteFaceUserResponseBody {
	return s.Body
}

func (s *DeleteFaceUserResponse) SetHeaders(v map[string]*string) *DeleteFaceUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceUserResponse) SetStatusCode(v int32) *DeleteFaceUserResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceUserResponse) SetBody(v *DeleteFaceUserResponseBody) *DeleteFaceUserResponse {
	s.Body = v
	return s
}

func (s *DeleteFaceUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
