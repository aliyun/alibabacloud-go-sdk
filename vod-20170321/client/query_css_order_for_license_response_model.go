// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCssOrderForLicenseResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCssOrderForLicenseResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCssOrderForLicenseResponse
	GetStatusCode() *int32
	SetBody(v *QueryCssOrderForLicenseResponseBody) *QueryCssOrderForLicenseResponse
	GetBody() *QueryCssOrderForLicenseResponseBody
}

type QueryCssOrderForLicenseResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCssOrderForLicenseResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCssOrderForLicenseResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCssOrderForLicenseResponse) GoString() string {
	return s.String()
}

func (s *QueryCssOrderForLicenseResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCssOrderForLicenseResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCssOrderForLicenseResponse) GetBody() *QueryCssOrderForLicenseResponseBody {
	return s.Body
}

func (s *QueryCssOrderForLicenseResponse) SetHeaders(v map[string]*string) *QueryCssOrderForLicenseResponse {
	s.Headers = v
	return s
}

func (s *QueryCssOrderForLicenseResponse) SetStatusCode(v int32) *QueryCssOrderForLicenseResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCssOrderForLicenseResponse) SetBody(v *QueryCssOrderForLicenseResponseBody) *QueryCssOrderForLicenseResponse {
	s.Body = v
	return s
}

func (s *QueryCssOrderForLicenseResponse) Validate() error {
	return dara.Validate(s)
}
