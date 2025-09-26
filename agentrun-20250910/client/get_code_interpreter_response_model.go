// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCodeInterpreterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCodeInterpreterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCodeInterpreterResponse
	GetStatusCode() *int32
	SetBody(v *CodeInterpreterResult) *GetCodeInterpreterResponse
	GetBody() *CodeInterpreterResult
}

type GetCodeInterpreterResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CodeInterpreterResult `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCodeInterpreterResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCodeInterpreterResponse) GoString() string {
	return s.String()
}

func (s *GetCodeInterpreterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCodeInterpreterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCodeInterpreterResponse) GetBody() *CodeInterpreterResult {
	return s.Body
}

func (s *GetCodeInterpreterResponse) SetHeaders(v map[string]*string) *GetCodeInterpreterResponse {
	s.Headers = v
	return s
}

func (s *GetCodeInterpreterResponse) SetStatusCode(v int32) *GetCodeInterpreterResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCodeInterpreterResponse) SetBody(v *CodeInterpreterResult) *GetCodeInterpreterResponse {
	s.Body = v
	return s
}

func (s *GetCodeInterpreterResponse) Validate() error {
	return dara.Validate(s)
}
