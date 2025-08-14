// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDocParserJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitDocParserJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitDocParserJobResponse
	GetStatusCode() *int32
	SetBody(v *SubmitDocParserJobResponseBody) *SubmitDocParserJobResponse
	GetBody() *SubmitDocParserJobResponseBody
}

type SubmitDocParserJobResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitDocParserJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitDocParserJobResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitDocParserJobResponse) GoString() string {
	return s.String()
}

func (s *SubmitDocParserJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitDocParserJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitDocParserJobResponse) GetBody() *SubmitDocParserJobResponseBody {
	return s.Body
}

func (s *SubmitDocParserJobResponse) SetHeaders(v map[string]*string) *SubmitDocParserJobResponse {
	s.Headers = v
	return s
}

func (s *SubmitDocParserJobResponse) SetStatusCode(v int32) *SubmitDocParserJobResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitDocParserJobResponse) SetBody(v *SubmitDocParserJobResponseBody) *SubmitDocParserJobResponse {
	s.Body = v
	return s
}

func (s *SubmitDocParserJobResponse) Validate() error {
	return dara.Validate(s)
}
