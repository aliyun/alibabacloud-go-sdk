// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifySubscriptionResponseBody) *ModifySubscriptionResponse
	GetBody() *ModifySubscriptionResponseBody
}

type ModifySubscriptionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySubscriptionResponse) GetBody() *ModifySubscriptionResponseBody {
	return s.Body
}

func (s *ModifySubscriptionResponse) SetHeaders(v map[string]*string) *ModifySubscriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifySubscriptionResponse) SetStatusCode(v int32) *ModifySubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySubscriptionResponse) SetBody(v *ModifySubscriptionResponseBody) *ModifySubscriptionResponse {
	s.Body = v
	return s
}

func (s *ModifySubscriptionResponse) Validate() error {
	return dara.Validate(s)
}
