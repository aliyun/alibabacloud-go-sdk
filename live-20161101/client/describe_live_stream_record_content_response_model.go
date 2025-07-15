// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamRecordContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamRecordContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamRecordContentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamRecordContentResponseBody) *DescribeLiveStreamRecordContentResponse
	GetBody() *DescribeLiveStreamRecordContentResponseBody
}

type DescribeLiveStreamRecordContentResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamRecordContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamRecordContentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordContentResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamRecordContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamRecordContentResponse) GetBody() *DescribeLiveStreamRecordContentResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamRecordContentResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamRecordContentResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamRecordContentResponse) SetStatusCode(v int32) *DescribeLiveStreamRecordContentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamRecordContentResponse) SetBody(v *DescribeLiveStreamRecordContentResponseBody) *DescribeLiveStreamRecordContentResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamRecordContentResponse) Validate() error {
	return dara.Validate(s)
}
