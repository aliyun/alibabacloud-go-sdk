// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSubscriptionResponseBody) *DeleteSubscriptionResponse
	GetBody() *DeleteSubscriptionResponseBody
}

type DeleteSubscriptionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *DeleteSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSubscriptionResponse) GetBody() *DeleteSubscriptionResponseBody {
	return s.Body
}

func (s *DeleteSubscriptionResponse) SetHeaders(v map[string]*string) *DeleteSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *DeleteSubscriptionResponse) SetStatusCode(v int32) *DeleteSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSubscriptionResponse) SetBody(v *DeleteSubscriptionResponseBody) *DeleteSubscriptionResponse {
	s.Body = v
	return s
}

func (s *DeleteSubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
