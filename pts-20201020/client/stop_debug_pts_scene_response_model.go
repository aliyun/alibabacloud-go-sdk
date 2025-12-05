// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDebugPtsSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopDebugPtsSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopDebugPtsSceneResponse
	GetStatusCode() *int32
	SetBody(v *StopDebugPtsSceneResponseBody) *StopDebugPtsSceneResponse
	GetBody() *StopDebugPtsSceneResponseBody
}

type StopDebugPtsSceneResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopDebugPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopDebugPtsSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s StopDebugPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *StopDebugPtsSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopDebugPtsSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopDebugPtsSceneResponse) GetBody() *StopDebugPtsSceneResponseBody {
	return s.Body
}

func (s *StopDebugPtsSceneResponse) SetHeaders(v map[string]*string) *StopDebugPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *StopDebugPtsSceneResponse) SetStatusCode(v int32) *StopDebugPtsSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *StopDebugPtsSceneResponse) SetBody(v *StopDebugPtsSceneResponseBody) *StopDebugPtsSceneResponse {
	s.Body = v
	return s
}

func (s *StopDebugPtsSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
