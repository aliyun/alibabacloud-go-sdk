// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerSensitiveInfoPhysicalDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CustomerSensitiveInfoPhysicalDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CustomerSensitiveInfoPhysicalDeleteResponse
	GetStatusCode() *int32
	SetBody(v *CustomerSensitiveInfoPhysicalDeleteResponseBody) *CustomerSensitiveInfoPhysicalDeleteResponse
	GetBody() *CustomerSensitiveInfoPhysicalDeleteResponseBody
}

type CustomerSensitiveInfoPhysicalDeleteResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomerSensitiveInfoPhysicalDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CustomerSensitiveInfoPhysicalDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s CustomerSensitiveInfoPhysicalDeleteResponse) GoString() string {
	return s.String()
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponse) GetBody() *CustomerSensitiveInfoPhysicalDeleteResponseBody {
	return s.Body
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponse) SetHeaders(v map[string]*string) *CustomerSensitiveInfoPhysicalDeleteResponse {
	s.Headers = v
	return s
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponse) SetStatusCode(v int32) *CustomerSensitiveInfoPhysicalDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponse) SetBody(v *CustomerSensitiveInfoPhysicalDeleteResponseBody) *CustomerSensitiveInfoPhysicalDeleteResponse {
	s.Body = v
	return s
}

func (s *CustomerSensitiveInfoPhysicalDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
