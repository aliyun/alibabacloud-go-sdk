// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitIsvCustomerTermsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitIsvCustomerTermsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitIsvCustomerTermsResponse
	GetStatusCode() *int32
	SetBody(v *SubmitIsvCustomerTermsResponseBody) *SubmitIsvCustomerTermsResponse
	GetBody() *SubmitIsvCustomerTermsResponseBody
}

type SubmitIsvCustomerTermsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitIsvCustomerTermsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitIsvCustomerTermsResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitIsvCustomerTermsResponse) GoString() string {
	return s.String()
}

func (s *SubmitIsvCustomerTermsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitIsvCustomerTermsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitIsvCustomerTermsResponse) GetBody() *SubmitIsvCustomerTermsResponseBody {
	return s.Body
}

func (s *SubmitIsvCustomerTermsResponse) SetHeaders(v map[string]*string) *SubmitIsvCustomerTermsResponse {
	s.Headers = v
	return s
}

func (s *SubmitIsvCustomerTermsResponse) SetStatusCode(v int32) *SubmitIsvCustomerTermsResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitIsvCustomerTermsResponse) SetBody(v *SubmitIsvCustomerTermsResponseBody) *SubmitIsvCustomerTermsResponse {
	s.Body = v
	return s
}

func (s *SubmitIsvCustomerTermsResponse) Validate() error {
	return dara.Validate(s)
}
