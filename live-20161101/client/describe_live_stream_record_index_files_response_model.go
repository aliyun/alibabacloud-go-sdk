// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamRecordIndexFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamRecordIndexFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamRecordIndexFilesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamRecordIndexFilesResponseBody) *DescribeLiveStreamRecordIndexFilesResponse
	GetBody() *DescribeLiveStreamRecordIndexFilesResponseBody
}

type DescribeLiveStreamRecordIndexFilesResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamRecordIndexFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamRecordIndexFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamRecordIndexFilesResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) GetBody() *DescribeLiveStreamRecordIndexFilesResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamRecordIndexFilesResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) SetStatusCode(v int32) *DescribeLiveStreamRecordIndexFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) SetBody(v *DescribeLiveStreamRecordIndexFilesResponseBody) *DescribeLiveStreamRecordIndexFilesResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamRecordIndexFilesResponse) Validate() error {
	return dara.Validate(s)
}
