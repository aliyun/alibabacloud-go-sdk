// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHeadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SegmentHeadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SegmentHeadResponse
	GetStatusCode() *int32
	SetBody(v *SegmentHeadResponseBody) *SegmentHeadResponse
	GetBody() *SegmentHeadResponseBody
}

type SegmentHeadResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentHeadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentHeadResponse) String() string {
	return dara.Prettify(s)
}

func (s SegmentHeadResponse) GoString() string {
	return s.String()
}

func (s *SegmentHeadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SegmentHeadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SegmentHeadResponse) GetBody() *SegmentHeadResponseBody {
	return s.Body
}

func (s *SegmentHeadResponse) SetHeaders(v map[string]*string) *SegmentHeadResponse {
	s.Headers = v
	return s
}

func (s *SegmentHeadResponse) SetStatusCode(v int32) *SegmentHeadResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentHeadResponse) SetBody(v *SegmentHeadResponseBody) *SegmentHeadResponse {
	s.Body = v
	return s
}

func (s *SegmentHeadResponse) Validate() error {
	return dara.Validate(s)
}
