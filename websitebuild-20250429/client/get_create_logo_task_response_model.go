// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCreateLogoTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCreateLogoTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCreateLogoTaskResponse
	GetStatusCode() *int32
	SetBody(v *GetCreateLogoTaskResponseBody) *GetCreateLogoTaskResponse
	GetBody() *GetCreateLogoTaskResponseBody
}

type GetCreateLogoTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCreateLogoTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCreateLogoTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCreateLogoTaskResponse) GoString() string {
	return s.String()
}

func (s *GetCreateLogoTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCreateLogoTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCreateLogoTaskResponse) GetBody() *GetCreateLogoTaskResponseBody {
	return s.Body
}

func (s *GetCreateLogoTaskResponse) SetHeaders(v map[string]*string) *GetCreateLogoTaskResponse {
	s.Headers = v
	return s
}

func (s *GetCreateLogoTaskResponse) SetStatusCode(v int32) *GetCreateLogoTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCreateLogoTaskResponse) SetBody(v *GetCreateLogoTaskResponseBody) *GetCreateLogoTaskResponse {
	s.Body = v
	return s
}

func (s *GetCreateLogoTaskResponse) Validate() error {
	return dara.Validate(s)
}
