// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeyPairsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListKeyPairsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListKeyPairsResponse
	GetStatusCode() *int32
	SetBody(v *ListKeyPairsResponseBody) *ListKeyPairsResponse
	GetBody() *ListKeyPairsResponseBody
}

type ListKeyPairsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListKeyPairsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListKeyPairsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListKeyPairsResponse) GoString() string {
	return s.String()
}

func (s *ListKeyPairsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListKeyPairsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListKeyPairsResponse) GetBody() *ListKeyPairsResponseBody {
	return s.Body
}

func (s *ListKeyPairsResponse) SetHeaders(v map[string]*string) *ListKeyPairsResponse {
	s.Headers = v
	return s
}

func (s *ListKeyPairsResponse) SetStatusCode(v int32) *ListKeyPairsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListKeyPairsResponse) SetBody(v *ListKeyPairsResponseBody) *ListKeyPairsResponse {
	s.Body = v
	return s
}

func (s *ListKeyPairsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
