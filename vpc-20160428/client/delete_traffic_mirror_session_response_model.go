// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficMirrorSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTrafficMirrorSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTrafficMirrorSessionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTrafficMirrorSessionResponseBody) *DeleteTrafficMirrorSessionResponse
	GetBody() *DeleteTrafficMirrorSessionResponseBody
}

type DeleteTrafficMirrorSessionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrafficMirrorSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrafficMirrorSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficMirrorSessionResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMirrorSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTrafficMirrorSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTrafficMirrorSessionResponse) GetBody() *DeleteTrafficMirrorSessionResponseBody {
	return s.Body
}

func (s *DeleteTrafficMirrorSessionResponse) SetHeaders(v map[string]*string) *DeleteTrafficMirrorSessionResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrafficMirrorSessionResponse) SetStatusCode(v int32) *DeleteTrafficMirrorSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrafficMirrorSessionResponse) SetBody(v *DeleteTrafficMirrorSessionResponseBody) *DeleteTrafficMirrorSessionResponse {
	s.Body = v
	return s
}

func (s *DeleteTrafficMirrorSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
