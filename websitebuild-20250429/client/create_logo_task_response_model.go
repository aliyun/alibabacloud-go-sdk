// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLogoTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLogoTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLogoTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateLogoTaskResponseBody) *CreateLogoTaskResponse
	GetBody() *CreateLogoTaskResponseBody
}

type CreateLogoTaskResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLogoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLogoTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLogoTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateLogoTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLogoTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLogoTaskResponse) GetBody() *CreateLogoTaskResponseBody {
	return s.Body
}

func (s *CreateLogoTaskResponse) SetHeaders(v map[string]*string) *CreateLogoTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateLogoTaskResponse) SetStatusCode(v int32) *CreateLogoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLogoTaskResponse) SetBody(v *CreateLogoTaskResponseBody) *CreateLogoTaskResponse {
	s.Body = v
	return s
}

func (s *CreateLogoTaskResponse) Validate() error {
	return dara.Validate(s)
}
