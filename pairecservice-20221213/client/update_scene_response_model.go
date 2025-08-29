// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSceneResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSceneResponseBody) *UpdateSceneResponse
	GetBody() *UpdateSceneResponseBody
}

type UpdateSceneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSceneResponse) GoString() string {
	return s.String()
}

func (s *UpdateSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSceneResponse) GetBody() *UpdateSceneResponseBody {
	return s.Body
}

func (s *UpdateSceneResponse) SetHeaders(v map[string]*string) *UpdateSceneResponse {
	s.Headers = v
	return s
}

func (s *UpdateSceneResponse) SetStatusCode(v int32) *UpdateSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSceneResponse) SetBody(v *UpdateSceneResponseBody) *UpdateSceneResponse {
	s.Body = v
	return s
}

func (s *UpdateSceneResponse) Validate() error {
	return dara.Validate(s)
}
