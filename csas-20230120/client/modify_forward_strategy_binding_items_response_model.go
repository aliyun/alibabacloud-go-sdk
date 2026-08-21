// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyForwardStrategyBindingItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyForwardStrategyBindingItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyForwardStrategyBindingItemsResponse
	GetStatusCode() *int32
	SetBody(v *ModifyForwardStrategyBindingItemsResponseBody) *ModifyForwardStrategyBindingItemsResponse
	GetBody() *ModifyForwardStrategyBindingItemsResponseBody
}

type ModifyForwardStrategyBindingItemsResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyForwardStrategyBindingItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyForwardStrategyBindingItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyForwardStrategyBindingItemsResponse) GoString() string {
	return s.String()
}

func (s *ModifyForwardStrategyBindingItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyForwardStrategyBindingItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyForwardStrategyBindingItemsResponse) GetBody() *ModifyForwardStrategyBindingItemsResponseBody {
	return s.Body
}

func (s *ModifyForwardStrategyBindingItemsResponse) SetHeaders(v map[string]*string) *ModifyForwardStrategyBindingItemsResponse {
	s.Headers = v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponse) SetStatusCode(v int32) *ModifyForwardStrategyBindingItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponse) SetBody(v *ModifyForwardStrategyBindingItemsResponseBody) *ModifyForwardStrategyBindingItemsResponse {
	s.Body = v
	return s
}

func (s *ModifyForwardStrategyBindingItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
