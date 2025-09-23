// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentVideoBodyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SegmentVideoBodyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SegmentVideoBodyResponse
	GetStatusCode() *int32
	SetBody(v *SegmentVideoBodyResponseBody) *SegmentVideoBodyResponse
	GetBody() *SegmentVideoBodyResponseBody
}

type SegmentVideoBodyResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentVideoBodyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentVideoBodyResponse) String() string {
	return dara.Prettify(s)
}

func (s SegmentVideoBodyResponse) GoString() string {
	return s.String()
}

func (s *SegmentVideoBodyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SegmentVideoBodyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SegmentVideoBodyResponse) GetBody() *SegmentVideoBodyResponseBody {
	return s.Body
}

func (s *SegmentVideoBodyResponse) SetHeaders(v map[string]*string) *SegmentVideoBodyResponse {
	s.Headers = v
	return s
}

func (s *SegmentVideoBodyResponse) SetStatusCode(v int32) *SegmentVideoBodyResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentVideoBodyResponse) SetBody(v *SegmentVideoBodyResponseBody) *SegmentVideoBodyResponse {
	s.Body = v
	return s
}

func (s *SegmentVideoBodyResponse) Validate() error {
	return dara.Validate(s)
}
