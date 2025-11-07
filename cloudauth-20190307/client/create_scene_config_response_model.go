// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSceneConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSceneConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSceneConfigResponse
	GetStatusCode() *int32
	SetBody(v *CreateSceneConfigResponseBody) *CreateSceneConfigResponse
	GetBody() *CreateSceneConfigResponseBody
}

type CreateSceneConfigResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSceneConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSceneConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSceneConfigResponse) GoString() string {
	return s.String()
}

func (s *CreateSceneConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSceneConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSceneConfigResponse) GetBody() *CreateSceneConfigResponseBody {
	return s.Body
}

func (s *CreateSceneConfigResponse) SetHeaders(v map[string]*string) *CreateSceneConfigResponse {
	s.Headers = v
	return s
}

func (s *CreateSceneConfigResponse) SetStatusCode(v int32) *CreateSceneConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSceneConfigResponse) SetBody(v *CreateSceneConfigResponseBody) *CreateSceneConfigResponse {
	s.Body = v
	return s
}

func (s *CreateSceneConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
