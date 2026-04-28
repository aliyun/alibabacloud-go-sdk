// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDeltaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDeltaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDeltaResponse
	GetStatusCode() *int32
	SetBody(v *ListDeltaResponseBody) *ListDeltaResponse
	GetBody() *ListDeltaResponseBody
}

type ListDeltaResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeltaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeltaResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDeltaResponse) GoString() string {
	return s.String()
}

func (s *ListDeltaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDeltaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDeltaResponse) GetBody() *ListDeltaResponseBody {
	return s.Body
}

func (s *ListDeltaResponse) SetHeaders(v map[string]*string) *ListDeltaResponse {
	s.Headers = v
	return s
}

func (s *ListDeltaResponse) SetStatusCode(v int32) *ListDeltaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeltaResponse) SetBody(v *ListDeltaResponseBody) *ListDeltaResponse {
	s.Body = v
	return s
}

func (s *ListDeltaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
