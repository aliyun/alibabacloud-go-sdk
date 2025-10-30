// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UnbindSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UnbindSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *UnbindSubscriptionResponseBody) *UnbindSubscriptionResponse
	GetBody() *UnbindSubscriptionResponseBody
}

type UnbindSubscriptionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UnbindSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UnbindSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s UnbindSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *UnbindSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UnbindSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UnbindSubscriptionResponse) GetBody() *UnbindSubscriptionResponseBody {
	return s.Body
}

func (s *UnbindSubscriptionResponse) SetHeaders(v map[string]*string) *UnbindSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *UnbindSubscriptionResponse) SetStatusCode(v int32) *UnbindSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *UnbindSubscriptionResponse) SetBody(v *UnbindSubscriptionResponseBody) *UnbindSubscriptionResponse {
	s.Body = v
	return s
}

func (s *UnbindSubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
