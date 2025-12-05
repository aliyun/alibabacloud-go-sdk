// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPtsSceneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartPtsSceneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartPtsSceneResponse
	GetStatusCode() *int32
	SetBody(v *StartPtsSceneResponseBody) *StartPtsSceneResponse
	GetBody() *StartPtsSceneResponseBody
}

type StartPtsSceneResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartPtsSceneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartPtsSceneResponse) String() string {
	return dara.Prettify(s)
}

func (s StartPtsSceneResponse) GoString() string {
	return s.String()
}

func (s *StartPtsSceneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartPtsSceneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartPtsSceneResponse) GetBody() *StartPtsSceneResponseBody {
	return s.Body
}

func (s *StartPtsSceneResponse) SetHeaders(v map[string]*string) *StartPtsSceneResponse {
	s.Headers = v
	return s
}

func (s *StartPtsSceneResponse) SetStatusCode(v int32) *StartPtsSceneResponse {
	s.StatusCode = &v
	return s
}

func (s *StartPtsSceneResponse) SetBody(v *StartPtsSceneResponseBody) *StartPtsSceneResponse {
	s.Body = v
	return s
}

func (s *StartPtsSceneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
