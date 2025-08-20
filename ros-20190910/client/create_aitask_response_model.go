// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAITaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAITaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAITaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateAITaskResponseBody) *CreateAITaskResponse
	GetBody() *CreateAITaskResponseBody
}

type CreateAITaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAITaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAITaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAITaskResponse) GoString() string {
	return s.String()
}

func (s *CreateAITaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAITaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAITaskResponse) GetBody() *CreateAITaskResponseBody {
	return s.Body
}

func (s *CreateAITaskResponse) SetHeaders(v map[string]*string) *CreateAITaskResponse {
	s.Headers = v
	return s
}

func (s *CreateAITaskResponse) SetStatusCode(v int32) *CreateAITaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAITaskResponse) SetBody(v *CreateAITaskResponseBody) *CreateAITaskResponse {
	s.Body = v
	return s
}

func (s *CreateAITaskResponse) Validate() error {
	return dara.Validate(s)
}
