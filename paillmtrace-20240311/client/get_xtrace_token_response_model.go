// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetXtraceTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetXtraceTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetXtraceTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetXtraceTokenResponseBody) *GetXtraceTokenResponse
	GetBody() *GetXtraceTokenResponseBody
}

type GetXtraceTokenResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetXtraceTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetXtraceTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetXtraceTokenResponse) GoString() string {
	return s.String()
}

func (s *GetXtraceTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetXtraceTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetXtraceTokenResponse) GetBody() *GetXtraceTokenResponseBody {
	return s.Body
}

func (s *GetXtraceTokenResponse) SetHeaders(v map[string]*string) *GetXtraceTokenResponse {
	s.Headers = v
	return s
}

func (s *GetXtraceTokenResponse) SetStatusCode(v int32) *GetXtraceTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetXtraceTokenResponse) SetBody(v *GetXtraceTokenResponseBody) *GetXtraceTokenResponse {
	s.Body = v
	return s
}

func (s *GetXtraceTokenResponse) Validate() error {
	return dara.Validate(s)
}
