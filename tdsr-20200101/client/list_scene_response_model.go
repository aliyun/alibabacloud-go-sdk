// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSceneResponse
	GetStatusCode() *int32
	SetBody(v *ListSceneResponseBody) *ListSceneResponse
	GetBody() *ListSceneResponseBody
}

type ListSceneResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSceneResponse) GoString() string {
	return s.String()
}

func (s *ListSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSceneResponse) GetBody() *ListSceneResponseBody {
	return s.Body
}

func (s *ListSceneResponse) SetHeaders(v map[string]*string) *ListSceneResponse {
	s.Headers = v
	return s
}

func (s *ListSceneResponse) SetStatusCode(v int32) *ListSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSceneResponse) SetBody(v *ListSceneResponseBody) *ListSceneResponse {
	s.Body = v
	return s
}

func (s *ListSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
