// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipCurrentStepResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SkipCurrentStepResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SkipCurrentStepResponse
	GetStatusCode() *int32
	SetBody(v *SkipCurrentStepResponseBody) *SkipCurrentStepResponse
	GetBody() *SkipCurrentStepResponseBody
}

type SkipCurrentStepResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SkipCurrentStepResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SkipCurrentStepResponse) String() string {
	return dara.Prettify(s)
}

func (s SkipCurrentStepResponse) GoString() string {
	return s.String()
}

func (s *SkipCurrentStepResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SkipCurrentStepResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SkipCurrentStepResponse) GetBody() *SkipCurrentStepResponseBody {
	return s.Body
}

func (s *SkipCurrentStepResponse) SetHeaders(v map[string]*string) *SkipCurrentStepResponse {
	s.Headers = v
	return s
}

func (s *SkipCurrentStepResponse) SetStatusCode(v int32) *SkipCurrentStepResponse {
	s.StatusCode = &v
	return s
}

func (s *SkipCurrentStepResponse) SetBody(v *SkipCurrentStepResponseBody) *SkipCurrentStepResponse {
	s.Body = v
	return s
}

func (s *SkipCurrentStepResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
