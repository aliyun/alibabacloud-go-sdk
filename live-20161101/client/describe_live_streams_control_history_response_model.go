// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamsControlHistoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamsControlHistoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamsControlHistoryResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamsControlHistoryResponseBody) *DescribeLiveStreamsControlHistoryResponse
	GetBody() *DescribeLiveStreamsControlHistoryResponseBody
}

type DescribeLiveStreamsControlHistoryResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamsControlHistoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamsControlHistoryResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamsControlHistoryResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamsControlHistoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamsControlHistoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamsControlHistoryResponse) GetBody() *DescribeLiveStreamsControlHistoryResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamsControlHistoryResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamsControlHistoryResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponse) SetStatusCode(v int32) *DescribeLiveStreamsControlHistoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponse) SetBody(v *DescribeLiveStreamsControlHistoryResponseBody) *DescribeLiveStreamsControlHistoryResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamsControlHistoryResponse) Validate() error {
	return dara.Validate(s)
}
