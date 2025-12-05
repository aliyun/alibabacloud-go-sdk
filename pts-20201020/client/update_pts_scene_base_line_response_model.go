// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePtsSceneBaseLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePtsSceneBaseLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePtsSceneBaseLineResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePtsSceneBaseLineResponseBody) *UpdatePtsSceneBaseLineResponse
	GetBody() *UpdatePtsSceneBaseLineResponseBody
}

type UpdatePtsSceneBaseLineResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePtsSceneBaseLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePtsSceneBaseLineResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePtsSceneBaseLineResponse) GoString() string {
	return s.String()
}

func (s *UpdatePtsSceneBaseLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePtsSceneBaseLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePtsSceneBaseLineResponse) GetBody() *UpdatePtsSceneBaseLineResponseBody {
	return s.Body
}

func (s *UpdatePtsSceneBaseLineResponse) SetHeaders(v map[string]*string) *UpdatePtsSceneBaseLineResponse {
	s.Headers = v
	return s
}

func (s *UpdatePtsSceneBaseLineResponse) SetStatusCode(v int32) *UpdatePtsSceneBaseLineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePtsSceneBaseLineResponse) SetBody(v *UpdatePtsSceneBaseLineResponseBody) *UpdatePtsSceneBaseLineResponse {
	s.Body = v
	return s
}

func (s *UpdatePtsSceneBaseLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
