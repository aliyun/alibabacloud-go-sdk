// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABTestSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateABTestSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateABTestSceneResponse
	GetStatusCode() *int32
	SetBody(v *UpdateABTestSceneResponseBody) *UpdateABTestSceneResponse
	GetBody() *UpdateABTestSceneResponseBody
}

type UpdateABTestSceneResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateABTestSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateABTestSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestSceneResponse) GoString() string {
	return s.String()
}

func (s *UpdateABTestSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateABTestSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateABTestSceneResponse) GetBody() *UpdateABTestSceneResponseBody {
	return s.Body
}

func (s *UpdateABTestSceneResponse) SetHeaders(v map[string]*string) *UpdateABTestSceneResponse {
	s.Headers = v
	return s
}

func (s *UpdateABTestSceneResponse) SetStatusCode(v int32) *UpdateABTestSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateABTestSceneResponse) SetBody(v *UpdateABTestSceneResponseBody) *UpdateABTestSceneResponse {
	s.Body = v
	return s
}

func (s *UpdateABTestSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
