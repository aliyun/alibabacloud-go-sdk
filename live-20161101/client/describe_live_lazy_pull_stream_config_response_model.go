// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveLazyPullStreamConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveLazyPullStreamConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveLazyPullStreamConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveLazyPullStreamConfigResponseBody) *DescribeLiveLazyPullStreamConfigResponse
	GetBody() *DescribeLiveLazyPullStreamConfigResponseBody
}

type DescribeLiveLazyPullStreamConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveLazyPullStreamConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveLazyPullStreamConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveLazyPullStreamConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveLazyPullStreamConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveLazyPullStreamConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveLazyPullStreamConfigResponse) GetBody() *DescribeLiveLazyPullStreamConfigResponseBody {
	return s.Body
}

func (s *DescribeLiveLazyPullStreamConfigResponse) SetHeaders(v map[string]*string) *DescribeLiveLazyPullStreamConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponse) SetStatusCode(v int32) *DescribeLiveLazyPullStreamConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponse) SetBody(v *DescribeLiveLazyPullStreamConfigResponseBody) *DescribeLiveLazyPullStreamConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveLazyPullStreamConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
