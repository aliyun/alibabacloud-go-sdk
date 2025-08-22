// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDcdnSLSRealTimeLogDeliveryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDcdnSLSRealTimeLogDeliveryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDcdnSLSRealTimeLogDeliveryResponse
	GetStatusCode() *int32
	SetBody(v *CreateDcdnSLSRealTimeLogDeliveryResponseBody) *CreateDcdnSLSRealTimeLogDeliveryResponse
	GetBody() *CreateDcdnSLSRealTimeLogDeliveryResponseBody
}

type CreateDcdnSLSRealTimeLogDeliveryResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDcdnSLSRealTimeLogDeliveryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDcdnSLSRealTimeLogDeliveryResponse) GoString() string {
	return s.String()
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponse) GetBody() *CreateDcdnSLSRealTimeLogDeliveryResponseBody {
	return s.Body
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponse) SetHeaders(v map[string]*string) *CreateDcdnSLSRealTimeLogDeliveryResponse {
	s.Headers = v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponse) SetStatusCode(v int32) *CreateDcdnSLSRealTimeLogDeliveryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponse) SetBody(v *CreateDcdnSLSRealTimeLogDeliveryResponseBody) *CreateDcdnSLSRealTimeLogDeliveryResponse {
	s.Body = v
	return s
}

func (s *CreateDcdnSLSRealTimeLogDeliveryResponse) Validate() error {
	return dara.Validate(s)
}
