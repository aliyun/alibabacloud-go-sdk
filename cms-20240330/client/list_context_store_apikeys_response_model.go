// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContextStoreAPIKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListContextStoreAPIKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListContextStoreAPIKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListContextStoreAPIKeysResponseBody) *ListContextStoreAPIKeysResponse
	GetBody() *ListContextStoreAPIKeysResponseBody
}

type ListContextStoreAPIKeysResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListContextStoreAPIKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListContextStoreAPIKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListContextStoreAPIKeysResponse) GoString() string {
	return s.String()
}

func (s *ListContextStoreAPIKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListContextStoreAPIKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListContextStoreAPIKeysResponse) GetBody() *ListContextStoreAPIKeysResponseBody {
	return s.Body
}

func (s *ListContextStoreAPIKeysResponse) SetHeaders(v map[string]*string) *ListContextStoreAPIKeysResponse {
	s.Headers = v
	return s
}

func (s *ListContextStoreAPIKeysResponse) SetStatusCode(v int32) *ListContextStoreAPIKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListContextStoreAPIKeysResponse) SetBody(v *ListContextStoreAPIKeysResponseBody) *ListContextStoreAPIKeysResponse {
	s.Body = v
	return s
}

func (s *ListContextStoreAPIKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
