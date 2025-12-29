// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRcuSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRcuSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRcuSceneResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRcuSceneResponseBody) *UpdateRcuSceneResponse
	GetBody() *UpdateRcuSceneResponseBody
}

type UpdateRcuSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRcuSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRcuSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRcuSceneResponse) GoString() string {
	return s.String()
}

func (s *UpdateRcuSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRcuSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRcuSceneResponse) GetBody() *UpdateRcuSceneResponseBody {
	return s.Body
}

func (s *UpdateRcuSceneResponse) SetHeaders(v map[string]*string) *UpdateRcuSceneResponse {
	s.Headers = v
	return s
}

func (s *UpdateRcuSceneResponse) SetStatusCode(v int32) *UpdateRcuSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRcuSceneResponse) SetBody(v *UpdateRcuSceneResponseBody) *UpdateRcuSceneResponse {
	s.Body = v
	return s
}

func (s *UpdateRcuSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
