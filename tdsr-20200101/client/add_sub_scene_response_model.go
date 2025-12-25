// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSubSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSubSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSubSceneResponse
	GetStatusCode() *int32
	SetBody(v *AddSubSceneResponseBody) *AddSubSceneResponse
	GetBody() *AddSubSceneResponseBody
}

type AddSubSceneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSubSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSubSceneResponse) GoString() string {
	return s.String()
}

func (s *AddSubSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSubSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSubSceneResponse) GetBody() *AddSubSceneResponseBody {
	return s.Body
}

func (s *AddSubSceneResponse) SetHeaders(v map[string]*string) *AddSubSceneResponse {
	s.Headers = v
	return s
}

func (s *AddSubSceneResponse) SetStatusCode(v int32) *AddSubSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSubSceneResponse) SetBody(v *AddSubSceneResponseBody) *AddSubSceneResponse {
	s.Body = v
	return s
}

func (s *AddSubSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
