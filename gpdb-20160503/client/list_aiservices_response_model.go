// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAIServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAIServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAIServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListAIServicesResponseBody) *ListAIServicesResponse
	GetBody() *ListAIServicesResponseBody
}

type ListAIServicesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAIServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAIServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAIServicesResponse) GoString() string {
	return s.String()
}

func (s *ListAIServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAIServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAIServicesResponse) GetBody() *ListAIServicesResponseBody {
	return s.Body
}

func (s *ListAIServicesResponse) SetHeaders(v map[string]*string) *ListAIServicesResponse {
	s.Headers = v
	return s
}

func (s *ListAIServicesResponse) SetStatusCode(v int32) *ListAIServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAIServicesResponse) SetBody(v *ListAIServicesResponseBody) *ListAIServicesResponse {
	s.Body = v
	return s
}

func (s *ListAIServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
