// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePatchBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePatchBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePatchBaselineResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePatchBaselineResponseBody) *UpdatePatchBaselineResponse
	GetBody() *UpdatePatchBaselineResponseBody
}

type UpdatePatchBaselineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePatchBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *UpdatePatchBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePatchBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePatchBaselineResponse) GetBody() *UpdatePatchBaselineResponseBody {
	return s.Body
}

func (s *UpdatePatchBaselineResponse) SetHeaders(v map[string]*string) *UpdatePatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *UpdatePatchBaselineResponse) SetStatusCode(v int32) *UpdatePatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePatchBaselineResponse) SetBody(v *UpdatePatchBaselineResponseBody) *UpdatePatchBaselineResponse {
	s.Body = v
	return s
}

func (s *UpdatePatchBaselineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
