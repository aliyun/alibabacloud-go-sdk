// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsCodeLimitConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSmsCodeLimitConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSmsCodeLimitConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetSmsCodeLimitConfigResponseBody) *GetSmsCodeLimitConfigResponse
	GetBody() *GetSmsCodeLimitConfigResponseBody
}

type GetSmsCodeLimitConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSmsCodeLimitConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSmsCodeLimitConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSmsCodeLimitConfigResponse) GoString() string {
	return s.String()
}

func (s *GetSmsCodeLimitConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSmsCodeLimitConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSmsCodeLimitConfigResponse) GetBody() *GetSmsCodeLimitConfigResponseBody {
	return s.Body
}

func (s *GetSmsCodeLimitConfigResponse) SetHeaders(v map[string]*string) *GetSmsCodeLimitConfigResponse {
	s.Headers = v
	return s
}

func (s *GetSmsCodeLimitConfigResponse) SetStatusCode(v int32) *GetSmsCodeLimitConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSmsCodeLimitConfigResponse) SetBody(v *GetSmsCodeLimitConfigResponseBody) *GetSmsCodeLimitConfigResponse {
	s.Body = v
	return s
}

func (s *GetSmsCodeLimitConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
