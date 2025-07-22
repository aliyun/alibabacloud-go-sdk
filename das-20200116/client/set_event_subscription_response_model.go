// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetEventSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetEventSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetEventSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *SetEventSubscriptionResponseBody) *SetEventSubscriptionResponse
	GetBody() *SetEventSubscriptionResponseBody
}

type SetEventSubscriptionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetEventSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetEventSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetEventSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *SetEventSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetEventSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetEventSubscriptionResponse) GetBody() *SetEventSubscriptionResponseBody {
	return s.Body
}

func (s *SetEventSubscriptionResponse) SetHeaders(v map[string]*string) *SetEventSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *SetEventSubscriptionResponse) SetStatusCode(v int32) *SetEventSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetEventSubscriptionResponse) SetBody(v *SetEventSubscriptionResponseBody) *SetEventSubscriptionResponse {
	s.Body = v
	return s
}

func (s *SetEventSubscriptionResponse) Validate() error {
	return dara.Validate(s)
}
