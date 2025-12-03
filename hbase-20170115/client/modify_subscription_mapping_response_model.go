// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySubscriptionMappingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySubscriptionMappingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySubscriptionMappingResponse
	GetStatusCode() *int32
	SetBody(v *ModifySubscriptionMappingResponseBody) *ModifySubscriptionMappingResponse
	GetBody() *ModifySubscriptionMappingResponseBody
}

type ModifySubscriptionMappingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySubscriptionMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySubscriptionMappingResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySubscriptionMappingResponse) GoString() string {
	return s.String()
}

func (s *ModifySubscriptionMappingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySubscriptionMappingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySubscriptionMappingResponse) GetBody() *ModifySubscriptionMappingResponseBody {
	return s.Body
}

func (s *ModifySubscriptionMappingResponse) SetHeaders(v map[string]*string) *ModifySubscriptionMappingResponse {
	s.Headers = v
	return s
}

func (s *ModifySubscriptionMappingResponse) SetStatusCode(v int32) *ModifySubscriptionMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySubscriptionMappingResponse) SetBody(v *ModifySubscriptionMappingResponseBody) *ModifySubscriptionMappingResponse {
	s.Body = v
	return s
}

func (s *ModifySubscriptionMappingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
