// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSceneResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSceneResponseBody) *DeleteSceneResponse
	GetBody() *DeleteSceneResponseBody
}

type DeleteSceneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSceneResponse) GoString() string {
	return s.String()
}

func (s *DeleteSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSceneResponse) GetBody() *DeleteSceneResponseBody {
	return s.Body
}

func (s *DeleteSceneResponse) SetHeaders(v map[string]*string) *DeleteSceneResponse {
	s.Headers = v
	return s
}

func (s *DeleteSceneResponse) SetStatusCode(v int32) *DeleteSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSceneResponse) SetBody(v *DeleteSceneResponseBody) *DeleteSceneResponse {
	s.Body = v
	return s
}

func (s *DeleteSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
