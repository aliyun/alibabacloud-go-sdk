// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySubscriptionDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySubscriptionDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySubscriptionDetailResponse
	GetStatusCode() *int32
	SetBody(v *QuerySubscriptionDetailResponseBody) *QuerySubscriptionDetailResponse
	GetBody() *QuerySubscriptionDetailResponseBody
}

type QuerySubscriptionDetailResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySubscriptionDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySubscriptionDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySubscriptionDetailResponse) GoString() string {
	return s.String()
}

func (s *QuerySubscriptionDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySubscriptionDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySubscriptionDetailResponse) GetBody() *QuerySubscriptionDetailResponseBody {
	return s.Body
}

func (s *QuerySubscriptionDetailResponse) SetHeaders(v map[string]*string) *QuerySubscriptionDetailResponse {
	s.Headers = v
	return s
}

func (s *QuerySubscriptionDetailResponse) SetStatusCode(v int32) *QuerySubscriptionDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySubscriptionDetailResponse) SetBody(v *QuerySubscriptionDetailResponseBody) *QuerySubscriptionDetailResponse {
	s.Body = v
	return s
}

func (s *QuerySubscriptionDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
