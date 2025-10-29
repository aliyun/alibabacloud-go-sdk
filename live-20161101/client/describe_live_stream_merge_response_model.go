// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamMergeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamMergeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamMergeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamMergeResponseBody) *DescribeLiveStreamMergeResponse
	GetBody() *DescribeLiveStreamMergeResponseBody
}

type DescribeLiveStreamMergeResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamMergeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamMergeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamMergeResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamMergeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamMergeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamMergeResponse) GetBody() *DescribeLiveStreamMergeResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamMergeResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamMergeResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamMergeResponse) SetStatusCode(v int32) *DescribeLiveStreamMergeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamMergeResponse) SetBody(v *DescribeLiveStreamMergeResponseBody) *DescribeLiveStreamMergeResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamMergeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
