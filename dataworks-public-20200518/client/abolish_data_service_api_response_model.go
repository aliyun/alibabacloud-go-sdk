// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbolishDataServiceApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AbolishDataServiceApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AbolishDataServiceApiResponse
	GetStatusCode() *int32
	SetBody(v *AbolishDataServiceApiResponseBody) *AbolishDataServiceApiResponse
	GetBody() *AbolishDataServiceApiResponseBody
}

type AbolishDataServiceApiResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbolishDataServiceApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbolishDataServiceApiResponse) String() string {
	return dara.Prettify(s)
}

func (s AbolishDataServiceApiResponse) GoString() string {
	return s.String()
}

func (s *AbolishDataServiceApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AbolishDataServiceApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AbolishDataServiceApiResponse) GetBody() *AbolishDataServiceApiResponseBody {
	return s.Body
}

func (s *AbolishDataServiceApiResponse) SetHeaders(v map[string]*string) *AbolishDataServiceApiResponse {
	s.Headers = v
	return s
}

func (s *AbolishDataServiceApiResponse) SetStatusCode(v int32) *AbolishDataServiceApiResponse {
	s.StatusCode = &v
	return s
}

func (s *AbolishDataServiceApiResponse) SetBody(v *AbolishDataServiceApiResponseBody) *AbolishDataServiceApiResponse {
	s.Body = v
	return s
}

func (s *AbolishDataServiceApiResponse) Validate() error {
	return dara.Validate(s)
}
