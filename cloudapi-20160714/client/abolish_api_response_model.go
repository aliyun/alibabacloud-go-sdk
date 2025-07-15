// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbolishApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AbolishApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AbolishApiResponse
	GetStatusCode() *int32
	SetBody(v *AbolishApiResponseBody) *AbolishApiResponse
	GetBody() *AbolishApiResponseBody
}

type AbolishApiResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbolishApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbolishApiResponse) String() string {
	return dara.Prettify(s)
}

func (s AbolishApiResponse) GoString() string {
	return s.String()
}

func (s *AbolishApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AbolishApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AbolishApiResponse) GetBody() *AbolishApiResponseBody {
	return s.Body
}

func (s *AbolishApiResponse) SetHeaders(v map[string]*string) *AbolishApiResponse {
	s.Headers = v
	return s
}

func (s *AbolishApiResponse) SetStatusCode(v int32) *AbolishApiResponse {
	s.StatusCode = &v
	return s
}

func (s *AbolishApiResponse) SetBody(v *AbolishApiResponseBody) *AbolishApiResponse {
	s.Body = v
	return s
}

func (s *AbolishApiResponse) Validate() error {
	return dara.Validate(s)
}
