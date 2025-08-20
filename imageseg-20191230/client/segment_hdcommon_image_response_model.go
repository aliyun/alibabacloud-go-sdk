// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHDCommonImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SegmentHDCommonImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SegmentHDCommonImageResponse
	GetStatusCode() *int32
	SetBody(v *SegmentHDCommonImageResponseBody) *SegmentHDCommonImageResponse
	GetBody() *SegmentHDCommonImageResponseBody
}

type SegmentHDCommonImageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentHDCommonImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentHDCommonImageResponse) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDCommonImageResponse) GoString() string {
	return s.String()
}

func (s *SegmentHDCommonImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SegmentHDCommonImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SegmentHDCommonImageResponse) GetBody() *SegmentHDCommonImageResponseBody {
	return s.Body
}

func (s *SegmentHDCommonImageResponse) SetHeaders(v map[string]*string) *SegmentHDCommonImageResponse {
	s.Headers = v
	return s
}

func (s *SegmentHDCommonImageResponse) SetStatusCode(v int32) *SegmentHDCommonImageResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentHDCommonImageResponse) SetBody(v *SegmentHDCommonImageResponseBody) *SegmentHDCommonImageResponse {
	s.Body = v
	return s
}

func (s *SegmentHDCommonImageResponse) Validate() error {
	return dara.Validate(s)
}
