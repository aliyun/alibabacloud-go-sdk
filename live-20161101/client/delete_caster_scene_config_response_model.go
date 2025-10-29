// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCasterSceneConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCasterSceneConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCasterSceneConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCasterSceneConfigResponseBody) *DeleteCasterSceneConfigResponse
	GetBody() *DeleteCasterSceneConfigResponseBody
}

type DeleteCasterSceneConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCasterSceneConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCasterSceneConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCasterSceneConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteCasterSceneConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCasterSceneConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCasterSceneConfigResponse) GetBody() *DeleteCasterSceneConfigResponseBody {
	return s.Body
}

func (s *DeleteCasterSceneConfigResponse) SetHeaders(v map[string]*string) *DeleteCasterSceneConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteCasterSceneConfigResponse) SetStatusCode(v int32) *DeleteCasterSceneConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCasterSceneConfigResponse) SetBody(v *DeleteCasterSceneConfigResponseBody) *DeleteCasterSceneConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteCasterSceneConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
