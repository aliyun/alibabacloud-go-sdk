// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentClothResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SegmentClothResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SegmentClothResponse
	GetStatusCode() *int32
	SetBody(v *SegmentClothResponseBody) *SegmentClothResponse
	GetBody() *SegmentClothResponseBody
}

type SegmentClothResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentClothResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentClothResponse) String() string {
	return dara.Prettify(s)
}

func (s SegmentClothResponse) GoString() string {
	return s.String()
}

func (s *SegmentClothResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SegmentClothResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SegmentClothResponse) GetBody() *SegmentClothResponseBody {
	return s.Body
}

func (s *SegmentClothResponse) SetHeaders(v map[string]*string) *SegmentClothResponse {
	s.Headers = v
	return s
}

func (s *SegmentClothResponse) SetStatusCode(v int32) *SegmentClothResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentClothResponse) SetBody(v *SegmentClothResponseBody) *SegmentClothResponse {
	s.Body = v
	return s
}

func (s *SegmentClothResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
