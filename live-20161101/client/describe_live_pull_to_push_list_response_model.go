// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePullToPushListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLivePullToPushListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLivePullToPushListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLivePullToPushListResponseBody) *DescribeLivePullToPushListResponse
	GetBody() *DescribeLivePullToPushListResponseBody
}

type DescribeLivePullToPushListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLivePullToPushListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLivePullToPushListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullToPushListResponse) GoString() string {
	return s.String()
}

func (s *DescribeLivePullToPushListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLivePullToPushListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLivePullToPushListResponse) GetBody() *DescribeLivePullToPushListResponseBody {
	return s.Body
}

func (s *DescribeLivePullToPushListResponse) SetHeaders(v map[string]*string) *DescribeLivePullToPushListResponse {
	s.Headers = v
	return s
}

func (s *DescribeLivePullToPushListResponse) SetStatusCode(v int32) *DescribeLivePullToPushListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLivePullToPushListResponse) SetBody(v *DescribeLivePullToPushListResponseBody) *DescribeLivePullToPushListResponse {
	s.Body = v
	return s
}

func (s *DescribeLivePullToPushListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
