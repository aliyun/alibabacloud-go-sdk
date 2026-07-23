// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReDoRoutineBuildResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReDoRoutineBuildResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReDoRoutineBuildResponse
	GetStatusCode() *int32
	SetBody(v *ReDoRoutineBuildResponseBody) *ReDoRoutineBuildResponse
	GetBody() *ReDoRoutineBuildResponseBody
}

type ReDoRoutineBuildResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReDoRoutineBuildResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReDoRoutineBuildResponse) String() string {
	return dara.Prettify(s)
}

func (s ReDoRoutineBuildResponse) GoString() string {
	return s.String()
}

func (s *ReDoRoutineBuildResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReDoRoutineBuildResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReDoRoutineBuildResponse) GetBody() *ReDoRoutineBuildResponseBody {
	return s.Body
}

func (s *ReDoRoutineBuildResponse) SetHeaders(v map[string]*string) *ReDoRoutineBuildResponse {
	s.Headers = v
	return s
}

func (s *ReDoRoutineBuildResponse) SetStatusCode(v int32) *ReDoRoutineBuildResponse {
	s.StatusCode = &v
	return s
}

func (s *ReDoRoutineBuildResponse) SetBody(v *ReDoRoutineBuildResponseBody) *ReDoRoutineBuildResponse {
	s.Body = v
	return s
}

func (s *ReDoRoutineBuildResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
