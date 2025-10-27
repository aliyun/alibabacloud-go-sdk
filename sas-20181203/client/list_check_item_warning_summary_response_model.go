// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCheckItemWarningSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCheckItemWarningSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCheckItemWarningSummaryResponse
	GetStatusCode() *int32
	SetBody(v *ListCheckItemWarningSummaryResponseBody) *ListCheckItemWarningSummaryResponse
	GetBody() *ListCheckItemWarningSummaryResponseBody
}

type ListCheckItemWarningSummaryResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCheckItemWarningSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCheckItemWarningSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCheckItemWarningSummaryResponse) GoString() string {
	return s.String()
}

func (s *ListCheckItemWarningSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCheckItemWarningSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCheckItemWarningSummaryResponse) GetBody() *ListCheckItemWarningSummaryResponseBody {
	return s.Body
}

func (s *ListCheckItemWarningSummaryResponse) SetHeaders(v map[string]*string) *ListCheckItemWarningSummaryResponse {
	s.Headers = v
	return s
}

func (s *ListCheckItemWarningSummaryResponse) SetStatusCode(v int32) *ListCheckItemWarningSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCheckItemWarningSummaryResponse) SetBody(v *ListCheckItemWarningSummaryResponseBody) *ListCheckItemWarningSummaryResponse {
	s.Body = v
	return s
}

func (s *ListCheckItemWarningSummaryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
