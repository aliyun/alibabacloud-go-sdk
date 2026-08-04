// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerInformationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomerInformationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomerInformationResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomerInformationResponseBody) *GetCustomerInformationResponse
	GetBody() *GetCustomerInformationResponseBody
}

type GetCustomerInformationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerInformationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerInformationResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerInformationResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerInformationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomerInformationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomerInformationResponse) GetBody() *GetCustomerInformationResponseBody {
	return s.Body
}

func (s *GetCustomerInformationResponse) SetHeaders(v map[string]*string) *GetCustomerInformationResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerInformationResponse) SetStatusCode(v int32) *GetCustomerInformationResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerInformationResponse) SetBody(v *GetCustomerInformationResponseBody) *GetCustomerInformationResponse {
	s.Body = v
	return s
}

func (s *GetCustomerInformationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
