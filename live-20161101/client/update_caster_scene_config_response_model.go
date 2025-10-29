// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCasterSceneConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCasterSceneConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCasterSceneConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCasterSceneConfigResponseBody) *UpdateCasterSceneConfigResponse
	GetBody() *UpdateCasterSceneConfigResponseBody
}

type UpdateCasterSceneConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCasterSceneConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCasterSceneConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCasterSceneConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateCasterSceneConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCasterSceneConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCasterSceneConfigResponse) GetBody() *UpdateCasterSceneConfigResponseBody {
	return s.Body
}

func (s *UpdateCasterSceneConfigResponse) SetHeaders(v map[string]*string) *UpdateCasterSceneConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateCasterSceneConfigResponse) SetStatusCode(v int32) *UpdateCasterSceneConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCasterSceneConfigResponse) SetBody(v *UpdateCasterSceneConfigResponseBody) *UpdateCasterSceneConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateCasterSceneConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
