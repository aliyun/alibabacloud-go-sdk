// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPtsSceneBaseLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPtsSceneBaseLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPtsSceneBaseLineResponse
	GetStatusCode() *int32
	SetBody(v *GetPtsSceneBaseLineResponseBody) *GetPtsSceneBaseLineResponse
	GetBody() *GetPtsSceneBaseLineResponseBody
}

type GetPtsSceneBaseLineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPtsSceneBaseLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPtsSceneBaseLineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPtsSceneBaseLineResponse) GoString() string {
	return s.String()
}

func (s *GetPtsSceneBaseLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPtsSceneBaseLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPtsSceneBaseLineResponse) GetBody() *GetPtsSceneBaseLineResponseBody {
	return s.Body
}

func (s *GetPtsSceneBaseLineResponse) SetHeaders(v map[string]*string) *GetPtsSceneBaseLineResponse {
	s.Headers = v
	return s
}

func (s *GetPtsSceneBaseLineResponse) SetStatusCode(v int32) *GetPtsSceneBaseLineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPtsSceneBaseLineResponse) SetBody(v *GetPtsSceneBaseLineResponseBody) *GetPtsSceneBaseLineResponse {
	s.Body = v
	return s
}

func (s *GetPtsSceneBaseLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
