// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunSearchGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunSearchGenerationResponse
	GetStatusCode() *int32
	SetBody(v *RunSearchGenerationResponseBody) *RunSearchGenerationResponse
	GetBody() *RunSearchGenerationResponseBody
}

type RunSearchGenerationResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSearchGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSearchGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunSearchGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunSearchGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunSearchGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunSearchGenerationResponse) GetBody() *RunSearchGenerationResponseBody {
	return s.Body
}

func (s *RunSearchGenerationResponse) SetHeaders(v map[string]*string) *RunSearchGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunSearchGenerationResponse) SetStatusCode(v int32) *RunSearchGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSearchGenerationResponse) SetBody(v *RunSearchGenerationResponseBody) *RunSearchGenerationResponse {
	s.Body = v
	return s
}

func (s *RunSearchGenerationResponse) Validate() error {
	return dara.Validate(s)
}
