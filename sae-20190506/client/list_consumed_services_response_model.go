// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumedServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConsumedServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConsumedServicesResponse
	GetStatusCode() *int32
	SetBody(v *ListConsumedServicesResponseBody) *ListConsumedServicesResponse
	GetBody() *ListConsumedServicesResponseBody
}

type ListConsumedServicesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConsumedServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConsumedServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConsumedServicesResponse) GoString() string {
	return s.String()
}

func (s *ListConsumedServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConsumedServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConsumedServicesResponse) GetBody() *ListConsumedServicesResponseBody {
	return s.Body
}

func (s *ListConsumedServicesResponse) SetHeaders(v map[string]*string) *ListConsumedServicesResponse {
	s.Headers = v
	return s
}

func (s *ListConsumedServicesResponse) SetStatusCode(v int32) *ListConsumedServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConsumedServicesResponse) SetBody(v *ListConsumedServicesResponseBody) *ListConsumedServicesResponse {
	s.Body = v
	return s
}

func (s *ListConsumedServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
