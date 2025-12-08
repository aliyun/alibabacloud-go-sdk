// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaceResponseBody) *DeleteFaceResponse
	GetBody() *DeleteFaceResponseBody
}

type DeleteFaceResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceResponse) GetBody() *DeleteFaceResponseBody {
	return s.Body
}

func (s *DeleteFaceResponse) SetHeaders(v map[string]*string) *DeleteFaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceResponse) SetStatusCode(v int32) *DeleteFaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceResponse) SetBody(v *DeleteFaceResponseBody) *DeleteFaceResponse {
	s.Body = v
	return s
}

func (s *DeleteFaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
