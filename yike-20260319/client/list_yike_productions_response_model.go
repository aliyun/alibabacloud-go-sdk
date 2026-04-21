// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListYikeProductionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListYikeProductionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListYikeProductionsResponse
	GetStatusCode() *int32
	SetBody(v *ListYikeProductionsResponseBody) *ListYikeProductionsResponse
	GetBody() *ListYikeProductionsResponseBody
}

type ListYikeProductionsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListYikeProductionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListYikeProductionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListYikeProductionsResponse) GoString() string {
	return s.String()
}

func (s *ListYikeProductionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListYikeProductionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListYikeProductionsResponse) GetBody() *ListYikeProductionsResponseBody {
	return s.Body
}

func (s *ListYikeProductionsResponse) SetHeaders(v map[string]*string) *ListYikeProductionsResponse {
	s.Headers = v
	return s
}

func (s *ListYikeProductionsResponse) SetStatusCode(v int32) *ListYikeProductionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListYikeProductionsResponse) SetBody(v *ListYikeProductionsResponseBody) *ListYikeProductionsResponse {
	s.Body = v
	return s
}

func (s *ListYikeProductionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
