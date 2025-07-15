// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsApplicationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlashSmsApplicationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlashSmsApplicationsResponse
	GetStatusCode() *int32
	SetBody(v *ListFlashSmsApplicationsResponseBody) *ListFlashSmsApplicationsResponse
	GetBody() *ListFlashSmsApplicationsResponseBody
}

type ListFlashSmsApplicationsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlashSmsApplicationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlashSmsApplicationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsApplicationsResponse) GoString() string {
	return s.String()
}

func (s *ListFlashSmsApplicationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlashSmsApplicationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlashSmsApplicationsResponse) GetBody() *ListFlashSmsApplicationsResponseBody {
	return s.Body
}

func (s *ListFlashSmsApplicationsResponse) SetHeaders(v map[string]*string) *ListFlashSmsApplicationsResponse {
	s.Headers = v
	return s
}

func (s *ListFlashSmsApplicationsResponse) SetStatusCode(v int32) *ListFlashSmsApplicationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlashSmsApplicationsResponse) SetBody(v *ListFlashSmsApplicationsResponseBody) *ListFlashSmsApplicationsResponse {
	s.Body = v
	return s
}

func (s *ListFlashSmsApplicationsResponse) Validate() error {
	return dara.Validate(s)
}
