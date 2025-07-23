// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSubscriptionAttributesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetSubscriptionAttributesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetSubscriptionAttributesResponse
	GetStatusCode() *int32
	SetBody(v *SetSubscriptionAttributesResponseBody) *SetSubscriptionAttributesResponse
	GetBody() *SetSubscriptionAttributesResponseBody
}

type SetSubscriptionAttributesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetSubscriptionAttributesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetSubscriptionAttributesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetSubscriptionAttributesResponse) GoString() string {
	return s.String()
}

func (s *SetSubscriptionAttributesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetSubscriptionAttributesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetSubscriptionAttributesResponse) GetBody() *SetSubscriptionAttributesResponseBody {
	return s.Body
}

func (s *SetSubscriptionAttributesResponse) SetHeaders(v map[string]*string) *SetSubscriptionAttributesResponse {
	s.Headers = v
	return s
}

func (s *SetSubscriptionAttributesResponse) SetStatusCode(v int32) *SetSubscriptionAttributesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetSubscriptionAttributesResponse) SetBody(v *SetSubscriptionAttributesResponseBody) *SetSubscriptionAttributesResponse {
	s.Body = v
	return s
}

func (s *SetSubscriptionAttributesResponse) Validate() error {
	return dara.Validate(s)
}
