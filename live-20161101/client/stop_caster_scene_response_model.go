// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopCasterSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopCasterSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopCasterSceneResponse
	GetStatusCode() *int32
	SetBody(v *StopCasterSceneResponseBody) *StopCasterSceneResponse
	GetBody() *StopCasterSceneResponseBody
}

type StopCasterSceneResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopCasterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopCasterSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s StopCasterSceneResponse) GoString() string {
	return s.String()
}

func (s *StopCasterSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopCasterSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopCasterSceneResponse) GetBody() *StopCasterSceneResponseBody {
	return s.Body
}

func (s *StopCasterSceneResponse) SetHeaders(v map[string]*string) *StopCasterSceneResponse {
	s.Headers = v
	return s
}

func (s *StopCasterSceneResponse) SetStatusCode(v int32) *StopCasterSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *StopCasterSceneResponse) SetBody(v *StopCasterSceneResponseBody) *StopCasterSceneResponse {
	s.Body = v
	return s
}

func (s *StopCasterSceneResponse) Validate() error {
	return dara.Validate(s)
}
