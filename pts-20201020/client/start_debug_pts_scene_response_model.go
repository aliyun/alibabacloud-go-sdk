// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDebugPtsSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartDebugPtsSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartDebugPtsSceneResponse
	GetStatusCode() *int32
	SetBody(v *StartDebugPtsSceneResponseBody) *StartDebugPtsSceneResponse
	GetBody() *StartDebugPtsSceneResponseBody
}

type StartDebugPtsSceneResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartDebugPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartDebugPtsSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s StartDebugPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *StartDebugPtsSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartDebugPtsSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartDebugPtsSceneResponse) GetBody() *StartDebugPtsSceneResponseBody {
	return s.Body
}

func (s *StartDebugPtsSceneResponse) SetHeaders(v map[string]*string) *StartDebugPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *StartDebugPtsSceneResponse) SetStatusCode(v int32) *StartDebugPtsSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *StartDebugPtsSceneResponse) SetBody(v *StartDebugPtsSceneResponseBody) *StartDebugPtsSceneResponse {
	s.Body = v
	return s
}

func (s *StartDebugPtsSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
