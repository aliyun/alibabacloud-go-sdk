// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddSourcesToTrafficMirrorSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddSourcesToTrafficMirrorSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddSourcesToTrafficMirrorSessionResponse
	GetStatusCode() *int32
	SetBody(v *AddSourcesToTrafficMirrorSessionResponseBody) *AddSourcesToTrafficMirrorSessionResponse
	GetBody() *AddSourcesToTrafficMirrorSessionResponseBody
}

type AddSourcesToTrafficMirrorSessionResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddSourcesToTrafficMirrorSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddSourcesToTrafficMirrorSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s AddSourcesToTrafficMirrorSessionResponse) GoString() string {
	return s.String()
}

func (s *AddSourcesToTrafficMirrorSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddSourcesToTrafficMirrorSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddSourcesToTrafficMirrorSessionResponse) GetBody() *AddSourcesToTrafficMirrorSessionResponseBody {
	return s.Body
}

func (s *AddSourcesToTrafficMirrorSessionResponse) SetHeaders(v map[string]*string) *AddSourcesToTrafficMirrorSessionResponse {
	s.Headers = v
	return s
}

func (s *AddSourcesToTrafficMirrorSessionResponse) SetStatusCode(v int32) *AddSourcesToTrafficMirrorSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *AddSourcesToTrafficMirrorSessionResponse) SetBody(v *AddSourcesToTrafficMirrorSessionResponseBody) *AddSourcesToTrafficMirrorSessionResponse {
	s.Body = v
	return s
}

func (s *AddSourcesToTrafficMirrorSessionResponse) Validate() error {
	return dara.Validate(s)
}
