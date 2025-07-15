// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficMirrorFilterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrafficMirrorFilterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrafficMirrorFilterResponse
	GetStatusCode() *int32
	SetBody(v *CreateTrafficMirrorFilterResponseBody) *CreateTrafficMirrorFilterResponse
	GetBody() *CreateTrafficMirrorFilterResponseBody
}

type CreateTrafficMirrorFilterResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrafficMirrorFilterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrafficMirrorFilterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficMirrorFilterResponse) GoString() string {
	return s.String()
}

func (s *CreateTrafficMirrorFilterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrafficMirrorFilterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrafficMirrorFilterResponse) GetBody() *CreateTrafficMirrorFilterResponseBody {
	return s.Body
}

func (s *CreateTrafficMirrorFilterResponse) SetHeaders(v map[string]*string) *CreateTrafficMirrorFilterResponse {
	s.Headers = v
	return s
}

func (s *CreateTrafficMirrorFilterResponse) SetStatusCode(v int32) *CreateTrafficMirrorFilterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrafficMirrorFilterResponse) SetBody(v *CreateTrafficMirrorFilterResponseBody) *CreateTrafficMirrorFilterResponse {
	s.Body = v
	return s
}

func (s *CreateTrafficMirrorFilterResponse) Validate() error {
	return dara.Validate(s)
}
