// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlashSmsProvidersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlashSmsProvidersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlashSmsProvidersResponse
	GetStatusCode() *int32
	SetBody(v *ListFlashSmsProvidersResponseBody) *ListFlashSmsProvidersResponse
	GetBody() *ListFlashSmsProvidersResponseBody
}

type ListFlashSmsProvidersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlashSmsProvidersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlashSmsProvidersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlashSmsProvidersResponse) GoString() string {
	return s.String()
}

func (s *ListFlashSmsProvidersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlashSmsProvidersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlashSmsProvidersResponse) GetBody() *ListFlashSmsProvidersResponseBody {
	return s.Body
}

func (s *ListFlashSmsProvidersResponse) SetHeaders(v map[string]*string) *ListFlashSmsProvidersResponse {
	s.Headers = v
	return s
}

func (s *ListFlashSmsProvidersResponse) SetStatusCode(v int32) *ListFlashSmsProvidersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlashSmsProvidersResponse) SetBody(v *ListFlashSmsProvidersResponseBody) *ListFlashSmsProvidersResponse {
	s.Body = v
	return s
}

func (s *ListFlashSmsProvidersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
