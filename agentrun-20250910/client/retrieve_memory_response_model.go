// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetrieveMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetrieveMemoryResponse
	GetStatusCode() *int32
	SetBody(v *RetrieveMemoryResponseBody) *RetrieveMemoryResponse
	GetBody() *RetrieveMemoryResponseBody
}

type RetrieveMemoryResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetrieveMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetrieveMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s RetrieveMemoryResponse) GoString() string {
	return s.String()
}

func (s *RetrieveMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetrieveMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetrieveMemoryResponse) GetBody() *RetrieveMemoryResponseBody {
	return s.Body
}

func (s *RetrieveMemoryResponse) SetHeaders(v map[string]*string) *RetrieveMemoryResponse {
	s.Headers = v
	return s
}

func (s *RetrieveMemoryResponse) SetStatusCode(v int32) *RetrieveMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *RetrieveMemoryResponse) SetBody(v *RetrieveMemoryResponseBody) *RetrieveMemoryResponse {
	s.Body = v
	return s
}

func (s *RetrieveMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
