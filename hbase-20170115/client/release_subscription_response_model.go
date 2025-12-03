// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseSubscriptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseSubscriptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseSubscriptionResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseSubscriptionResponseBody) *ReleaseSubscriptionResponse
	GetBody() *ReleaseSubscriptionResponseBody
}

type ReleaseSubscriptionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseSubscriptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseSubscriptionResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseSubscriptionResponse) GoString() string {
	return s.String()
}

func (s *ReleaseSubscriptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseSubscriptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseSubscriptionResponse) GetBody() *ReleaseSubscriptionResponseBody {
	return s.Body
}

func (s *ReleaseSubscriptionResponse) SetHeaders(v map[string]*string) *ReleaseSubscriptionResponse {
	s.Headers = v
	return s
}

func (s *ReleaseSubscriptionResponse) SetStatusCode(v int32) *ReleaseSubscriptionResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseSubscriptionResponse) SetBody(v *ReleaseSubscriptionResponseBody) *ReleaseSubscriptionResponse {
	s.Body = v
	return s
}

func (s *ReleaseSubscriptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
