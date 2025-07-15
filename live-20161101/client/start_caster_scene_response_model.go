// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCasterSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartCasterSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartCasterSceneResponse
	GetStatusCode() *int32
	SetBody(v *StartCasterSceneResponseBody) *StartCasterSceneResponse
	GetBody() *StartCasterSceneResponseBody
}

type StartCasterSceneResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCasterSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartCasterSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s StartCasterSceneResponse) GoString() string {
	return s.String()
}

func (s *StartCasterSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartCasterSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartCasterSceneResponse) GetBody() *StartCasterSceneResponseBody {
	return s.Body
}

func (s *StartCasterSceneResponse) SetHeaders(v map[string]*string) *StartCasterSceneResponse {
	s.Headers = v
	return s
}

func (s *StartCasterSceneResponse) SetStatusCode(v int32) *StartCasterSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCasterSceneResponse) SetBody(v *StartCasterSceneResponseBody) *StartCasterSceneResponse {
	s.Body = v
	return s
}

func (s *StartCasterSceneResponse) Validate() error {
	return dara.Validate(s)
}
