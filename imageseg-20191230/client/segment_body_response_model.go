// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentBodyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SegmentBodyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SegmentBodyResponse
	GetStatusCode() *int32
	SetBody(v *SegmentBodyResponseBody) *SegmentBodyResponse
	GetBody() *SegmentBodyResponseBody
}

type SegmentBodyResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentBodyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentBodyResponse) String() string {
	return dara.Prettify(s)
}

func (s SegmentBodyResponse) GoString() string {
	return s.String()
}

func (s *SegmentBodyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SegmentBodyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SegmentBodyResponse) GetBody() *SegmentBodyResponseBody {
	return s.Body
}

func (s *SegmentBodyResponse) SetHeaders(v map[string]*string) *SegmentBodyResponse {
	s.Headers = v
	return s
}

func (s *SegmentBodyResponse) SetStatusCode(v int32) *SegmentBodyResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentBodyResponse) SetBody(v *SegmentBodyResponseBody) *SegmentBodyResponse {
	s.Body = v
	return s
}

func (s *SegmentBodyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
