// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSubSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSubSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSubSceneResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSubSceneResponseBody) *UpdateSubSceneResponse
	GetBody() *UpdateSubSceneResponseBody
}

type UpdateSubSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSubSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSubSceneResponse) GoString() string {
	return s.String()
}

func (s *UpdateSubSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSubSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSubSceneResponse) GetBody() *UpdateSubSceneResponseBody {
	return s.Body
}

func (s *UpdateSubSceneResponse) SetHeaders(v map[string]*string) *UpdateSubSceneResponse {
	s.Headers = v
	return s
}

func (s *UpdateSubSceneResponse) SetStatusCode(v int32) *UpdateSubSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSubSceneResponse) SetBody(v *UpdateSubSceneResponseBody) *UpdateSubSceneResponse {
	s.Body = v
	return s
}

func (s *UpdateSubSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
