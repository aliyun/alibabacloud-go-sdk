// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRealtimeDeliveryFieldResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRealtimeDeliveryFieldResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRealtimeDeliveryFieldResponse
	GetStatusCode() *int32
	SetBody(v *GetRealtimeDeliveryFieldResponseBody) *GetRealtimeDeliveryFieldResponse
	GetBody() *GetRealtimeDeliveryFieldResponseBody
}

type GetRealtimeDeliveryFieldResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRealtimeDeliveryFieldResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRealtimeDeliveryFieldResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRealtimeDeliveryFieldResponse) GoString() string {
	return s.String()
}

func (s *GetRealtimeDeliveryFieldResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRealtimeDeliveryFieldResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRealtimeDeliveryFieldResponse) GetBody() *GetRealtimeDeliveryFieldResponseBody {
	return s.Body
}

func (s *GetRealtimeDeliveryFieldResponse) SetHeaders(v map[string]*string) *GetRealtimeDeliveryFieldResponse {
	s.Headers = v
	return s
}

func (s *GetRealtimeDeliveryFieldResponse) SetStatusCode(v int32) *GetRealtimeDeliveryFieldResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRealtimeDeliveryFieldResponse) SetBody(v *GetRealtimeDeliveryFieldResponseBody) *GetRealtimeDeliveryFieldResponse {
	s.Body = v
	return s
}

func (s *GetRealtimeDeliveryFieldResponse) Validate() error {
	return dara.Validate(s)
}
