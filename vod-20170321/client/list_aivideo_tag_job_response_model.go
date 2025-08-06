// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIVideoTagJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIVideoTagJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIVideoTagJobResponse
	GetStatusCode() *int32
	SetBody(v *ListAIVideoTagJobResponseBody) *ListAIVideoTagJobResponse
	GetBody() *ListAIVideoTagJobResponseBody
}

type ListAIVideoTagJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIVideoTagJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIVideoTagJobResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIVideoTagJobResponse) GoString() string {
	return s.String()
}

func (s *ListAIVideoTagJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIVideoTagJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIVideoTagJobResponse) GetBody() *ListAIVideoTagJobResponseBody {
	return s.Body
}

func (s *ListAIVideoTagJobResponse) SetHeaders(v map[string]*string) *ListAIVideoTagJobResponse {
	s.Headers = v
	return s
}

func (s *ListAIVideoTagJobResponse) SetStatusCode(v int32) *ListAIVideoTagJobResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIVideoTagJobResponse) SetBody(v *ListAIVideoTagJobResponseBody) *ListAIVideoTagJobResponse {
	s.Body = v
	return s
}

func (s *ListAIVideoTagJobResponse) Validate() error {
	return dara.Validate(s)
}
