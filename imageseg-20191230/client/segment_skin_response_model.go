// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSegmentSkinResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SegmentSkinResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SegmentSkinResponse
	GetStatusCode() *int32
	SetBody(v *SegmentSkinResponseBody) *SegmentSkinResponse
	GetBody() *SegmentSkinResponseBody
}

type SegmentSkinResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SegmentSkinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SegmentSkinResponse) String() string {
	return dara.Prettify(s)
}

func (s SegmentSkinResponse) GoString() string {
	return s.String()
}

func (s *SegmentSkinResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SegmentSkinResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SegmentSkinResponse) GetBody() *SegmentSkinResponseBody {
	return s.Body
}

func (s *SegmentSkinResponse) SetHeaders(v map[string]*string) *SegmentSkinResponse {
	s.Headers = v
	return s
}

func (s *SegmentSkinResponse) SetStatusCode(v int32) *SegmentSkinResponse {
	s.StatusCode = &v
	return s
}

func (s *SegmentSkinResponse) SetBody(v *SegmentSkinResponseBody) *SegmentSkinResponse {
	s.Body = v
	return s
}

func (s *SegmentSkinResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
