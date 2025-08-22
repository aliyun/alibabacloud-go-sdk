// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateModuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateModuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateModuleResponseBody) *CreateModuleResponse
	GetBody() *CreateModuleResponseBody
}

type CreateModuleResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateModuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateModuleResponse) GoString() string {
	return s.String()
}

func (s *CreateModuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateModuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateModuleResponse) GetBody() *CreateModuleResponseBody {
	return s.Body
}

func (s *CreateModuleResponse) SetHeaders(v map[string]*string) *CreateModuleResponse {
	s.Headers = v
	return s
}

func (s *CreateModuleResponse) SetStatusCode(v int32) *CreateModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateModuleResponse) SetBody(v *CreateModuleResponseBody) *CreateModuleResponse {
	s.Body = v
	return s
}

func (s *CreateModuleResponse) Validate() error {
	return dara.Validate(s)
}
