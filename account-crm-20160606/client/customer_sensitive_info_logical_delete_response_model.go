// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCustomerSensitiveInfoLogicalDeleteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CustomerSensitiveInfoLogicalDeleteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CustomerSensitiveInfoLogicalDeleteResponse
	GetStatusCode() *int32
	SetBody(v *CustomerSensitiveInfoLogicalDeleteResponseBody) *CustomerSensitiveInfoLogicalDeleteResponse
	GetBody() *CustomerSensitiveInfoLogicalDeleteResponseBody
}

type CustomerSensitiveInfoLogicalDeleteResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CustomerSensitiveInfoLogicalDeleteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CustomerSensitiveInfoLogicalDeleteResponse) String() string {
	return dara.Prettify(s)
}

func (s CustomerSensitiveInfoLogicalDeleteResponse) GoString() string {
	return s.String()
}

func (s *CustomerSensitiveInfoLogicalDeleteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CustomerSensitiveInfoLogicalDeleteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CustomerSensitiveInfoLogicalDeleteResponse) GetBody() *CustomerSensitiveInfoLogicalDeleteResponseBody {
	return s.Body
}

func (s *CustomerSensitiveInfoLogicalDeleteResponse) SetHeaders(v map[string]*string) *CustomerSensitiveInfoLogicalDeleteResponse {
	s.Headers = v
	return s
}

func (s *CustomerSensitiveInfoLogicalDeleteResponse) SetStatusCode(v int32) *CustomerSensitiveInfoLogicalDeleteResponse {
	s.StatusCode = &v
	return s
}

func (s *CustomerSensitiveInfoLogicalDeleteResponse) SetBody(v *CustomerSensitiveInfoLogicalDeleteResponseBody) *CustomerSensitiveInfoLogicalDeleteResponse {
	s.Body = v
	return s
}

func (s *CustomerSensitiveInfoLogicalDeleteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
