// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProbeMcpOriginResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProbeMcpOriginResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProbeMcpOriginResponse
	GetStatusCode() *int32
	SetBody(v *ProbeMcpOriginResponseBody) *ProbeMcpOriginResponse
	GetBody() *ProbeMcpOriginResponseBody
}

type ProbeMcpOriginResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProbeMcpOriginResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProbeMcpOriginResponse) String() string {
	return dara.Prettify(s)
}

func (s ProbeMcpOriginResponse) GoString() string {
	return s.String()
}

func (s *ProbeMcpOriginResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProbeMcpOriginResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProbeMcpOriginResponse) GetBody() *ProbeMcpOriginResponseBody {
	return s.Body
}

func (s *ProbeMcpOriginResponse) SetHeaders(v map[string]*string) *ProbeMcpOriginResponse {
	s.Headers = v
	return s
}

func (s *ProbeMcpOriginResponse) SetStatusCode(v int32) *ProbeMcpOriginResponse {
	s.StatusCode = &v
	return s
}

func (s *ProbeMcpOriginResponse) SetBody(v *ProbeMcpOriginResponseBody) *ProbeMcpOriginResponse {
	s.Body = v
	return s
}

func (s *ProbeMcpOriginResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
