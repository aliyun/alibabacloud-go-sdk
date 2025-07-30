// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubscriptionInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSubscriptionInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSubscriptionInstanceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSubscriptionInstanceResponseBody) *DeleteSubscriptionInstanceResponse
	GetBody() *DeleteSubscriptionInstanceResponseBody
}

type DeleteSubscriptionInstanceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSubscriptionInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSubscriptionInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubscriptionInstanceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSubscriptionInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSubscriptionInstanceResponse) GetBody() *DeleteSubscriptionInstanceResponseBody {
	return s.Body
}

func (s *DeleteSubscriptionInstanceResponse) SetHeaders(v map[string]*string) *DeleteSubscriptionInstanceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubscriptionInstanceResponse) SetStatusCode(v int32) *DeleteSubscriptionInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubscriptionInstanceResponse) SetBody(v *DeleteSubscriptionInstanceResponseBody) *DeleteSubscriptionInstanceResponse {
	s.Body = v
	return s
}

func (s *DeleteSubscriptionInstanceResponse) Validate() error {
	return dara.Validate(s)
}
