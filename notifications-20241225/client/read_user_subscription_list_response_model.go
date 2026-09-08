// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadUserSubscriptionListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadUserSubscriptionListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadUserSubscriptionListResponse
	GetStatusCode() *int32
	SetBody(v *ReadUserSubscriptionListResponseBody) *ReadUserSubscriptionListResponse
	GetBody() *ReadUserSubscriptionListResponseBody
}

type ReadUserSubscriptionListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadUserSubscriptionListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadUserSubscriptionListResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadUserSubscriptionListResponse) GoString() string {
	return s.String()
}

func (s *ReadUserSubscriptionListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadUserSubscriptionListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadUserSubscriptionListResponse) GetBody() *ReadUserSubscriptionListResponseBody {
	return s.Body
}

func (s *ReadUserSubscriptionListResponse) SetHeaders(v map[string]*string) *ReadUserSubscriptionListResponse {
	s.Headers = v
	return s
}

func (s *ReadUserSubscriptionListResponse) SetStatusCode(v int32) *ReadUserSubscriptionListResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadUserSubscriptionListResponse) SetBody(v *ReadUserSubscriptionListResponseBody) *ReadUserSubscriptionListResponse {
	s.Body = v
	return s
}

func (s *ReadUserSubscriptionListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
