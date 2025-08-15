// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeliveryHistoryJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeliveryHistoryJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeliveryHistoryJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateDeliveryHistoryJobResponseBody) *CreateDeliveryHistoryJobResponse
	GetBody() *CreateDeliveryHistoryJobResponseBody
}

type CreateDeliveryHistoryJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeliveryHistoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeliveryHistoryJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDeliveryHistoryJobResponse) GoString() string {
	return s.String()
}

func (s *CreateDeliveryHistoryJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeliveryHistoryJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeliveryHistoryJobResponse) GetBody() *CreateDeliveryHistoryJobResponseBody {
	return s.Body
}

func (s *CreateDeliveryHistoryJobResponse) SetHeaders(v map[string]*string) *CreateDeliveryHistoryJobResponse {
	s.Headers = v
	return s
}

func (s *CreateDeliveryHistoryJobResponse) SetStatusCode(v int32) *CreateDeliveryHistoryJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDeliveryHistoryJobResponse) SetBody(v *CreateDeliveryHistoryJobResponseBody) *CreateDeliveryHistoryJobResponse {
	s.Body = v
	return s
}

func (s *CreateDeliveryHistoryJobResponse) Validate() error {
	return dara.Validate(s)
}
