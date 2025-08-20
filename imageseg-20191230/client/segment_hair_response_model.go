// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHairResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SegmentHairResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SegmentHairResponse
	GetStatusCode() *int32
	SetBody(v *SegmentHairResponseBody) *SegmentHairResponse
	GetBody() *SegmentHairResponseBody
}

type SegmentHairResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentHairResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentHairResponse) String() string {
	return dara.Prettify(s)
}

func (s SegmentHairResponse) GoString() string {
	return s.String()
}

func (s *SegmentHairResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SegmentHairResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SegmentHairResponse) GetBody() *SegmentHairResponseBody {
	return s.Body
}

func (s *SegmentHairResponse) SetHeaders(v map[string]*string) *SegmentHairResponse {
	s.Headers = v
	return s
}

func (s *SegmentHairResponse) SetStatusCode(v int32) *SegmentHairResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentHairResponse) SetBody(v *SegmentHairResponseBody) *SegmentHairResponse {
	s.Body = v
	return s
}

func (s *SegmentHairResponse) Validate() error {
	return dara.Validate(s)
}
