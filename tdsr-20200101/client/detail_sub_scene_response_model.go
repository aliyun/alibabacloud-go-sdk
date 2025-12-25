// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetailSubSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetailSubSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetailSubSceneResponse
	GetStatusCode() *int32
	SetBody(v *DetailSubSceneResponseBody) *DetailSubSceneResponse
	GetBody() *DetailSubSceneResponseBody
}

type DetailSubSceneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetailSubSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetailSubSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s DetailSubSceneResponse) GoString() string {
	return s.String()
}

func (s *DetailSubSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetailSubSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetailSubSceneResponse) GetBody() *DetailSubSceneResponseBody {
	return s.Body
}

func (s *DetailSubSceneResponse) SetHeaders(v map[string]*string) *DetailSubSceneResponse {
	s.Headers = v
	return s
}

func (s *DetailSubSceneResponse) SetStatusCode(v int32) *DetailSubSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *DetailSubSceneResponse) SetBody(v *DetailSubSceneResponseBody) *DetailSubSceneResponse {
	s.Body = v
	return s
}

func (s *DetailSubSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
