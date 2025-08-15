// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeliveryHistoryJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeliveryHistoryJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeliveryHistoryJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDeliveryHistoryJobResponseBody) *DeleteDeliveryHistoryJobResponse
	GetBody() *DeleteDeliveryHistoryJobResponseBody
}

type DeleteDeliveryHistoryJobResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDeliveryHistoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDeliveryHistoryJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeliveryHistoryJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeliveryHistoryJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeliveryHistoryJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeliveryHistoryJobResponse) GetBody() *DeleteDeliveryHistoryJobResponseBody {
	return s.Body
}

func (s *DeleteDeliveryHistoryJobResponse) SetHeaders(v map[string]*string) *DeleteDeliveryHistoryJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeliveryHistoryJobResponse) SetStatusCode(v int32) *DeleteDeliveryHistoryJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeliveryHistoryJobResponse) SetBody(v *DeleteDeliveryHistoryJobResponseBody) *DeleteDeliveryHistoryJobResponse {
	s.Body = v
	return s
}

func (s *DeleteDeliveryHistoryJobResponse) Validate() error {
	return dara.Validate(s)
}
