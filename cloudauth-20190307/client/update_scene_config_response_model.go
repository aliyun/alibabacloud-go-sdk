// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSceneConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSceneConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSceneConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSceneConfigResponseBody) *UpdateSceneConfigResponse
	GetBody() *UpdateSceneConfigResponseBody
}

type UpdateSceneConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSceneConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSceneConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSceneConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateSceneConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSceneConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSceneConfigResponse) GetBody() *UpdateSceneConfigResponseBody {
	return s.Body
}

func (s *UpdateSceneConfigResponse) SetHeaders(v map[string]*string) *UpdateSceneConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateSceneConfigResponse) SetStatusCode(v int32) *UpdateSceneConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSceneConfigResponse) SetBody(v *UpdateSceneConfigResponseBody) *UpdateSceneConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateSceneConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
