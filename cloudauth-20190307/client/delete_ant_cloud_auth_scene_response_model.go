// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAntCloudAuthSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAntCloudAuthSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAntCloudAuthSceneResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAntCloudAuthSceneResponseBody) *DeleteAntCloudAuthSceneResponse
	GetBody() *DeleteAntCloudAuthSceneResponseBody
}

type DeleteAntCloudAuthSceneResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAntCloudAuthSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAntCloudAuthSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAntCloudAuthSceneResponse) GoString() string {
	return s.String()
}

func (s *DeleteAntCloudAuthSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAntCloudAuthSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAntCloudAuthSceneResponse) GetBody() *DeleteAntCloudAuthSceneResponseBody {
	return s.Body
}

func (s *DeleteAntCloudAuthSceneResponse) SetHeaders(v map[string]*string) *DeleteAntCloudAuthSceneResponse {
	s.Headers = v
	return s
}

func (s *DeleteAntCloudAuthSceneResponse) SetStatusCode(v int32) *DeleteAntCloudAuthSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAntCloudAuthSceneResponse) SetBody(v *DeleteAntCloudAuthSceneResponseBody) *DeleteAntCloudAuthSceneResponse {
	s.Body = v
	return s
}

func (s *DeleteAntCloudAuthSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
