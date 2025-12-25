// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubSceneResponse
	GetStatusCode() *int32
	SetBody(v *ListSubSceneResponseBody) *ListSubSceneResponse
	GetBody() *ListSubSceneResponseBody
}

type ListSubSceneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubSceneResponse) GoString() string {
	return s.String()
}

func (s *ListSubSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubSceneResponse) GetBody() *ListSubSceneResponseBody {
	return s.Body
}

func (s *ListSubSceneResponse) SetHeaders(v map[string]*string) *ListSubSceneResponse {
	s.Headers = v
	return s
}

func (s *ListSubSceneResponse) SetStatusCode(v int32) *ListSubSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubSceneResponse) SetBody(v *ListSubSceneResponseBody) *ListSubSceneResponse {
	s.Body = v
	return s
}

func (s *ListSubSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
