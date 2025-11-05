// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubscriptionBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubscriptionBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubscriptionBillResponse
	GetStatusCode() *int32
	SetBody(v *SubscriptionBillResponseBody) *SubscriptionBillResponse
	GetBody() *SubscriptionBillResponseBody
}

type SubscriptionBillResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubscriptionBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubscriptionBillResponse) String() string {
	return dara.Prettify(s)
}

func (s SubscriptionBillResponse) GoString() string {
	return s.String()
}

func (s *SubscriptionBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubscriptionBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubscriptionBillResponse) GetBody() *SubscriptionBillResponseBody {
	return s.Body
}

func (s *SubscriptionBillResponse) SetHeaders(v map[string]*string) *SubscriptionBillResponse {
	s.Headers = v
	return s
}

func (s *SubscriptionBillResponse) SetStatusCode(v int32) *SubscriptionBillResponse {
	s.StatusCode = &v
	return s
}

func (s *SubscriptionBillResponse) SetBody(v *SubscriptionBillResponseBody) *SubscriptionBillResponse {
	s.Body = v
	return s
}

func (s *SubscriptionBillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
