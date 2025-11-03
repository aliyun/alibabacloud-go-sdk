// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficMirrorSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrafficMirrorSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrafficMirrorSessionResponse
	GetStatusCode() *int32
	SetBody(v *CreateTrafficMirrorSessionResponseBody) *CreateTrafficMirrorSessionResponse
	GetBody() *CreateTrafficMirrorSessionResponseBody
}

type CreateTrafficMirrorSessionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrafficMirrorSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrafficMirrorSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrafficMirrorSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrafficMirrorSessionResponse) GetBody() *CreateTrafficMirrorSessionResponseBody {
	return s.Body
}

func (s *CreateTrafficMirrorSessionResponse) SetHeaders(v map[string]*string) *CreateTrafficMirrorSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateTrafficMirrorSessionResponse) SetStatusCode(v int32) *CreateTrafficMirrorSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrafficMirrorSessionResponse) SetBody(v *CreateTrafficMirrorSessionResponseBody) *CreateTrafficMirrorSessionResponse {
	s.Body = v
	return s
}

func (s *CreateTrafficMirrorSessionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
