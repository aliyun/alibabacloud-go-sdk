// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePullToPushResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLivePullToPushResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLivePullToPushResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLivePullToPushResponseBody) *DescribeLivePullToPushResponse
	GetBody() *DescribeLivePullToPushResponseBody
}

type DescribeLivePullToPushResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLivePullToPushResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLivePullToPushResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullToPushResponse) GoString() string {
	return s.String()
}

func (s *DescribeLivePullToPushResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLivePullToPushResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLivePullToPushResponse) GetBody() *DescribeLivePullToPushResponseBody {
	return s.Body
}

func (s *DescribeLivePullToPushResponse) SetHeaders(v map[string]*string) *DescribeLivePullToPushResponse {
	s.Headers = v
	return s
}

func (s *DescribeLivePullToPushResponse) SetStatusCode(v int32) *DescribeLivePullToPushResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLivePullToPushResponse) SetBody(v *DescribeLivePullToPushResponseBody) *DescribeLivePullToPushResponse {
	s.Body = v
	return s
}

func (s *DescribeLivePullToPushResponse) Validate() error {
	return dara.Validate(s)
}
