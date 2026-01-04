// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubscriptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubscriptionsResponse
	GetStatusCode() *int32
	SetBody(v *ListSubscriptionsResponseBody) *ListSubscriptionsResponse
	GetBody() *ListSubscriptionsResponseBody
}

type ListSubscriptionsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubscriptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubscriptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionsResponse) GoString() string {
	return s.String()
}

func (s *ListSubscriptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubscriptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubscriptionsResponse) GetBody() *ListSubscriptionsResponseBody {
	return s.Body
}

func (s *ListSubscriptionsResponse) SetHeaders(v map[string]*string) *ListSubscriptionsResponse {
	s.Headers = v
	return s
}

func (s *ListSubscriptionsResponse) SetStatusCode(v int32) *ListSubscriptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubscriptionsResponse) SetBody(v *ListSubscriptionsResponseBody) *ListSubscriptionsResponse {
	s.Body = v
	return s
}

func (s *ListSubscriptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
