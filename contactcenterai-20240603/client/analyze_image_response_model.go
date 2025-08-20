// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAnalyzeImageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AnalyzeImageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AnalyzeImageResponse
	GetStatusCode() *int32
	SetBody(v *AnalyzeImageResponseBody) *AnalyzeImageResponse
	GetBody() *AnalyzeImageResponseBody
}

type AnalyzeImageResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AnalyzeImageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AnalyzeImageResponse) String() string {
	return dara.Prettify(s)
}

func (s AnalyzeImageResponse) GoString() string {
	return s.String()
}

func (s *AnalyzeImageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AnalyzeImageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AnalyzeImageResponse) GetBody() *AnalyzeImageResponseBody {
	return s.Body
}

func (s *AnalyzeImageResponse) SetHeaders(v map[string]*string) *AnalyzeImageResponse {
	s.Headers = v
	return s
}

func (s *AnalyzeImageResponse) SetStatusCode(v int32) *AnalyzeImageResponse {
	s.StatusCode = &v
	return s
}

func (s *AnalyzeImageResponse) SetBody(v *AnalyzeImageResponseBody) *AnalyzeImageResponse {
	s.Body = v
	return s
}

func (s *AnalyzeImageResponse) Validate() error {
	return dara.Validate(s)
}
