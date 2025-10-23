// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendValidateFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendValidateFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendValidateFileResponse
	GetStatusCode() *int32
	SetBody(v *SendValidateFileResponseBody) *SendValidateFileResponse
	GetBody() *SendValidateFileResponseBody
}

type SendValidateFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendValidateFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendValidateFileResponse) String() string {
	return dara.Prettify(s)
}

func (s SendValidateFileResponse) GoString() string {
	return s.String()
}

func (s *SendValidateFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendValidateFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendValidateFileResponse) GetBody() *SendValidateFileResponseBody {
	return s.Body
}

func (s *SendValidateFileResponse) SetHeaders(v map[string]*string) *SendValidateFileResponse {
	s.Headers = v
	return s
}

func (s *SendValidateFileResponse) SetStatusCode(v int32) *SendValidateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SendValidateFileResponse) SetBody(v *SendValidateFileResponseBody) *SendValidateFileResponse {
	s.Body = v
	return s
}

func (s *SendValidateFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
