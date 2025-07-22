// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEventSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEventSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEventSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *GetEventSubscriptionResponseBody) *GetEventSubscriptionResponse
	GetBody() *GetEventSubscriptionResponseBody
}

type GetEventSubscriptionResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEventSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEventSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEventSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *GetEventSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEventSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEventSubscriptionResponse) GetBody() *GetEventSubscriptionResponseBody {
	return s.Body
}

func (s *GetEventSubscriptionResponse) SetHeaders(v map[string]*string) *GetEventSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *GetEventSubscriptionResponse) SetStatusCode(v int32) *GetEventSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEventSubscriptionResponse) SetBody(v *GetEventSubscriptionResponseBody) *GetEventSubscriptionResponse {
	s.Body = v
	return s
}

func (s *GetEventSubscriptionResponse) Validate() error {
	return dara.Validate(s)
}
