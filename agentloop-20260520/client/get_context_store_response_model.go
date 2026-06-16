// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetContextStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetContextStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetContextStoreResponse
	GetStatusCode() *int32
	SetBody(v *GetContextStoreResponseBody) *GetContextStoreResponse
	GetBody() *GetContextStoreResponseBody
}

type GetContextStoreResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetContextStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetContextStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s GetContextStoreResponse) GoString() string {
	return s.String()
}

func (s *GetContextStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetContextStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetContextStoreResponse) GetBody() *GetContextStoreResponseBody {
	return s.Body
}

func (s *GetContextStoreResponse) SetHeaders(v map[string]*string) *GetContextStoreResponse {
	s.Headers = v
	return s
}

func (s *GetContextStoreResponse) SetStatusCode(v int32) *GetContextStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *GetContextStoreResponse) SetBody(v *GetContextStoreResponseBody) *GetContextStoreResponse {
	s.Body = v
	return s
}

func (s *GetContextStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
