// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceRenewPriceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetResourceRenewPriceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetResourceRenewPriceResponse
	GetStatusCode() *int32
	SetBody(v *GetResourceRenewPriceResponseBody) *GetResourceRenewPriceResponse
	GetBody() *GetResourceRenewPriceResponseBody
}

type GetResourceRenewPriceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetResourceRenewPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetResourceRenewPriceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetResourceRenewPriceResponse) GoString() string {
	return s.String()
}

func (s *GetResourceRenewPriceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetResourceRenewPriceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetResourceRenewPriceResponse) GetBody() *GetResourceRenewPriceResponseBody {
	return s.Body
}

func (s *GetResourceRenewPriceResponse) SetHeaders(v map[string]*string) *GetResourceRenewPriceResponse {
	s.Headers = v
	return s
}

func (s *GetResourceRenewPriceResponse) SetStatusCode(v int32) *GetResourceRenewPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetResourceRenewPriceResponse) SetBody(v *GetResourceRenewPriceResponseBody) *GetResourceRenewPriceResponse {
	s.Body = v
	return s
}

func (s *GetResourceRenewPriceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
