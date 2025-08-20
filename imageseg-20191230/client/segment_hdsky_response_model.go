// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentHDSkyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SegmentHDSkyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SegmentHDSkyResponse
	GetStatusCode() *int32
	SetBody(v *SegmentHDSkyResponseBody) *SegmentHDSkyResponse
	GetBody() *SegmentHDSkyResponseBody
}

type SegmentHDSkyResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentHDSkyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentHDSkyResponse) String() string {
	return dara.Prettify(s)
}

func (s SegmentHDSkyResponse) GoString() string {
	return s.String()
}

func (s *SegmentHDSkyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SegmentHDSkyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SegmentHDSkyResponse) GetBody() *SegmentHDSkyResponseBody {
	return s.Body
}

func (s *SegmentHDSkyResponse) SetHeaders(v map[string]*string) *SegmentHDSkyResponse {
	s.Headers = v
	return s
}

func (s *SegmentHDSkyResponse) SetStatusCode(v int32) *SegmentHDSkyResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentHDSkyResponse) SetBody(v *SegmentHDSkyResponseBody) *SegmentHDSkyResponse {
	s.Body = v
	return s
}

func (s *SegmentHDSkyResponse) Validate() error {
	return dara.Validate(s)
}
