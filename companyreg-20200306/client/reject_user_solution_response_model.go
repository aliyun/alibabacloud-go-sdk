// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRejectUserSolutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RejectUserSolutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RejectUserSolutionResponse
	GetStatusCode() *int32
	SetBody(v *RejectUserSolutionResponseBody) *RejectUserSolutionResponse
	GetBody() *RejectUserSolutionResponseBody
}

type RejectUserSolutionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RejectUserSolutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RejectUserSolutionResponse) String() string {
	return dara.Prettify(s)
}

func (s RejectUserSolutionResponse) GoString() string {
	return s.String()
}

func (s *RejectUserSolutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RejectUserSolutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RejectUserSolutionResponse) GetBody() *RejectUserSolutionResponseBody {
	return s.Body
}

func (s *RejectUserSolutionResponse) SetHeaders(v map[string]*string) *RejectUserSolutionResponse {
	s.Headers = v
	return s
}

func (s *RejectUserSolutionResponse) SetStatusCode(v int32) *RejectUserSolutionResponse {
	s.StatusCode = &v
	return s
}

func (s *RejectUserSolutionResponse) SetBody(v *RejectUserSolutionResponseBody) *RejectUserSolutionResponse {
	s.Body = v
	return s
}

func (s *RejectUserSolutionResponse) Validate() error {
	return dara.Validate(s)
}
