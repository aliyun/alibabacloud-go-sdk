// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLimitationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLimitationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLimitationsResponse
	GetStatusCode() *int32
	SetBody(v *ListLimitationsResponseBody) *ListLimitationsResponse
	GetBody() *ListLimitationsResponseBody
}

type ListLimitationsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLimitationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLimitationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLimitationsResponse) GoString() string {
	return s.String()
}

func (s *ListLimitationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLimitationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLimitationsResponse) GetBody() *ListLimitationsResponseBody {
	return s.Body
}

func (s *ListLimitationsResponse) SetHeaders(v map[string]*string) *ListLimitationsResponse {
	s.Headers = v
	return s
}

func (s *ListLimitationsResponse) SetStatusCode(v int32) *ListLimitationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLimitationsResponse) SetBody(v *ListLimitationsResponseBody) *ListLimitationsResponse {
	s.Body = v
	return s
}

func (s *ListLimitationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
