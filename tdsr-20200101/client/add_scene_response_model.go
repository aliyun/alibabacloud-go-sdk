// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSceneResponse
	GetStatusCode() *int32
	SetBody(v *AddSceneResponseBody) *AddSceneResponse
	GetBody() *AddSceneResponseBody
}

type AddSceneResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSceneResponse) GoString() string {
	return s.String()
}

func (s *AddSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSceneResponse) GetBody() *AddSceneResponseBody {
	return s.Body
}

func (s *AddSceneResponse) SetHeaders(v map[string]*string) *AddSceneResponse {
	s.Headers = v
	return s
}

func (s *AddSceneResponse) SetStatusCode(v int32) *AddSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSceneResponse) SetBody(v *AddSceneResponseBody) *AddSceneResponse {
	s.Body = v
	return s
}

func (s *AddSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
