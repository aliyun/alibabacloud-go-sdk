// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRcuSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRcuSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRcuSceneResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRcuSceneResponseBody) *DeleteRcuSceneResponse
	GetBody() *DeleteRcuSceneResponseBody
}

type DeleteRcuSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRcuSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRcuSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRcuSceneResponse) GoString() string {
	return s.String()
}

func (s *DeleteRcuSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRcuSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRcuSceneResponse) GetBody() *DeleteRcuSceneResponseBody {
	return s.Body
}

func (s *DeleteRcuSceneResponse) SetHeaders(v map[string]*string) *DeleteRcuSceneResponse {
	s.Headers = v
	return s
}

func (s *DeleteRcuSceneResponse) SetStatusCode(v int32) *DeleteRcuSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRcuSceneResponse) SetBody(v *DeleteRcuSceneResponseBody) *DeleteRcuSceneResponse {
	s.Body = v
	return s
}

func (s *DeleteRcuSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
