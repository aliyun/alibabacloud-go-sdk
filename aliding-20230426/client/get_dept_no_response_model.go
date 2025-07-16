// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeptNoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeptNoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeptNoResponse
	GetStatusCode() *int32
	SetBody(v *GetDeptNoResponseBody) *GetDeptNoResponse
	GetBody() *GetDeptNoResponseBody
}

type GetDeptNoResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeptNoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeptNoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeptNoResponse) GoString() string {
	return s.String()
}

func (s *GetDeptNoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeptNoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeptNoResponse) GetBody() *GetDeptNoResponseBody {
	return s.Body
}

func (s *GetDeptNoResponse) SetHeaders(v map[string]*string) *GetDeptNoResponse {
	s.Headers = v
	return s
}

func (s *GetDeptNoResponse) SetStatusCode(v int32) *GetDeptNoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeptNoResponse) SetBody(v *GetDeptNoResponseBody) *GetDeptNoResponse {
	s.Body = v
	return s
}

func (s *GetDeptNoResponse) Validate() error {
	return dara.Validate(s)
}
