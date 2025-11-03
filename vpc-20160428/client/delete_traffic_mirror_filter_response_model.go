// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrafficMirrorFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTrafficMirrorFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTrafficMirrorFilterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTrafficMirrorFilterResponseBody) *DeleteTrafficMirrorFilterResponse
	GetBody() *DeleteTrafficMirrorFilterResponseBody
}

type DeleteTrafficMirrorFilterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrafficMirrorFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrafficMirrorFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrafficMirrorFilterResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrafficMirrorFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTrafficMirrorFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTrafficMirrorFilterResponse) GetBody() *DeleteTrafficMirrorFilterResponseBody {
	return s.Body
}

func (s *DeleteTrafficMirrorFilterResponse) SetHeaders(v map[string]*string) *DeleteTrafficMirrorFilterResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrafficMirrorFilterResponse) SetStatusCode(v int32) *DeleteTrafficMirrorFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrafficMirrorFilterResponse) SetBody(v *DeleteTrafficMirrorFilterResponseBody) *DeleteTrafficMirrorFilterResponse {
	s.Body = v
	return s
}

func (s *DeleteTrafficMirrorFilterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
