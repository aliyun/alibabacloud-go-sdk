// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubscriptionSharedPackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSubscriptionSharedPackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSubscriptionSharedPackagesResponse
	GetStatusCode() *int32
	SetBody(v *ListSubscriptionSharedPackagesResponseBody) *ListSubscriptionSharedPackagesResponse
	GetBody() *ListSubscriptionSharedPackagesResponseBody
}

type ListSubscriptionSharedPackagesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSubscriptionSharedPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSubscriptionSharedPackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSubscriptionSharedPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListSubscriptionSharedPackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSubscriptionSharedPackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSubscriptionSharedPackagesResponse) GetBody() *ListSubscriptionSharedPackagesResponseBody {
	return s.Body
}

func (s *ListSubscriptionSharedPackagesResponse) SetHeaders(v map[string]*string) *ListSubscriptionSharedPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListSubscriptionSharedPackagesResponse) SetStatusCode(v int32) *ListSubscriptionSharedPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSubscriptionSharedPackagesResponse) SetBody(v *ListSubscriptionSharedPackagesResponseBody) *ListSubscriptionSharedPackagesResponse {
	s.Body = v
	return s
}

func (s *ListSubscriptionSharedPackagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
