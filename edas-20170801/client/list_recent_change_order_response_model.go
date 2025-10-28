// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRecentChangeOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRecentChangeOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRecentChangeOrderResponse
	GetStatusCode() *int32
	SetBody(v *ListRecentChangeOrderResponseBody) *ListRecentChangeOrderResponse
	GetBody() *ListRecentChangeOrderResponseBody
}

type ListRecentChangeOrderResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecentChangeOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecentChangeOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRecentChangeOrderResponse) GoString() string {
	return s.String()
}

func (s *ListRecentChangeOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRecentChangeOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRecentChangeOrderResponse) GetBody() *ListRecentChangeOrderResponseBody {
	return s.Body
}

func (s *ListRecentChangeOrderResponse) SetHeaders(v map[string]*string) *ListRecentChangeOrderResponse {
	s.Headers = v
	return s
}

func (s *ListRecentChangeOrderResponse) SetStatusCode(v int32) *ListRecentChangeOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecentChangeOrderResponse) SetBody(v *ListRecentChangeOrderResponseBody) *ListRecentChangeOrderResponse {
	s.Body = v
	return s
}

func (s *ListRecentChangeOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
