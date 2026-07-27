// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgQueryRowDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgQueryRowDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgQueryRowDetailResponse
	GetStatusCode() *int32
	SetBody(v *DsgQueryRowDetailResponseBody) *DsgQueryRowDetailResponse
	GetBody() *DsgQueryRowDetailResponseBody
}

type DsgQueryRowDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgQueryRowDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgQueryRowDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgQueryRowDetailResponse) GoString() string {
	return s.String()
}

func (s *DsgQueryRowDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgQueryRowDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgQueryRowDetailResponse) GetBody() *DsgQueryRowDetailResponseBody {
	return s.Body
}

func (s *DsgQueryRowDetailResponse) SetHeaders(v map[string]*string) *DsgQueryRowDetailResponse {
	s.Headers = v
	return s
}

func (s *DsgQueryRowDetailResponse) SetStatusCode(v int32) *DsgQueryRowDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgQueryRowDetailResponse) SetBody(v *DsgQueryRowDetailResponseBody) *DsgQueryRowDetailResponse {
	s.Body = v
	return s
}

func (s *DsgQueryRowDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
