// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePtsSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePtsSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePtsSceneResponse
	GetStatusCode() *int32
	SetBody(v *DeletePtsSceneResponseBody) *DeletePtsSceneResponse
	GetBody() *DeletePtsSceneResponseBody
}

type DeletePtsSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePtsSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePtsSceneResponse) GoString() string {
	return s.String()
}

func (s *DeletePtsSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePtsSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePtsSceneResponse) GetBody() *DeletePtsSceneResponseBody {
	return s.Body
}

func (s *DeletePtsSceneResponse) SetHeaders(v map[string]*string) *DeletePtsSceneResponse {
	s.Headers = v
	return s
}

func (s *DeletePtsSceneResponse) SetStatusCode(v int32) *DeletePtsSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePtsSceneResponse) SetBody(v *DeletePtsSceneResponseBody) *DeletePtsSceneResponse {
	s.Body = v
	return s
}

func (s *DeletePtsSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
