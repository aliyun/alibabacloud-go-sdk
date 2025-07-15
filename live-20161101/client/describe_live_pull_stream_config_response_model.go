// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLivePullStreamConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLivePullStreamConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLivePullStreamConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLivePullStreamConfigResponseBody) *DescribeLivePullStreamConfigResponse
	GetBody() *DescribeLivePullStreamConfigResponseBody
}

type DescribeLivePullStreamConfigResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLivePullStreamConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLivePullStreamConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLivePullStreamConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeLivePullStreamConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLivePullStreamConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLivePullStreamConfigResponse) GetBody() *DescribeLivePullStreamConfigResponseBody {
	return s.Body
}

func (s *DescribeLivePullStreamConfigResponse) SetHeaders(v map[string]*string) *DescribeLivePullStreamConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeLivePullStreamConfigResponse) SetStatusCode(v int32) *DescribeLivePullStreamConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLivePullStreamConfigResponse) SetBody(v *DescribeLivePullStreamConfigResponseBody) *DescribeLivePullStreamConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeLivePullStreamConfigResponse) Validate() error {
	return dara.Validate(s)
}
