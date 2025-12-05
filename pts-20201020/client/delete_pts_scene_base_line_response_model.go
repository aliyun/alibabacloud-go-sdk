// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePtsSceneBaseLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePtsSceneBaseLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePtsSceneBaseLineResponse
	GetStatusCode() *int32
	SetBody(v *DeletePtsSceneBaseLineResponseBody) *DeletePtsSceneBaseLineResponse
	GetBody() *DeletePtsSceneBaseLineResponseBody
}

type DeletePtsSceneBaseLineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePtsSceneBaseLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePtsSceneBaseLineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePtsSceneBaseLineResponse) GoString() string {
	return s.String()
}

func (s *DeletePtsSceneBaseLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePtsSceneBaseLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePtsSceneBaseLineResponse) GetBody() *DeletePtsSceneBaseLineResponseBody {
	return s.Body
}

func (s *DeletePtsSceneBaseLineResponse) SetHeaders(v map[string]*string) *DeletePtsSceneBaseLineResponse {
	s.Headers = v
	return s
}

func (s *DeletePtsSceneBaseLineResponse) SetStatusCode(v int32) *DeletePtsSceneBaseLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePtsSceneBaseLineResponse) SetBody(v *DeletePtsSceneBaseLineResponseBody) *DeletePtsSceneBaseLineResponse {
	s.Body = v
	return s
}

func (s *DeletePtsSceneBaseLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
