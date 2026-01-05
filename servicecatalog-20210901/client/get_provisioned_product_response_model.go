// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProvisionedProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProvisionedProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProvisionedProductResponse
	GetStatusCode() *int32
	SetBody(v *GetProvisionedProductResponseBody) *GetProvisionedProductResponse
	GetBody() *GetProvisionedProductResponseBody
}

type GetProvisionedProductResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProvisionedProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProvisionedProductResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProvisionedProductResponse) GoString() string {
	return s.String()
}

func (s *GetProvisionedProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProvisionedProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProvisionedProductResponse) GetBody() *GetProvisionedProductResponseBody {
	return s.Body
}

func (s *GetProvisionedProductResponse) SetHeaders(v map[string]*string) *GetProvisionedProductResponse {
	s.Headers = v
	return s
}

func (s *GetProvisionedProductResponse) SetStatusCode(v int32) *GetProvisionedProductResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProvisionedProductResponse) SetBody(v *GetProvisionedProductResponseBody) *GetProvisionedProductResponse {
	s.Body = v
	return s
}

func (s *GetProvisionedProductResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
