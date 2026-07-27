// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDsgGetVisitDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DsgGetVisitDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DsgGetVisitDetailResponse
	GetStatusCode() *int32
	SetBody(v *DsgGetVisitDetailResponseBody) *DsgGetVisitDetailResponse
	GetBody() *DsgGetVisitDetailResponseBody
}

type DsgGetVisitDetailResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DsgGetVisitDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DsgGetVisitDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s DsgGetVisitDetailResponse) GoString() string {
	return s.String()
}

func (s *DsgGetVisitDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DsgGetVisitDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DsgGetVisitDetailResponse) GetBody() *DsgGetVisitDetailResponseBody {
	return s.Body
}

func (s *DsgGetVisitDetailResponse) SetHeaders(v map[string]*string) *DsgGetVisitDetailResponse {
	s.Headers = v
	return s
}

func (s *DsgGetVisitDetailResponse) SetStatusCode(v int32) *DsgGetVisitDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *DsgGetVisitDetailResponse) SetBody(v *DsgGetVisitDetailResponseBody) *DsgGetVisitDetailResponse {
	s.Body = v
	return s
}

func (s *DsgGetVisitDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
