// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInterceptionTargetPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInterceptionTargetPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInterceptionTargetPageResponse
	GetStatusCode() *int32
	SetBody(v *ListInterceptionTargetPageResponseBody) *ListInterceptionTargetPageResponse
	GetBody() *ListInterceptionTargetPageResponseBody
}

type ListInterceptionTargetPageResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInterceptionTargetPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInterceptionTargetPageResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInterceptionTargetPageResponse) GoString() string {
	return s.String()
}

func (s *ListInterceptionTargetPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInterceptionTargetPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInterceptionTargetPageResponse) GetBody() *ListInterceptionTargetPageResponseBody {
	return s.Body
}

func (s *ListInterceptionTargetPageResponse) SetHeaders(v map[string]*string) *ListInterceptionTargetPageResponse {
	s.Headers = v
	return s
}

func (s *ListInterceptionTargetPageResponse) SetStatusCode(v int32) *ListInterceptionTargetPageResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInterceptionTargetPageResponse) SetBody(v *ListInterceptionTargetPageResponseBody) *ListInterceptionTargetPageResponse {
	s.Body = v
	return s
}

func (s *ListInterceptionTargetPageResponse) Validate() error {
	return dara.Validate(s)
}
