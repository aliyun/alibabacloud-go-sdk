// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRoutineBuildResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRoutineBuildResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRoutineBuildResponse
	GetStatusCode() *int32
	SetBody(v *CreateRoutineBuildResponseBody) *CreateRoutineBuildResponse
	GetBody() *CreateRoutineBuildResponseBody
}

type CreateRoutineBuildResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRoutineBuildResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRoutineBuildResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRoutineBuildResponse) GoString() string {
	return s.String()
}

func (s *CreateRoutineBuildResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRoutineBuildResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRoutineBuildResponse) GetBody() *CreateRoutineBuildResponseBody {
	return s.Body
}

func (s *CreateRoutineBuildResponse) SetHeaders(v map[string]*string) *CreateRoutineBuildResponse {
	s.Headers = v
	return s
}

func (s *CreateRoutineBuildResponse) SetStatusCode(v int32) *CreateRoutineBuildResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRoutineBuildResponse) SetBody(v *CreateRoutineBuildResponseBody) *CreateRoutineBuildResponse {
	s.Body = v
	return s
}

func (s *CreateRoutineBuildResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
