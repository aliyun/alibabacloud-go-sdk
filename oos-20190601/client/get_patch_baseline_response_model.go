// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPatchBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetPatchBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetPatchBaselineResponse
	GetStatusCode() *int32
	SetBody(v *GetPatchBaselineResponseBody) *GetPatchBaselineResponse
	GetBody() *GetPatchBaselineResponseBody
}

type GetPatchBaselineResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetPatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetPatchBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s GetPatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *GetPatchBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetPatchBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetPatchBaselineResponse) GetBody() *GetPatchBaselineResponseBody {
	return s.Body
}

func (s *GetPatchBaselineResponse) SetHeaders(v map[string]*string) *GetPatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *GetPatchBaselineResponse) SetStatusCode(v int32) *GetPatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *GetPatchBaselineResponse) SetBody(v *GetPatchBaselineResponseBody) *GetPatchBaselineResponse {
	s.Body = v
	return s
}

func (s *GetPatchBaselineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
