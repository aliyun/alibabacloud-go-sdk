// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApiMcpServerValidateHclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApiMcpServerValidateHclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApiMcpServerValidateHclResponse
	GetStatusCode() *int32
	SetBody(v *ApiMcpServerValidateHclResponseBody) *ApiMcpServerValidateHclResponse
	GetBody() *ApiMcpServerValidateHclResponseBody
}

type ApiMcpServerValidateHclResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApiMcpServerValidateHclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApiMcpServerValidateHclResponse) String() string {
	return dara.Prettify(s)
}

func (s ApiMcpServerValidateHclResponse) GoString() string {
	return s.String()
}

func (s *ApiMcpServerValidateHclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApiMcpServerValidateHclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApiMcpServerValidateHclResponse) GetBody() *ApiMcpServerValidateHclResponseBody {
	return s.Body
}

func (s *ApiMcpServerValidateHclResponse) SetHeaders(v map[string]*string) *ApiMcpServerValidateHclResponse {
	s.Headers = v
	return s
}

func (s *ApiMcpServerValidateHclResponse) SetStatusCode(v int32) *ApiMcpServerValidateHclResponse {
	s.StatusCode = &v
	return s
}

func (s *ApiMcpServerValidateHclResponse) SetBody(v *ApiMcpServerValidateHclResponseBody) *ApiMcpServerValidateHclResponse {
	s.Body = v
	return s
}

func (s *ApiMcpServerValidateHclResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
