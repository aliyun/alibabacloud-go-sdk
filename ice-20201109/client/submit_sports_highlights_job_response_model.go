// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSportsHighlightsJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSportsHighlightsJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSportsHighlightsJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSportsHighlightsJobResponseBody) *SubmitSportsHighlightsJobResponse
	GetBody() *SubmitSportsHighlightsJobResponseBody
}

type SubmitSportsHighlightsJobResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSportsHighlightsJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSportsHighlightsJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSportsHighlightsJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitSportsHighlightsJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSportsHighlightsJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSportsHighlightsJobResponse) GetBody() *SubmitSportsHighlightsJobResponseBody {
	return s.Body
}

func (s *SubmitSportsHighlightsJobResponse) SetHeaders(v map[string]*string) *SubmitSportsHighlightsJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitSportsHighlightsJobResponse) SetStatusCode(v int32) *SubmitSportsHighlightsJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSportsHighlightsJobResponse) SetBody(v *SubmitSportsHighlightsJobResponseBody) *SubmitSportsHighlightsJobResponse {
	s.Body = v
	return s
}

func (s *SubmitSportsHighlightsJobResponse) Validate() error {
	return dara.Validate(s)
}
