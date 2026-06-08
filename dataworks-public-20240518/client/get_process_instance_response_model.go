// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProcessInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProcessInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProcessInstanceResponse
	GetStatusCode() *int32
	SetBody(v *GetProcessInstanceResponseBody) *GetProcessInstanceResponse
	GetBody() *GetProcessInstanceResponseBody
}

type GetProcessInstanceResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProcessInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProcessInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProcessInstanceResponse) GoString() string {
	return s.String()
}

func (s *GetProcessInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProcessInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProcessInstanceResponse) GetBody() *GetProcessInstanceResponseBody {
	return s.Body
}

func (s *GetProcessInstanceResponse) SetHeaders(v map[string]*string) *GetProcessInstanceResponse {
	s.Headers = v
	return s
}

func (s *GetProcessInstanceResponse) SetStatusCode(v int32) *GetProcessInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProcessInstanceResponse) SetBody(v *GetProcessInstanceResponseBody) *GetProcessInstanceResponse {
	s.Body = v
	return s
}

func (s *GetProcessInstanceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
