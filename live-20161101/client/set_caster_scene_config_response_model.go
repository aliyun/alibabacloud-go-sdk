// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetCasterSceneConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetCasterSceneConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetCasterSceneConfigResponse
	GetStatusCode() *int32
	SetBody(v *SetCasterSceneConfigResponseBody) *SetCasterSceneConfigResponse
	GetBody() *SetCasterSceneConfigResponseBody
}

type SetCasterSceneConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetCasterSceneConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetCasterSceneConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s SetCasterSceneConfigResponse) GoString() string {
	return s.String()
}

func (s *SetCasterSceneConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetCasterSceneConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetCasterSceneConfigResponse) GetBody() *SetCasterSceneConfigResponseBody {
	return s.Body
}

func (s *SetCasterSceneConfigResponse) SetHeaders(v map[string]*string) *SetCasterSceneConfigResponse {
	s.Headers = v
	return s
}

func (s *SetCasterSceneConfigResponse) SetStatusCode(v int32) *SetCasterSceneConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *SetCasterSceneConfigResponse) SetBody(v *SetCasterSceneConfigResponseBody) *SetCasterSceneConfigResponse {
	s.Body = v
	return s
}

func (s *SetCasterSceneConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
