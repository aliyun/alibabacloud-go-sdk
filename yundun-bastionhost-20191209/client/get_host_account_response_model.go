// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHostAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHostAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHostAccountResponse
	GetStatusCode() *int32
	SetBody(v *GetHostAccountResponseBody) *GetHostAccountResponse
	GetBody() *GetHostAccountResponseBody
}

type GetHostAccountResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHostAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHostAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHostAccountResponse) GoString() string {
	return s.String()
}

func (s *GetHostAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHostAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHostAccountResponse) GetBody() *GetHostAccountResponseBody {
	return s.Body
}

func (s *GetHostAccountResponse) SetHeaders(v map[string]*string) *GetHostAccountResponse {
	s.Headers = v
	return s
}

func (s *GetHostAccountResponse) SetStatusCode(v int32) *GetHostAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHostAccountResponse) SetBody(v *GetHostAccountResponseBody) *GetHostAccountResponse {
	s.Body = v
	return s
}

func (s *GetHostAccountResponse) Validate() error {
	return dara.Validate(s)
}
