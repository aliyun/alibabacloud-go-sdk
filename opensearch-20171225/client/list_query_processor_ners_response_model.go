// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQueryProcessorNersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListQueryProcessorNersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListQueryProcessorNersResponse
	GetStatusCode() *int32
	SetBody(v *ListQueryProcessorNersResponseBody) *ListQueryProcessorNersResponse
	GetBody() *ListQueryProcessorNersResponseBody
}

type ListQueryProcessorNersResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListQueryProcessorNersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListQueryProcessorNersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListQueryProcessorNersResponse) GoString() string {
	return s.String()
}

func (s *ListQueryProcessorNersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListQueryProcessorNersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListQueryProcessorNersResponse) GetBody() *ListQueryProcessorNersResponseBody {
	return s.Body
}

func (s *ListQueryProcessorNersResponse) SetHeaders(v map[string]*string) *ListQueryProcessorNersResponse {
	s.Headers = v
	return s
}

func (s *ListQueryProcessorNersResponse) SetStatusCode(v int32) *ListQueryProcessorNersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListQueryProcessorNersResponse) SetBody(v *ListQueryProcessorNersResponseBody) *ListQueryProcessorNersResponse {
	s.Body = v
	return s
}

func (s *ListQueryProcessorNersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
