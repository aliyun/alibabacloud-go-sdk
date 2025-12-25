// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPackSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PackSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PackSceneResponse
	GetStatusCode() *int32
	SetBody(v *PackSceneResponseBody) *PackSceneResponse
	GetBody() *PackSceneResponseBody
}

type PackSceneResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PackSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PackSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s PackSceneResponse) GoString() string {
	return s.String()
}

func (s *PackSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PackSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PackSceneResponse) GetBody() *PackSceneResponseBody {
	return s.Body
}

func (s *PackSceneResponse) SetHeaders(v map[string]*string) *PackSceneResponse {
	s.Headers = v
	return s
}

func (s *PackSceneResponse) SetStatusCode(v int32) *PackSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *PackSceneResponse) SetBody(v *PackSceneResponseBody) *PackSceneResponse {
	s.Body = v
	return s
}

func (s *PackSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
