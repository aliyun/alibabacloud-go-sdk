// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAutoMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ClearAutoMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ClearAutoMemoryResponse
	GetStatusCode() *int32
	SetBody(v *ClearAutoMemoryResponseBody) *ClearAutoMemoryResponse
	GetBody() *ClearAutoMemoryResponseBody
}

type ClearAutoMemoryResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearAutoMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearAutoMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ClearAutoMemoryResponse) GoString() string {
	return s.String()
}

func (s *ClearAutoMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ClearAutoMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ClearAutoMemoryResponse) GetBody() *ClearAutoMemoryResponseBody {
	return s.Body
}

func (s *ClearAutoMemoryResponse) SetHeaders(v map[string]*string) *ClearAutoMemoryResponse {
	s.Headers = v
	return s
}

func (s *ClearAutoMemoryResponse) SetStatusCode(v int32) *ClearAutoMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearAutoMemoryResponse) SetBody(v *ClearAutoMemoryResponseBody) *ClearAutoMemoryResponse {
	s.Body = v
	return s
}

func (s *ClearAutoMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
