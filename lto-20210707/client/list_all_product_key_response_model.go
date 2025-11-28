// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllProductKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllProductKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllProductKeyResponse
	GetStatusCode() *int32
	SetBody(v *ListAllProductKeyResponseBody) *ListAllProductKeyResponse
	GetBody() *ListAllProductKeyResponseBody
}

type ListAllProductKeyResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllProductKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllProductKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllProductKeyResponse) GoString() string {
	return s.String()
}

func (s *ListAllProductKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllProductKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllProductKeyResponse) GetBody() *ListAllProductKeyResponseBody {
	return s.Body
}

func (s *ListAllProductKeyResponse) SetHeaders(v map[string]*string) *ListAllProductKeyResponse {
	s.Headers = v
	return s
}

func (s *ListAllProductKeyResponse) SetStatusCode(v int32) *ListAllProductKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllProductKeyResponse) SetBody(v *ListAllProductKeyResponseBody) *ListAllProductKeyResponse {
	s.Body = v
	return s
}

func (s *ListAllProductKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
