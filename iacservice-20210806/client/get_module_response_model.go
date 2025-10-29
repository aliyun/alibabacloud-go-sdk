// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetModuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetModuleResponse
	GetStatusCode() *int32
	SetBody(v *GetModuleResponseBody) *GetModuleResponse
	GetBody() *GetModuleResponseBody
}

type GetModuleResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetModuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetModuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetModuleResponse) GoString() string {
	return s.String()
}

func (s *GetModuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetModuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetModuleResponse) GetBody() *GetModuleResponseBody {
	return s.Body
}

func (s *GetModuleResponse) SetHeaders(v map[string]*string) *GetModuleResponse {
	s.Headers = v
	return s
}

func (s *GetModuleResponse) SetStatusCode(v int32) *GetModuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetModuleResponse) SetBody(v *GetModuleResponseBody) *GetModuleResponse {
	s.Body = v
	return s
}

func (s *GetModuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
