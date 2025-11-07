// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAntCloudAuthSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAntCloudAuthSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAntCloudAuthSceneResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAntCloudAuthSceneResponseBody) *UpdateAntCloudAuthSceneResponse
	GetBody() *UpdateAntCloudAuthSceneResponseBody
}

type UpdateAntCloudAuthSceneResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAntCloudAuthSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAntCloudAuthSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAntCloudAuthSceneResponse) GoString() string {
	return s.String()
}

func (s *UpdateAntCloudAuthSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAntCloudAuthSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAntCloudAuthSceneResponse) GetBody() *UpdateAntCloudAuthSceneResponseBody {
	return s.Body
}

func (s *UpdateAntCloudAuthSceneResponse) SetHeaders(v map[string]*string) *UpdateAntCloudAuthSceneResponse {
	s.Headers = v
	return s
}

func (s *UpdateAntCloudAuthSceneResponse) SetStatusCode(v int32) *UpdateAntCloudAuthSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAntCloudAuthSceneResponse) SetBody(v *UpdateAntCloudAuthSceneResponseBody) *UpdateAntCloudAuthSceneResponse {
	s.Body = v
	return s
}

func (s *UpdateAntCloudAuthSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
