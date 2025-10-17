// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseAITaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseAITaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseAITaskResponse
	GetStatusCode() *int32
	SetBody(v *CloseAITaskResponseBody) *CloseAITaskResponse
	GetBody() *CloseAITaskResponseBody
}

type CloseAITaskResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseAITaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseAITaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseAITaskResponse) GoString() string {
	return s.String()
}

func (s *CloseAITaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseAITaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseAITaskResponse) GetBody() *CloseAITaskResponseBody {
	return s.Body
}

func (s *CloseAITaskResponse) SetHeaders(v map[string]*string) *CloseAITaskResponse {
	s.Headers = v
	return s
}

func (s *CloseAITaskResponse) SetStatusCode(v int32) *CloseAITaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseAITaskResponse) SetBody(v *CloseAITaskResponseBody) *CloseAITaskResponse {
	s.Body = v
	return s
}

func (s *CloseAITaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
