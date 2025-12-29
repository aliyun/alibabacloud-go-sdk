// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRcuSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRcuSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRcuSceneResponse
	GetStatusCode() *int32
	SetBody(v *CreateRcuSceneResponseBody) *CreateRcuSceneResponse
	GetBody() *CreateRcuSceneResponseBody
}

type CreateRcuSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRcuSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRcuSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRcuSceneResponse) GoString() string {
	return s.String()
}

func (s *CreateRcuSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRcuSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRcuSceneResponse) GetBody() *CreateRcuSceneResponseBody {
	return s.Body
}

func (s *CreateRcuSceneResponse) SetHeaders(v map[string]*string) *CreateRcuSceneResponse {
	s.Headers = v
	return s
}

func (s *CreateRcuSceneResponse) SetStatusCode(v int32) *CreateRcuSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRcuSceneResponse) SetBody(v *CreateRcuSceneResponseBody) *CreateRcuSceneResponse {
	s.Body = v
	return s
}

func (s *CreateRcuSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
