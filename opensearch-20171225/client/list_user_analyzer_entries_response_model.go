// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserAnalyzerEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserAnalyzerEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserAnalyzerEntriesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserAnalyzerEntriesResponseBody) *ListUserAnalyzerEntriesResponse
	GetBody() *ListUserAnalyzerEntriesResponseBody
}

type ListUserAnalyzerEntriesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserAnalyzerEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserAnalyzerEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserAnalyzerEntriesResponse) GoString() string {
	return s.String()
}

func (s *ListUserAnalyzerEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserAnalyzerEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserAnalyzerEntriesResponse) GetBody() *ListUserAnalyzerEntriesResponseBody {
	return s.Body
}

func (s *ListUserAnalyzerEntriesResponse) SetHeaders(v map[string]*string) *ListUserAnalyzerEntriesResponse {
	s.Headers = v
	return s
}

func (s *ListUserAnalyzerEntriesResponse) SetStatusCode(v int32) *ListUserAnalyzerEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserAnalyzerEntriesResponse) SetBody(v *ListUserAnalyzerEntriesResponseBody) *ListUserAnalyzerEntriesResponse {
	s.Body = v
	return s
}

func (s *ListUserAnalyzerEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
