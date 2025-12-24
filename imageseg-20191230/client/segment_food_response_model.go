// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentFoodResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SegmentFoodResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SegmentFoodResponse
	GetStatusCode() *int32
	SetBody(v *SegmentFoodResponseBody) *SegmentFoodResponse
	GetBody() *SegmentFoodResponseBody
}

type SegmentFoodResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentFoodResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentFoodResponse) String() string {
	return dara.Prettify(s)
}

func (s SegmentFoodResponse) GoString() string {
	return s.String()
}

func (s *SegmentFoodResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SegmentFoodResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SegmentFoodResponse) GetBody() *SegmentFoodResponseBody {
	return s.Body
}

func (s *SegmentFoodResponse) SetHeaders(v map[string]*string) *SegmentFoodResponse {
	s.Headers = v
	return s
}

func (s *SegmentFoodResponse) SetStatusCode(v int32) *SegmentFoodResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentFoodResponse) SetBody(v *SegmentFoodResponseBody) *SegmentFoodResponse {
	s.Body = v
	return s
}

func (s *SegmentFoodResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
