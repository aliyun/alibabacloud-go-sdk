// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListWarmPoolsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListWarmPoolsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListWarmPoolsResponse
	GetStatusCode() *int32
	SetBody(v *ListWarmPoolsResponseBody) *ListWarmPoolsResponse
	GetBody() *ListWarmPoolsResponseBody
}

type ListWarmPoolsResponse struct {
	Headers    map[string]*string         `json:"headers" xml:"headers"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListWarmPoolsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListWarmPoolsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListWarmPoolsResponse) GoString() string {
	return s.String()
}

func (s *ListWarmPoolsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListWarmPoolsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListWarmPoolsResponse) GetBody() *ListWarmPoolsResponseBody {
	return s.Body
}

func (s *ListWarmPoolsResponse) SetHeaders(v map[string]*string) *ListWarmPoolsResponse {
	s.Headers = v
	return s
}

func (s *ListWarmPoolsResponse) SetStatusCode(v int32) *ListWarmPoolsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListWarmPoolsResponse) SetBody(v *ListWarmPoolsResponseBody) *ListWarmPoolsResponse {
	s.Body = v
	return s
}

func (s *ListWarmPoolsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
