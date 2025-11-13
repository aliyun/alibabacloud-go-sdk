// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunContractResultGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunContractResultGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunContractResultGenerationResponse
	GetStatusCode() *int32
	SetBody(v *RunContractResultGenerationResponseBody) *RunContractResultGenerationResponse
	GetBody() *RunContractResultGenerationResponseBody
}

type RunContractResultGenerationResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunContractResultGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunContractResultGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunContractResultGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunContractResultGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunContractResultGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunContractResultGenerationResponse) GetBody() *RunContractResultGenerationResponseBody {
	return s.Body
}

func (s *RunContractResultGenerationResponse) SetHeaders(v map[string]*string) *RunContractResultGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunContractResultGenerationResponse) SetStatusCode(v int32) *RunContractResultGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunContractResultGenerationResponse) SetBody(v *RunContractResultGenerationResponseBody) *RunContractResultGenerationResponse {
	s.Body = v
	return s
}

func (s *RunContractResultGenerationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
