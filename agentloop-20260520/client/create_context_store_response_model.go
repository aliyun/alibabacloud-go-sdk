// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContextStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContextStoreResponse
	GetStatusCode() *int32
	SetBody(v *CreateContextStoreResponseBody) *CreateContextStoreResponse
	GetBody() *CreateContextStoreResponseBody
}

type CreateContextStoreResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContextStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContextStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContextStoreResponse) GoString() string {
	return s.String()
}

func (s *CreateContextStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContextStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContextStoreResponse) GetBody() *CreateContextStoreResponseBody {
	return s.Body
}

func (s *CreateContextStoreResponse) SetHeaders(v map[string]*string) *CreateContextStoreResponse {
	s.Headers = v
	return s
}

func (s *CreateContextStoreResponse) SetStatusCode(v int32) *CreateContextStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContextStoreResponse) SetBody(v *CreateContextStoreResponseBody) *CreateContextStoreResponse {
	s.Body = v
	return s
}

func (s *CreateContextStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
