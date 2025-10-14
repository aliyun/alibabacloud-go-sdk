// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLogStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLogStoreResponse
	GetStatusCode() *int32
	SetBody(v *CreateLogStoreResponseBody) *CreateLogStoreResponse
	GetBody() *CreateLogStoreResponseBody
}

type CreateLogStoreResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLogStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLogStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLogStoreResponse) GoString() string {
	return s.String()
}

func (s *CreateLogStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLogStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLogStoreResponse) GetBody() *CreateLogStoreResponseBody {
	return s.Body
}

func (s *CreateLogStoreResponse) SetHeaders(v map[string]*string) *CreateLogStoreResponse {
	s.Headers = v
	return s
}

func (s *CreateLogStoreResponse) SetStatusCode(v int32) *CreateLogStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLogStoreResponse) SetBody(v *CreateLogStoreResponseBody) *CreateLogStoreResponse {
	s.Body = v
	return s
}

func (s *CreateLogStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
