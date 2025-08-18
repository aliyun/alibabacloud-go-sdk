// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSiteLogDeliveryQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSiteLogDeliveryQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSiteLogDeliveryQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetSiteLogDeliveryQuotaResponseBody) *GetSiteLogDeliveryQuotaResponse
	GetBody() *GetSiteLogDeliveryQuotaResponseBody
}

type GetSiteLogDeliveryQuotaResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSiteLogDeliveryQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSiteLogDeliveryQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSiteLogDeliveryQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetSiteLogDeliveryQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSiteLogDeliveryQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSiteLogDeliveryQuotaResponse) GetBody() *GetSiteLogDeliveryQuotaResponseBody {
	return s.Body
}

func (s *GetSiteLogDeliveryQuotaResponse) SetHeaders(v map[string]*string) *GetSiteLogDeliveryQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetSiteLogDeliveryQuotaResponse) SetStatusCode(v int32) *GetSiteLogDeliveryQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSiteLogDeliveryQuotaResponse) SetBody(v *GetSiteLogDeliveryQuotaResponseBody) *GetSiteLogDeliveryQuotaResponse {
	s.Body = v
	return s
}

func (s *GetSiteLogDeliveryQuotaResponse) Validate() error {
	return dara.Validate(s)
}
