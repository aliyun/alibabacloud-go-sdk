// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSlsLogStoreResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSlsLogStoreResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSlsLogStoreResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSlsLogStoreResponseBody) *UpdateSlsLogStoreResponse
	GetBody() *UpdateSlsLogStoreResponseBody
}

type UpdateSlsLogStoreResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSlsLogStoreResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSlsLogStoreResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSlsLogStoreResponse) GoString() string {
	return s.String()
}

func (s *UpdateSlsLogStoreResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSlsLogStoreResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSlsLogStoreResponse) GetBody() *UpdateSlsLogStoreResponseBody {
	return s.Body
}

func (s *UpdateSlsLogStoreResponse) SetHeaders(v map[string]*string) *UpdateSlsLogStoreResponse {
	s.Headers = v
	return s
}

func (s *UpdateSlsLogStoreResponse) SetStatusCode(v int32) *UpdateSlsLogStoreResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSlsLogStoreResponse) SetBody(v *UpdateSlsLogStoreResponseBody) *UpdateSlsLogStoreResponse {
	s.Body = v
	return s
}

func (s *UpdateSlsLogStoreResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
