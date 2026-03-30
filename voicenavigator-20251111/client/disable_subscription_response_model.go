// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *DisableSubscriptionResponseBody) *DisableSubscriptionResponse
	GetBody() *DisableSubscriptionResponseBody
}

type DisableSubscriptionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *DisableSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableSubscriptionResponse) GetBody() *DisableSubscriptionResponseBody {
	return s.Body
}

func (s *DisableSubscriptionResponse) SetHeaders(v map[string]*string) *DisableSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *DisableSubscriptionResponse) SetStatusCode(v int32) *DisableSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableSubscriptionResponse) SetBody(v *DisableSubscriptionResponseBody) *DisableSubscriptionResponse {
	s.Body = v
	return s
}

func (s *DisableSubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
