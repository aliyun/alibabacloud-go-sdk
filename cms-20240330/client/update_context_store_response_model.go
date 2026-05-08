// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContextStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContextStoreResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContextStoreResponseBody) *UpdateContextStoreResponse
	GetBody() *UpdateContextStoreResponseBody
}

type UpdateContextStoreResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContextStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContextStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextStoreResponse) GoString() string {
	return s.String()
}

func (s *UpdateContextStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContextStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContextStoreResponse) GetBody() *UpdateContextStoreResponseBody {
	return s.Body
}

func (s *UpdateContextStoreResponse) SetHeaders(v map[string]*string) *UpdateContextStoreResponse {
	s.Headers = v
	return s
}

func (s *UpdateContextStoreResponse) SetStatusCode(v int32) *UpdateContextStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContextStoreResponse) SetBody(v *UpdateContextStoreResponseBody) *UpdateContextStoreResponse {
	s.Body = v
	return s
}

func (s *UpdateContextStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
