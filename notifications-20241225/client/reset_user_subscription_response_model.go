// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetUserSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResetUserSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResetUserSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *ResetUserSubscriptionResponseBody) *ResetUserSubscriptionResponse
	GetBody() *ResetUserSubscriptionResponseBody
}

type ResetUserSubscriptionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResetUserSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResetUserSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ResetUserSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *ResetUserSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResetUserSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResetUserSubscriptionResponse) GetBody() *ResetUserSubscriptionResponseBody {
	return s.Body
}

func (s *ResetUserSubscriptionResponse) SetHeaders(v map[string]*string) *ResetUserSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *ResetUserSubscriptionResponse) SetStatusCode(v int32) *ResetUserSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ResetUserSubscriptionResponse) SetBody(v *ResetUserSubscriptionResponseBody) *ResetUserSubscriptionResponse {
	s.Body = v
	return s
}

func (s *ResetUserSubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
