// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckImageExistsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckImageExistsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckImageExistsResponse
	GetStatusCode() *int32
	SetBody(v *CheckImageExistsResponseBody) *CheckImageExistsResponse
	GetBody() *CheckImageExistsResponseBody
}

type CheckImageExistsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckImageExistsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckImageExistsResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckImageExistsResponse) GoString() string {
	return s.String()
}

func (s *CheckImageExistsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckImageExistsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckImageExistsResponse) GetBody() *CheckImageExistsResponseBody {
	return s.Body
}

func (s *CheckImageExistsResponse) SetHeaders(v map[string]*string) *CheckImageExistsResponse {
	s.Headers = v
	return s
}

func (s *CheckImageExistsResponse) SetStatusCode(v int32) *CheckImageExistsResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckImageExistsResponse) SetBody(v *CheckImageExistsResponseBody) *CheckImageExistsResponse {
	s.Body = v
	return s
}

func (s *CheckImageExistsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
