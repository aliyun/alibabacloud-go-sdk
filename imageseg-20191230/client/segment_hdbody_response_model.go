// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHDBodyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SegmentHDBodyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SegmentHDBodyResponse
	GetStatusCode() *int32
	SetBody(v *SegmentHDBodyResponseBody) *SegmentHDBodyResponse
	GetBody() *SegmentHDBodyResponseBody
}

type SegmentHDBodyResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentHDBodyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentHDBodyResponse) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDBodyResponse) GoString() string {
	return s.String()
}

func (s *SegmentHDBodyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SegmentHDBodyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SegmentHDBodyResponse) GetBody() *SegmentHDBodyResponseBody {
	return s.Body
}

func (s *SegmentHDBodyResponse) SetHeaders(v map[string]*string) *SegmentHDBodyResponse {
	s.Headers = v
	return s
}

func (s *SegmentHDBodyResponse) SetStatusCode(v int32) *SegmentHDBodyResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentHDBodyResponse) SetBody(v *SegmentHDBodyResponseBody) *SegmentHDBodyResponse {
	s.Body = v
	return s
}

func (s *SegmentHDBodyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
