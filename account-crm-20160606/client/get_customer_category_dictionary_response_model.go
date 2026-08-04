// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerCategoryDictionaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomerCategoryDictionaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomerCategoryDictionaryResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomerCategoryDictionaryResponseBody) *GetCustomerCategoryDictionaryResponse
	GetBody() *GetCustomerCategoryDictionaryResponseBody
}

type GetCustomerCategoryDictionaryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerCategoryDictionaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerCategoryDictionaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerCategoryDictionaryResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerCategoryDictionaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomerCategoryDictionaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomerCategoryDictionaryResponse) GetBody() *GetCustomerCategoryDictionaryResponseBody {
	return s.Body
}

func (s *GetCustomerCategoryDictionaryResponse) SetHeaders(v map[string]*string) *GetCustomerCategoryDictionaryResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerCategoryDictionaryResponse) SetStatusCode(v int32) *GetCustomerCategoryDictionaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerCategoryDictionaryResponse) SetBody(v *GetCustomerCategoryDictionaryResponseBody) *GetCustomerCategoryDictionaryResponse {
	s.Body = v
	return s
}

func (s *GetCustomerCategoryDictionaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
