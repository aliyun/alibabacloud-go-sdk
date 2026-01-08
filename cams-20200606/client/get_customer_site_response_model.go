// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomerSiteResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCustomerSiteResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCustomerSiteResponse
	GetStatusCode() *int32
	SetBody(v *GetCustomerSiteResponseBody) *GetCustomerSiteResponse
	GetBody() *GetCustomerSiteResponseBody
}

type GetCustomerSiteResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCustomerSiteResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCustomerSiteResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCustomerSiteResponse) GoString() string {
	return s.String()
}

func (s *GetCustomerSiteResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCustomerSiteResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCustomerSiteResponse) GetBody() *GetCustomerSiteResponseBody {
	return s.Body
}

func (s *GetCustomerSiteResponse) SetHeaders(v map[string]*string) *GetCustomerSiteResponse {
	s.Headers = v
	return s
}

func (s *GetCustomerSiteResponse) SetStatusCode(v int32) *GetCustomerSiteResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCustomerSiteResponse) SetBody(v *GetCustomerSiteResponseBody) *GetCustomerSiteResponse {
	s.Body = v
	return s
}

func (s *GetCustomerSiteResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
