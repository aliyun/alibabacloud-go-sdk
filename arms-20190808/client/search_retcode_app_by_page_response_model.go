// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchRetcodeAppByPageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchRetcodeAppByPageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchRetcodeAppByPageResponse
	GetStatusCode() *int32
	SetBody(v *SearchRetcodeAppByPageResponseBody) *SearchRetcodeAppByPageResponse
	GetBody() *SearchRetcodeAppByPageResponseBody
}

type SearchRetcodeAppByPageResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchRetcodeAppByPageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchRetcodeAppByPageResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchRetcodeAppByPageResponse) GoString() string {
	return s.String()
}

func (s *SearchRetcodeAppByPageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchRetcodeAppByPageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchRetcodeAppByPageResponse) GetBody() *SearchRetcodeAppByPageResponseBody {
	return s.Body
}

func (s *SearchRetcodeAppByPageResponse) SetHeaders(v map[string]*string) *SearchRetcodeAppByPageResponse {
	s.Headers = v
	return s
}

func (s *SearchRetcodeAppByPageResponse) SetStatusCode(v int32) *SearchRetcodeAppByPageResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchRetcodeAppByPageResponse) SetBody(v *SearchRetcodeAppByPageResponseBody) *SearchRetcodeAppByPageResponse {
	s.Body = v
	return s
}

func (s *SearchRetcodeAppByPageResponse) Validate() error {
	return dara.Validate(s)
}
