// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListKeysResponseBody) *ListKeysResponse
	GetBody() *ListKeysResponseBody
}

type ListKeysResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKeysResponse) GoString() string {
	return s.String()
}

func (s *ListKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKeysResponse) GetBody() *ListKeysResponseBody {
	return s.Body
}

func (s *ListKeysResponse) SetHeaders(v map[string]*string) *ListKeysResponse {
	s.Headers = v
	return s
}

func (s *ListKeysResponse) SetStatusCode(v int32) *ListKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKeysResponse) SetBody(v *ListKeysResponseBody) *ListKeysResponse {
	s.Body = v
	return s
}

func (s *ListKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
