// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConfigFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateConfigFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateConfigFileResponse
	GetStatusCode() *int32
	SetBody(v *CreateConfigFileResponseBody) *CreateConfigFileResponse
	GetBody() *CreateConfigFileResponseBody
}

type CreateConfigFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateConfigFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateConfigFileResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateConfigFileResponse) GoString() string {
	return s.String()
}

func (s *CreateConfigFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateConfigFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateConfigFileResponse) GetBody() *CreateConfigFileResponseBody {
	return s.Body
}

func (s *CreateConfigFileResponse) SetHeaders(v map[string]*string) *CreateConfigFileResponse {
	s.Headers = v
	return s
}

func (s *CreateConfigFileResponse) SetStatusCode(v int32) *CreateConfigFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateConfigFileResponse) SetBody(v *CreateConfigFileResponseBody) *CreateConfigFileResponse {
	s.Body = v
	return s
}

func (s *CreateConfigFileResponse) Validate() error {
	return dara.Validate(s)
}
