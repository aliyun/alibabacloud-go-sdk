// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMcubeMiniTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMcubeMiniTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMcubeMiniTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateMcubeMiniTaskResponseBody) *CreateMcubeMiniTaskResponse
	GetBody() *CreateMcubeMiniTaskResponseBody
}

type CreateMcubeMiniTaskResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMcubeMiniTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMcubeMiniTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMcubeMiniTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateMcubeMiniTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMcubeMiniTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMcubeMiniTaskResponse) GetBody() *CreateMcubeMiniTaskResponseBody {
	return s.Body
}

func (s *CreateMcubeMiniTaskResponse) SetHeaders(v map[string]*string) *CreateMcubeMiniTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateMcubeMiniTaskResponse) SetStatusCode(v int32) *CreateMcubeMiniTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMcubeMiniTaskResponse) SetBody(v *CreateMcubeMiniTaskResponseBody) *CreateMcubeMiniTaskResponse {
	s.Body = v
	return s
}

func (s *CreateMcubeMiniTaskResponse) Validate() error {
	return dara.Validate(s)
}
