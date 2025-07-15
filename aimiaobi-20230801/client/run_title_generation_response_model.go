// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunTitleGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunTitleGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunTitleGenerationResponse
	GetStatusCode() *int32
	SetBody(v *RunTitleGenerationResponseBody) *RunTitleGenerationResponse
	GetBody() *RunTitleGenerationResponseBody
}

type RunTitleGenerationResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunTitleGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunTitleGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunTitleGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunTitleGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunTitleGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunTitleGenerationResponse) GetBody() *RunTitleGenerationResponseBody {
	return s.Body
}

func (s *RunTitleGenerationResponse) SetHeaders(v map[string]*string) *RunTitleGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunTitleGenerationResponse) SetStatusCode(v int32) *RunTitleGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunTitleGenerationResponse) SetBody(v *RunTitleGenerationResponseBody) *RunTitleGenerationResponse {
	s.Body = v
	return s
}

func (s *RunTitleGenerationResponse) Validate() error {
	return dara.Validate(s)
}
