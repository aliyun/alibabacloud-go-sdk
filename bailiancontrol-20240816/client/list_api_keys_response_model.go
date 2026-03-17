// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApiKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApiKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApiKeysResponse
	GetStatusCode() *int32
	SetBody(v *ListApiKeysResponseBody) *ListApiKeysResponse
	GetBody() *ListApiKeysResponseBody
}

type ListApiKeysResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApiKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApiKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApiKeysResponse) GoString() string {
	return s.String()
}

func (s *ListApiKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApiKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApiKeysResponse) GetBody() *ListApiKeysResponseBody {
	return s.Body
}

func (s *ListApiKeysResponse) SetHeaders(v map[string]*string) *ListApiKeysResponse {
	s.Headers = v
	return s
}

func (s *ListApiKeysResponse) SetStatusCode(v int32) *ListApiKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApiKeysResponse) SetBody(v *ListApiKeysResponseBody) *ListApiKeysResponse {
	s.Body = v
	return s
}

func (s *ListApiKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
