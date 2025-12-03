// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionDescriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySubscriptionDescriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySubscriptionDescriptionResponse
	GetStatusCode() *int32
	SetBody(v *ModifySubscriptionDescriptionResponseBody) *ModifySubscriptionDescriptionResponse
	GetBody() *ModifySubscriptionDescriptionResponseBody
}

type ModifySubscriptionDescriptionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySubscriptionDescriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySubscriptionDescriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionDescriptionResponse) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionDescriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySubscriptionDescriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySubscriptionDescriptionResponse) GetBody() *ModifySubscriptionDescriptionResponseBody {
	return s.Body
}

func (s *ModifySubscriptionDescriptionResponse) SetHeaders(v map[string]*string) *ModifySubscriptionDescriptionResponse {
	s.Headers = v
	return s
}

func (s *ModifySubscriptionDescriptionResponse) SetStatusCode(v int32) *ModifySubscriptionDescriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySubscriptionDescriptionResponse) SetBody(v *ModifySubscriptionDescriptionResponseBody) *ModifySubscriptionDescriptionResponse {
	s.Body = v
	return s
}

func (s *ModifySubscriptionDescriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
