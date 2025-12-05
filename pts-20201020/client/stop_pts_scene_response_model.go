// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopPtsSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopPtsSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopPtsSceneResponse
	GetStatusCode() *int32
	SetBody(v *StopPtsSceneResponseBody) *StopPtsSceneResponse
	GetBody() *StopPtsSceneResponseBody
}

type StopPtsSceneResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopPtsSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s StopPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *StopPtsSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopPtsSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopPtsSceneResponse) GetBody() *StopPtsSceneResponseBody {
	return s.Body
}

func (s *StopPtsSceneResponse) SetHeaders(v map[string]*string) *StopPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *StopPtsSceneResponse) SetStatusCode(v int32) *StopPtsSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *StopPtsSceneResponse) SetBody(v *StopPtsSceneResponseBody) *StopPtsSceneResponse {
	s.Body = v
	return s
}

func (s *StopPtsSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
