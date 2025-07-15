// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterDefaultPatchBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterDefaultPatchBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterDefaultPatchBaselineResponse
	GetStatusCode() *int32
	SetBody(v *RegisterDefaultPatchBaselineResponseBody) *RegisterDefaultPatchBaselineResponse
	GetBody() *RegisterDefaultPatchBaselineResponseBody
}

type RegisterDefaultPatchBaselineResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterDefaultPatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterDefaultPatchBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterDefaultPatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *RegisterDefaultPatchBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterDefaultPatchBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterDefaultPatchBaselineResponse) GetBody() *RegisterDefaultPatchBaselineResponseBody {
	return s.Body
}

func (s *RegisterDefaultPatchBaselineResponse) SetHeaders(v map[string]*string) *RegisterDefaultPatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *RegisterDefaultPatchBaselineResponse) SetStatusCode(v int32) *RegisterDefaultPatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterDefaultPatchBaselineResponse) SetBody(v *RegisterDefaultPatchBaselineResponseBody) *RegisterDefaultPatchBaselineResponse {
	s.Body = v
	return s
}

func (s *RegisterDefaultPatchBaselineResponse) Validate() error {
	return dara.Validate(s)
}
