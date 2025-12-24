// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentCommonImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SegmentCommonImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SegmentCommonImageResponse
	GetStatusCode() *int32
	SetBody(v *SegmentCommonImageResponseBody) *SegmentCommonImageResponse
	GetBody() *SegmentCommonImageResponseBody
}

type SegmentCommonImageResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentCommonImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentCommonImageResponse) String() string {
	return dara.Prettify(s)
}

func (s SegmentCommonImageResponse) GoString() string {
	return s.String()
}

func (s *SegmentCommonImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SegmentCommonImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SegmentCommonImageResponse) GetBody() *SegmentCommonImageResponseBody {
	return s.Body
}

func (s *SegmentCommonImageResponse) SetHeaders(v map[string]*string) *SegmentCommonImageResponse {
	s.Headers = v
	return s
}

func (s *SegmentCommonImageResponse) SetStatusCode(v int32) *SegmentCommonImageResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentCommonImageResponse) SetBody(v *SegmentCommonImageResponseBody) *SegmentCommonImageResponse {
	s.Body = v
	return s
}

func (s *SegmentCommonImageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
