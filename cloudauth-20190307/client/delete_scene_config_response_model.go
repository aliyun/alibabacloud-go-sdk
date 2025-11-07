// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSceneConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSceneConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSceneConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSceneConfigResponseBody) *DeleteSceneConfigResponse
	GetBody() *DeleteSceneConfigResponseBody
}

type DeleteSceneConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSceneConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSceneConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSceneConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteSceneConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSceneConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSceneConfigResponse) GetBody() *DeleteSceneConfigResponseBody {
	return s.Body
}

func (s *DeleteSceneConfigResponse) SetHeaders(v map[string]*string) *DeleteSceneConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteSceneConfigResponse) SetStatusCode(v int32) *DeleteSceneConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSceneConfigResponse) SetBody(v *DeleteSceneConfigResponseBody) *DeleteSceneConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteSceneConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
