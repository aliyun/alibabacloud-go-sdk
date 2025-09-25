// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunLibraryChatGenerationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunLibraryChatGenerationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunLibraryChatGenerationResponse
	GetStatusCode() *int32
	SetBody(v *RunLibraryChatGenerationResponseBody) *RunLibraryChatGenerationResponse
	GetBody() *RunLibraryChatGenerationResponseBody
}

type RunLibraryChatGenerationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunLibraryChatGenerationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunLibraryChatGenerationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunLibraryChatGenerationResponse) GoString() string {
	return s.String()
}

func (s *RunLibraryChatGenerationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunLibraryChatGenerationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunLibraryChatGenerationResponse) GetBody() *RunLibraryChatGenerationResponseBody {
	return s.Body
}

func (s *RunLibraryChatGenerationResponse) SetHeaders(v map[string]*string) *RunLibraryChatGenerationResponse {
	s.Headers = v
	return s
}

func (s *RunLibraryChatGenerationResponse) SetStatusCode(v int32) *RunLibraryChatGenerationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunLibraryChatGenerationResponse) SetBody(v *RunLibraryChatGenerationResponseBody) *RunLibraryChatGenerationResponse {
	s.Body = v
	return s
}

func (s *RunLibraryChatGenerationResponse) Validate() error {
	return dara.Validate(s)
}
