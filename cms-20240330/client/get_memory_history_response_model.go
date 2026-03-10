// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMemoryHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMemoryHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMemoryHistoryResponse
	GetStatusCode() *int32
	SetBody(v *GetMemoryHistoryResponseBody) *GetMemoryHistoryResponse
	GetBody() *GetMemoryHistoryResponseBody
}

type GetMemoryHistoryResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMemoryHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMemoryHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMemoryHistoryResponse) GoString() string {
	return s.String()
}

func (s *GetMemoryHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMemoryHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMemoryHistoryResponse) GetBody() *GetMemoryHistoryResponseBody {
	return s.Body
}

func (s *GetMemoryHistoryResponse) SetHeaders(v map[string]*string) *GetMemoryHistoryResponse {
	s.Headers = v
	return s
}

func (s *GetMemoryHistoryResponse) SetStatusCode(v int32) *GetMemoryHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMemoryHistoryResponse) SetBody(v *GetMemoryHistoryResponseBody) *GetMemoryHistoryResponse {
	s.Body = v
	return s
}

func (s *GetMemoryHistoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
