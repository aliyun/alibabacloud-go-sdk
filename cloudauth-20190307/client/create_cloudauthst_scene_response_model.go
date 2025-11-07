// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudauthstSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudauthstSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudauthstSceneResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudauthstSceneResponseBody) *CreateCloudauthstSceneResponse
	GetBody() *CreateCloudauthstSceneResponseBody
}

type CreateCloudauthstSceneResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudauthstSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudauthstSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudauthstSceneResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudauthstSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudauthstSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudauthstSceneResponse) GetBody() *CreateCloudauthstSceneResponseBody {
	return s.Body
}

func (s *CreateCloudauthstSceneResponse) SetHeaders(v map[string]*string) *CreateCloudauthstSceneResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudauthstSceneResponse) SetStatusCode(v int32) *CreateCloudauthstSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudauthstSceneResponse) SetBody(v *CreateCloudauthstSceneResponseBody) *CreateCloudauthstSceneResponse {
	s.Body = v
	return s
}

func (s *CreateCloudauthstSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
