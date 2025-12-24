// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentSkyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SegmentSkyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SegmentSkyResponse
	GetStatusCode() *int32
	SetBody(v *SegmentSkyResponseBody) *SegmentSkyResponse
	GetBody() *SegmentSkyResponseBody
}

type SegmentSkyResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentSkyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentSkyResponse) String() string {
	return dara.Prettify(s)
}

func (s SegmentSkyResponse) GoString() string {
	return s.String()
}

func (s *SegmentSkyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SegmentSkyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SegmentSkyResponse) GetBody() *SegmentSkyResponseBody {
	return s.Body
}

func (s *SegmentSkyResponse) SetHeaders(v map[string]*string) *SegmentSkyResponse {
	s.Headers = v
	return s
}

func (s *SegmentSkyResponse) SetStatusCode(v int32) *SegmentSkyResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentSkyResponse) SetBody(v *SegmentSkyResponseBody) *SegmentSkyResponse {
	s.Body = v
	return s
}

func (s *SegmentSkyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
