// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSummaryJobDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SummaryJobDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SummaryJobDetailResponse
	GetStatusCode() *int32
	SetBody(v *SummaryJobDetailResponseBody) *SummaryJobDetailResponse
	GetBody() *SummaryJobDetailResponseBody
}

type SummaryJobDetailResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SummaryJobDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SummaryJobDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s SummaryJobDetailResponse) GoString() string {
	return s.String()
}

func (s *SummaryJobDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SummaryJobDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SummaryJobDetailResponse) GetBody() *SummaryJobDetailResponseBody {
	return s.Body
}

func (s *SummaryJobDetailResponse) SetHeaders(v map[string]*string) *SummaryJobDetailResponse {
	s.Headers = v
	return s
}

func (s *SummaryJobDetailResponse) SetStatusCode(v int32) *SummaryJobDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *SummaryJobDetailResponse) SetBody(v *SummaryJobDetailResponseBody) *SummaryJobDetailResponse {
	s.Body = v
	return s
}

func (s *SummaryJobDetailResponse) Validate() error {
	return dara.Validate(s)
}
