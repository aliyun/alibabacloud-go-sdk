// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSearchCaseFullTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunSearchCaseFullTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunSearchCaseFullTextResponse
	GetStatusCode() *int32
	SetBody(v *RunSearchCaseFullTextResponseBody) *RunSearchCaseFullTextResponse
	GetBody() *RunSearchCaseFullTextResponseBody
}

type RunSearchCaseFullTextResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSearchCaseFullTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSearchCaseFullTextResponse) String() string {
	return dara.Prettify(s)
}

func (s RunSearchCaseFullTextResponse) GoString() string {
	return s.String()
}

func (s *RunSearchCaseFullTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunSearchCaseFullTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunSearchCaseFullTextResponse) GetBody() *RunSearchCaseFullTextResponseBody {
	return s.Body
}

func (s *RunSearchCaseFullTextResponse) SetHeaders(v map[string]*string) *RunSearchCaseFullTextResponse {
	s.Headers = v
	return s
}

func (s *RunSearchCaseFullTextResponse) SetStatusCode(v int32) *RunSearchCaseFullTextResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSearchCaseFullTextResponse) SetBody(v *RunSearchCaseFullTextResponseBody) *RunSearchCaseFullTextResponse {
	s.Body = v
	return s
}

func (s *RunSearchCaseFullTextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
