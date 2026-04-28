// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListForwardStrategyBindingItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListForwardStrategyBindingItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListForwardStrategyBindingItemsResponse
	GetStatusCode() *int32
	SetBody(v *ListForwardStrategyBindingItemsResponseBody) *ListForwardStrategyBindingItemsResponse
	GetBody() *ListForwardStrategyBindingItemsResponseBody
}

type ListForwardStrategyBindingItemsResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListForwardStrategyBindingItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListForwardStrategyBindingItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListForwardStrategyBindingItemsResponse) GoString() string {
	return s.String()
}

func (s *ListForwardStrategyBindingItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListForwardStrategyBindingItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListForwardStrategyBindingItemsResponse) GetBody() *ListForwardStrategyBindingItemsResponseBody {
	return s.Body
}

func (s *ListForwardStrategyBindingItemsResponse) SetHeaders(v map[string]*string) *ListForwardStrategyBindingItemsResponse {
	s.Headers = v
	return s
}

func (s *ListForwardStrategyBindingItemsResponse) SetStatusCode(v int32) *ListForwardStrategyBindingItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListForwardStrategyBindingItemsResponse) SetBody(v *ListForwardStrategyBindingItemsResponseBody) *ListForwardStrategyBindingItemsResponse {
	s.Body = v
	return s
}

func (s *ListForwardStrategyBindingItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
