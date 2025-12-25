// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetailSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetailSceneResponse
	GetStatusCode() *int32
	SetBody(v *DetailSceneResponseBody) *DetailSceneResponse
	GetBody() *DetailSceneResponseBody
}

type DetailSceneResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetailSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetailSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DetailSceneResponse) GoString() string {
	return s.String()
}

func (s *DetailSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetailSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetailSceneResponse) GetBody() *DetailSceneResponseBody {
	return s.Body
}

func (s *DetailSceneResponse) SetHeaders(v map[string]*string) *DetailSceneResponse {
	s.Headers = v
	return s
}

func (s *DetailSceneResponse) SetStatusCode(v int32) *DetailSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DetailSceneResponse) SetBody(v *DetailSceneResponseBody) *DetailSceneResponse {
	s.Body = v
	return s
}

func (s *DetailSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
