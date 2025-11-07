// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudauthstSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudauthstSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudauthstSceneResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudauthstSceneResponseBody) *DeleteCloudauthstSceneResponse
	GetBody() *DeleteCloudauthstSceneResponseBody
}

type DeleteCloudauthstSceneResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudauthstSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudauthstSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudauthstSceneResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudauthstSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudauthstSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudauthstSceneResponse) GetBody() *DeleteCloudauthstSceneResponseBody {
	return s.Body
}

func (s *DeleteCloudauthstSceneResponse) SetHeaders(v map[string]*string) *DeleteCloudauthstSceneResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudauthstSceneResponse) SetStatusCode(v int32) *DeleteCloudauthstSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudauthstSceneResponse) SetBody(v *DeleteCloudauthstSceneResponseBody) *DeleteCloudauthstSceneResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudauthstSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
