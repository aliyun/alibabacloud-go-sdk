// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSceneResponse
	GetStatusCode() *int32
	SetBody(v *CreateSceneResponseBody) *CreateSceneResponse
	GetBody() *CreateSceneResponseBody
}

type CreateSceneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSceneResponse) GoString() string {
	return s.String()
}

func (s *CreateSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSceneResponse) GetBody() *CreateSceneResponseBody {
	return s.Body
}

func (s *CreateSceneResponse) SetHeaders(v map[string]*string) *CreateSceneResponse {
	s.Headers = v
	return s
}

func (s *CreateSceneResponse) SetStatusCode(v int32) *CreateSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSceneResponse) SetBody(v *CreateSceneResponseBody) *CreateSceneResponse {
	s.Body = v
	return s
}

func (s *CreateSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
