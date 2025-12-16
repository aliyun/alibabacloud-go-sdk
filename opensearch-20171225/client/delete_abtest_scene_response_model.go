// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteABTestSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteABTestSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteABTestSceneResponse
	GetStatusCode() *int32
	SetBody(v *DeleteABTestSceneResponseBody) *DeleteABTestSceneResponse
	GetBody() *DeleteABTestSceneResponseBody
}

type DeleteABTestSceneResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteABTestSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteABTestSceneResponse) GoString() string {
	return s.String()
}

func (s *DeleteABTestSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteABTestSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteABTestSceneResponse) GetBody() *DeleteABTestSceneResponseBody {
	return s.Body
}

func (s *DeleteABTestSceneResponse) SetHeaders(v map[string]*string) *DeleteABTestSceneResponse {
	s.Headers = v
	return s
}

func (s *DeleteABTestSceneResponse) SetStatusCode(v int32) *DeleteABTestSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteABTestSceneResponse) SetBody(v *DeleteABTestSceneResponseBody) *DeleteABTestSceneResponse {
	s.Body = v
	return s
}

func (s *DeleteABTestSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
