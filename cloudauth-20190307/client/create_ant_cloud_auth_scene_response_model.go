// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAntCloudAuthSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAntCloudAuthSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAntCloudAuthSceneResponse
	GetStatusCode() *int32
	SetBody(v *CreateAntCloudAuthSceneResponseBody) *CreateAntCloudAuthSceneResponse
	GetBody() *CreateAntCloudAuthSceneResponseBody
}

type CreateAntCloudAuthSceneResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAntCloudAuthSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAntCloudAuthSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAntCloudAuthSceneResponse) GoString() string {
	return s.String()
}

func (s *CreateAntCloudAuthSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAntCloudAuthSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAntCloudAuthSceneResponse) GetBody() *CreateAntCloudAuthSceneResponseBody {
	return s.Body
}

func (s *CreateAntCloudAuthSceneResponse) SetHeaders(v map[string]*string) *CreateAntCloudAuthSceneResponse {
	s.Headers = v
	return s
}

func (s *CreateAntCloudAuthSceneResponse) SetStatusCode(v int32) *CreateAntCloudAuthSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAntCloudAuthSceneResponse) SetBody(v *CreateAntCloudAuthSceneResponseBody) *CreateAntCloudAuthSceneResponse {
	s.Body = v
	return s
}

func (s *CreateAntCloudAuthSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
