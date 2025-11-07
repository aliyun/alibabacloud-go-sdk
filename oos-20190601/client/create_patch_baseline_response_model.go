// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePatchBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePatchBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePatchBaselineResponse
	GetStatusCode() *int32
	SetBody(v *CreatePatchBaselineResponseBody) *CreatePatchBaselineResponse
	GetBody() *CreatePatchBaselineResponseBody
}

type CreatePatchBaselineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePatchBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *CreatePatchBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePatchBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePatchBaselineResponse) GetBody() *CreatePatchBaselineResponseBody {
	return s.Body
}

func (s *CreatePatchBaselineResponse) SetHeaders(v map[string]*string) *CreatePatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *CreatePatchBaselineResponse) SetStatusCode(v int32) *CreatePatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePatchBaselineResponse) SetBody(v *CreatePatchBaselineResponseBody) *CreatePatchBaselineResponse {
	s.Body = v
	return s
}

func (s *CreatePatchBaselineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
