// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveAnalysisColumnResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveAnalysisColumnResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveAnalysisColumnResponse
	GetStatusCode() *int32
	SetBody(v *SaveAnalysisColumnResponseBody) *SaveAnalysisColumnResponse
	GetBody() *SaveAnalysisColumnResponseBody
}

type SaveAnalysisColumnResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveAnalysisColumnResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveAnalysisColumnResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveAnalysisColumnResponse) GoString() string {
	return s.String()
}

func (s *SaveAnalysisColumnResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveAnalysisColumnResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveAnalysisColumnResponse) GetBody() *SaveAnalysisColumnResponseBody {
	return s.Body
}

func (s *SaveAnalysisColumnResponse) SetHeaders(v map[string]*string) *SaveAnalysisColumnResponse {
	s.Headers = v
	return s
}

func (s *SaveAnalysisColumnResponse) SetStatusCode(v int32) *SaveAnalysisColumnResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveAnalysisColumnResponse) SetBody(v *SaveAnalysisColumnResponseBody) *SaveAnalysisColumnResponse {
	s.Body = v
	return s
}

func (s *SaveAnalysisColumnResponse) Validate() error {
	return dara.Validate(s)
}
