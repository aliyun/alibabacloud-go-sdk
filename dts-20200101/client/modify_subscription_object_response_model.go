// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySubscriptionObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySubscriptionObjectResponse
	GetStatusCode() *int32
	SetBody(v *ModifySubscriptionObjectResponseBody) *ModifySubscriptionObjectResponse
	GetBody() *ModifySubscriptionObjectResponseBody
}

type ModifySubscriptionObjectResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySubscriptionObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySubscriptionObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionObjectResponse) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySubscriptionObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySubscriptionObjectResponse) GetBody() *ModifySubscriptionObjectResponseBody {
	return s.Body
}

func (s *ModifySubscriptionObjectResponse) SetHeaders(v map[string]*string) *ModifySubscriptionObjectResponse {
	s.Headers = v
	return s
}

func (s *ModifySubscriptionObjectResponse) SetStatusCode(v int32) *ModifySubscriptionObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySubscriptionObjectResponse) SetBody(v *ModifySubscriptionObjectResponseBody) *ModifySubscriptionObjectResponse {
	s.Body = v
	return s
}

func (s *ModifySubscriptionObjectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
