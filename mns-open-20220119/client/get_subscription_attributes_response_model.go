// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubscriptionAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSubscriptionAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSubscriptionAttributesResponse
	GetStatusCode() *int32
	SetBody(v *GetSubscriptionAttributesResponseBody) *GetSubscriptionAttributesResponse
	GetBody() *GetSubscriptionAttributesResponseBody
}

type GetSubscriptionAttributesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSubscriptionAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSubscriptionAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSubscriptionAttributesResponse) GoString() string {
	return s.String()
}

func (s *GetSubscriptionAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSubscriptionAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSubscriptionAttributesResponse) GetBody() *GetSubscriptionAttributesResponseBody {
	return s.Body
}

func (s *GetSubscriptionAttributesResponse) SetHeaders(v map[string]*string) *GetSubscriptionAttributesResponse {
	s.Headers = v
	return s
}

func (s *GetSubscriptionAttributesResponse) SetStatusCode(v int32) *GetSubscriptionAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSubscriptionAttributesResponse) SetBody(v *GetSubscriptionAttributesResponseBody) *GetSubscriptionAttributesResponse {
	s.Body = v
	return s
}

func (s *GetSubscriptionAttributesResponse) Validate() error {
	return dara.Validate(s)
}
