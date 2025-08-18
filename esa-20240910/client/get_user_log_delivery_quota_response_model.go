// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserLogDeliveryQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserLogDeliveryQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserLogDeliveryQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetUserLogDeliveryQuotaResponseBody) *GetUserLogDeliveryQuotaResponse
	GetBody() *GetUserLogDeliveryQuotaResponseBody
}

type GetUserLogDeliveryQuotaResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserLogDeliveryQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserLogDeliveryQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserLogDeliveryQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetUserLogDeliveryQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserLogDeliveryQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserLogDeliveryQuotaResponse) GetBody() *GetUserLogDeliveryQuotaResponseBody {
	return s.Body
}

func (s *GetUserLogDeliveryQuotaResponse) SetHeaders(v map[string]*string) *GetUserLogDeliveryQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetUserLogDeliveryQuotaResponse) SetStatusCode(v int32) *GetUserLogDeliveryQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserLogDeliveryQuotaResponse) SetBody(v *GetUserLogDeliveryQuotaResponseBody) *GetUserLogDeliveryQuotaResponse {
	s.Body = v
	return s
}

func (s *GetUserLogDeliveryQuotaResponse) Validate() error {
	return dara.Validate(s)
}
