// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelSubscriptionBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelSubscriptionBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelSubscriptionBillResponse
	GetStatusCode() *int32
	SetBody(v *CancelSubscriptionBillResponseBody) *CancelSubscriptionBillResponse
	GetBody() *CancelSubscriptionBillResponseBody
}

type CancelSubscriptionBillResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelSubscriptionBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelSubscriptionBillResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelSubscriptionBillResponse) GoString() string {
	return s.String()
}

func (s *CancelSubscriptionBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelSubscriptionBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelSubscriptionBillResponse) GetBody() *CancelSubscriptionBillResponseBody {
	return s.Body
}

func (s *CancelSubscriptionBillResponse) SetHeaders(v map[string]*string) *CancelSubscriptionBillResponse {
	s.Headers = v
	return s
}

func (s *CancelSubscriptionBillResponse) SetStatusCode(v int32) *CancelSubscriptionBillResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelSubscriptionBillResponse) SetBody(v *CancelSubscriptionBillResponseBody) *CancelSubscriptionBillResponse {
	s.Body = v
	return s
}

func (s *CancelSubscriptionBillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
