// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCodeInterpreterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCodeInterpreterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCodeInterpreterResponse
	GetStatusCode() *int32
	SetBody(v *CodeInterpreterResult) *CreateCodeInterpreterResponse
	GetBody() *CodeInterpreterResult
}

type CreateCodeInterpreterResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CodeInterpreterResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCodeInterpreterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCodeInterpreterResponse) GoString() string {
	return s.String()
}

func (s *CreateCodeInterpreterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCodeInterpreterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCodeInterpreterResponse) GetBody() *CodeInterpreterResult {
	return s.Body
}

func (s *CreateCodeInterpreterResponse) SetHeaders(v map[string]*string) *CreateCodeInterpreterResponse {
	s.Headers = v
	return s
}

func (s *CreateCodeInterpreterResponse) SetStatusCode(v int32) *CreateCodeInterpreterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCodeInterpreterResponse) SetBody(v *CodeInterpreterResult) *CreateCodeInterpreterResponse {
	s.Body = v
	return s
}

func (s *CreateCodeInterpreterResponse) Validate() error {
	return dara.Validate(s)
}
