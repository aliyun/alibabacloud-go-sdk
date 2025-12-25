// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScenePublishResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ScenePublishResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ScenePublishResponse
	GetStatusCode() *int32
	SetBody(v *ScenePublishResponseBody) *ScenePublishResponse
	GetBody() *ScenePublishResponseBody
}

type ScenePublishResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScenePublishResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScenePublishResponse) String() string {
	return dara.Prettify(s)
}

func (s ScenePublishResponse) GoString() string {
	return s.String()
}

func (s *ScenePublishResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ScenePublishResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ScenePublishResponse) GetBody() *ScenePublishResponseBody {
	return s.Body
}

func (s *ScenePublishResponse) SetHeaders(v map[string]*string) *ScenePublishResponse {
	s.Headers = v
	return s
}

func (s *ScenePublishResponse) SetStatusCode(v int32) *ScenePublishResponse {
	s.StatusCode = &v
	return s
}

func (s *ScenePublishResponse) SetBody(v *ScenePublishResponseBody) *ScenePublishResponse {
	s.Body = v
	return s
}

func (s *ScenePublishResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
