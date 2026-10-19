// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceManualMemoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReplaceManualMemoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReplaceManualMemoryResponse
	GetStatusCode() *int32
	SetBody(v *ReplaceManualMemoryResponseBody) *ReplaceManualMemoryResponse
	GetBody() *ReplaceManualMemoryResponseBody
}

type ReplaceManualMemoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReplaceManualMemoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReplaceManualMemoryResponse) String() string {
	return dara.Prettify(s)
}

func (s ReplaceManualMemoryResponse) GoString() string {
	return s.String()
}

func (s *ReplaceManualMemoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReplaceManualMemoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReplaceManualMemoryResponse) GetBody() *ReplaceManualMemoryResponseBody {
	return s.Body
}

func (s *ReplaceManualMemoryResponse) SetHeaders(v map[string]*string) *ReplaceManualMemoryResponse {
	s.Headers = v
	return s
}

func (s *ReplaceManualMemoryResponse) SetStatusCode(v int32) *ReplaceManualMemoryResponse {
	s.StatusCode = &v
	return s
}

func (s *ReplaceManualMemoryResponse) SetBody(v *ReplaceManualMemoryResponseBody) *ReplaceManualMemoryResponse {
	s.Body = v
	return s
}

func (s *ReplaceManualMemoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
