// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAICaptionExtractionJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitAICaptionExtractionJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitAICaptionExtractionJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitAICaptionExtractionJobResponseBody) *SubmitAICaptionExtractionJobResponse
	GetBody() *SubmitAICaptionExtractionJobResponseBody
}

type SubmitAICaptionExtractionJobResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitAICaptionExtractionJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitAICaptionExtractionJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitAICaptionExtractionJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitAICaptionExtractionJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitAICaptionExtractionJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitAICaptionExtractionJobResponse) GetBody() *SubmitAICaptionExtractionJobResponseBody {
	return s.Body
}

func (s *SubmitAICaptionExtractionJobResponse) SetHeaders(v map[string]*string) *SubmitAICaptionExtractionJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitAICaptionExtractionJobResponse) SetStatusCode(v int32) *SubmitAICaptionExtractionJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitAICaptionExtractionJobResponse) SetBody(v *SubmitAICaptionExtractionJobResponseBody) *SubmitAICaptionExtractionJobResponse {
	s.Body = v
	return s
}

func (s *SubmitAICaptionExtractionJobResponse) Validate() error {
	return dara.Validate(s)
}
