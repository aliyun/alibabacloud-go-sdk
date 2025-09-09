// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceTestTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServiceTestTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServiceTestTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetServiceTestTaskResponseBody) *GetServiceTestTaskResponse
	GetBody() *GetServiceTestTaskResponseBody
}

type GetServiceTestTaskResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServiceTestTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServiceTestTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServiceTestTaskResponse) GoString() string {
	return s.String()
}

func (s *GetServiceTestTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServiceTestTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServiceTestTaskResponse) GetBody() *GetServiceTestTaskResponseBody {
	return s.Body
}

func (s *GetServiceTestTaskResponse) SetHeaders(v map[string]*string) *GetServiceTestTaskResponse {
	s.Headers = v
	return s
}

func (s *GetServiceTestTaskResponse) SetStatusCode(v int32) *GetServiceTestTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServiceTestTaskResponse) SetBody(v *GetServiceTestTaskResponseBody) *GetServiceTestTaskResponse {
	s.Body = v
	return s
}

func (s *GetServiceTestTaskResponse) Validate() error {
	return dara.Validate(s)
}
