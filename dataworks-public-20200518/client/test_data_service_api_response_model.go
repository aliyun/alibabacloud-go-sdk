// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTestDataServiceApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TestDataServiceApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TestDataServiceApiResponse
	GetStatusCode() *int32
	SetBody(v *TestDataServiceApiResponseBody) *TestDataServiceApiResponse
	GetBody() *TestDataServiceApiResponseBody
}

type TestDataServiceApiResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TestDataServiceApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TestDataServiceApiResponse) String() string {
	return dara.Prettify(s)
}

func (s TestDataServiceApiResponse) GoString() string {
	return s.String()
}

func (s *TestDataServiceApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TestDataServiceApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TestDataServiceApiResponse) GetBody() *TestDataServiceApiResponseBody {
	return s.Body
}

func (s *TestDataServiceApiResponse) SetHeaders(v map[string]*string) *TestDataServiceApiResponse {
	s.Headers = v
	return s
}

func (s *TestDataServiceApiResponse) SetStatusCode(v int32) *TestDataServiceApiResponse {
	s.StatusCode = &v
	return s
}

func (s *TestDataServiceApiResponse) SetBody(v *TestDataServiceApiResponseBody) *TestDataServiceApiResponse {
	s.Body = v
	return s
}

func (s *TestDataServiceApiResponse) Validate() error {
	return dara.Validate(s)
}
