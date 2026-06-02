// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCreditPackageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCreditPackageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCreditPackageResponse
	GetStatusCode() *int32
	SetBody(v *CreateCreditPackageResponseBody) *CreateCreditPackageResponse
	GetBody() *CreateCreditPackageResponseBody
}

type CreateCreditPackageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCreditPackageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCreditPackageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCreditPackageResponse) GoString() string {
	return s.String()
}

func (s *CreateCreditPackageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCreditPackageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCreditPackageResponse) GetBody() *CreateCreditPackageResponseBody {
	return s.Body
}

func (s *CreateCreditPackageResponse) SetHeaders(v map[string]*string) *CreateCreditPackageResponse {
	s.Headers = v
	return s
}

func (s *CreateCreditPackageResponse) SetStatusCode(v int32) *CreateCreditPackageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCreditPackageResponse) SetBody(v *CreateCreditPackageResponseBody) *CreateCreditPackageResponse {
	s.Body = v
	return s
}

func (s *CreateCreditPackageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
