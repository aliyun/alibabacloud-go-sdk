// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAICaptionExtractionJobsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAICaptionExtractionJobsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAICaptionExtractionJobsResponse
	GetStatusCode() *int32
	SetBody(v *GetAICaptionExtractionJobsResponseBody) *GetAICaptionExtractionJobsResponse
	GetBody() *GetAICaptionExtractionJobsResponseBody
}

type GetAICaptionExtractionJobsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAICaptionExtractionJobsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAICaptionExtractionJobsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAICaptionExtractionJobsResponse) GoString() string {
	return s.String()
}

func (s *GetAICaptionExtractionJobsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAICaptionExtractionJobsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAICaptionExtractionJobsResponse) GetBody() *GetAICaptionExtractionJobsResponseBody {
	return s.Body
}

func (s *GetAICaptionExtractionJobsResponse) SetHeaders(v map[string]*string) *GetAICaptionExtractionJobsResponse {
	s.Headers = v
	return s
}

func (s *GetAICaptionExtractionJobsResponse) SetStatusCode(v int32) *GetAICaptionExtractionJobsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAICaptionExtractionJobsResponse) SetBody(v *GetAICaptionExtractionJobsResponseBody) *GetAICaptionExtractionJobsResponse {
	s.Body = v
	return s
}

func (s *GetAICaptionExtractionJobsResponse) Validate() error {
	return dara.Validate(s)
}
