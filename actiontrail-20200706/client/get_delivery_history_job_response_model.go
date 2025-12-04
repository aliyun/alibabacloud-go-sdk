// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDeliveryHistoryJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDeliveryHistoryJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDeliveryHistoryJobResponse
	GetStatusCode() *int32
	SetBody(v *GetDeliveryHistoryJobResponseBody) *GetDeliveryHistoryJobResponse
	GetBody() *GetDeliveryHistoryJobResponseBody
}

type GetDeliveryHistoryJobResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDeliveryHistoryJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDeliveryHistoryJobResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDeliveryHistoryJobResponse) GoString() string {
	return s.String()
}

func (s *GetDeliveryHistoryJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDeliveryHistoryJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDeliveryHistoryJobResponse) GetBody() *GetDeliveryHistoryJobResponseBody {
	return s.Body
}

func (s *GetDeliveryHistoryJobResponse) SetHeaders(v map[string]*string) *GetDeliveryHistoryJobResponse {
	s.Headers = v
	return s
}

func (s *GetDeliveryHistoryJobResponse) SetStatusCode(v int32) *GetDeliveryHistoryJobResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDeliveryHistoryJobResponse) SetBody(v *GetDeliveryHistoryJobResponseBody) *GetDeliveryHistoryJobResponse {
	s.Body = v
	return s
}

func (s *GetDeliveryHistoryJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
