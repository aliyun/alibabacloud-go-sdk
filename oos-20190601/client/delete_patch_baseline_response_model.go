// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePatchBaselineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePatchBaselineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePatchBaselineResponse
	GetStatusCode() *int32
	SetBody(v *DeletePatchBaselineResponseBody) *DeletePatchBaselineResponse
	GetBody() *DeletePatchBaselineResponseBody
}

type DeletePatchBaselineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePatchBaselineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePatchBaselineResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePatchBaselineResponse) GoString() string {
	return s.String()
}

func (s *DeletePatchBaselineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePatchBaselineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePatchBaselineResponse) GetBody() *DeletePatchBaselineResponseBody {
	return s.Body
}

func (s *DeletePatchBaselineResponse) SetHeaders(v map[string]*string) *DeletePatchBaselineResponse {
	s.Headers = v
	return s
}

func (s *DeletePatchBaselineResponse) SetStatusCode(v int32) *DeletePatchBaselineResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePatchBaselineResponse) SetBody(v *DeletePatchBaselineResponseBody) *DeletePatchBaselineResponse {
	s.Body = v
	return s
}

func (s *DeletePatchBaselineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
