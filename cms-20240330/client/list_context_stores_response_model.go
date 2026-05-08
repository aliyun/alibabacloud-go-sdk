// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextStoresResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextStoresResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextStoresResponse
	GetStatusCode() *int32
	SetBody(v *ListContextStoresResponseBody) *ListContextStoresResponse
	GetBody() *ListContextStoresResponseBody
}

type ListContextStoresResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextStoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextStoresResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextStoresResponse) GoString() string {
	return s.String()
}

func (s *ListContextStoresResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextStoresResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextStoresResponse) GetBody() *ListContextStoresResponseBody {
	return s.Body
}

func (s *ListContextStoresResponse) SetHeaders(v map[string]*string) *ListContextStoresResponse {
	s.Headers = v
	return s
}

func (s *ListContextStoresResponse) SetStatusCode(v int32) *ListContextStoresResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextStoresResponse) SetBody(v *ListContextStoresResponseBody) *ListContextStoresResponse {
	s.Body = v
	return s
}

func (s *ListContextStoresResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
