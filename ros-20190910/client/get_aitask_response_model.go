// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAITaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAITaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAITaskResponse
	GetStatusCode() *int32
	SetBody(v *GetAITaskResponseBody) *GetAITaskResponse
	GetBody() *GetAITaskResponseBody
}

type GetAITaskResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAITaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAITaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAITaskResponse) GoString() string {
	return s.String()
}

func (s *GetAITaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAITaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAITaskResponse) GetBody() *GetAITaskResponseBody {
	return s.Body
}

func (s *GetAITaskResponse) SetHeaders(v map[string]*string) *GetAITaskResponse {
	s.Headers = v
	return s
}

func (s *GetAITaskResponse) SetStatusCode(v int32) *GetAITaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAITaskResponse) SetBody(v *GetAITaskResponseBody) *GetAITaskResponse {
	s.Body = v
	return s
}

func (s *GetAITaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
