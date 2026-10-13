// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProductSubscriptionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListProductSubscriptionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListProductSubscriptionsResponse
	GetStatusCode() *int32
	SetBody(v *ListProductSubscriptionsResponseBody) *ListProductSubscriptionsResponse
	GetBody() *ListProductSubscriptionsResponseBody
}

type ListProductSubscriptionsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListProductSubscriptionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListProductSubscriptionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListProductSubscriptionsResponse) GoString() string {
	return s.String()
}

func (s *ListProductSubscriptionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListProductSubscriptionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListProductSubscriptionsResponse) GetBody() *ListProductSubscriptionsResponseBody {
	return s.Body
}

func (s *ListProductSubscriptionsResponse) SetHeaders(v map[string]*string) *ListProductSubscriptionsResponse {
	s.Headers = v
	return s
}

func (s *ListProductSubscriptionsResponse) SetStatusCode(v int32) *ListProductSubscriptionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListProductSubscriptionsResponse) SetBody(v *ListProductSubscriptionsResponseBody) *ListProductSubscriptionsResponse {
	s.Body = v
	return s
}

func (s *ListProductSubscriptionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
