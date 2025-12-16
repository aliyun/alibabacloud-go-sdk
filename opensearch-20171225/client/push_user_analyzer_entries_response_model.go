// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushUserAnalyzerEntriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushUserAnalyzerEntriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushUserAnalyzerEntriesResponse
	GetStatusCode() *int32
	SetBody(v *PushUserAnalyzerEntriesResponseBody) *PushUserAnalyzerEntriesResponse
	GetBody() *PushUserAnalyzerEntriesResponseBody
}

type PushUserAnalyzerEntriesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushUserAnalyzerEntriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushUserAnalyzerEntriesResponse) String() string {
	return dara.Prettify(s)
}

func (s PushUserAnalyzerEntriesResponse) GoString() string {
	return s.String()
}

func (s *PushUserAnalyzerEntriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushUserAnalyzerEntriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushUserAnalyzerEntriesResponse) GetBody() *PushUserAnalyzerEntriesResponseBody {
	return s.Body
}

func (s *PushUserAnalyzerEntriesResponse) SetHeaders(v map[string]*string) *PushUserAnalyzerEntriesResponse {
	s.Headers = v
	return s
}

func (s *PushUserAnalyzerEntriesResponse) SetStatusCode(v int32) *PushUserAnalyzerEntriesResponse {
	s.StatusCode = &v
	return s
}

func (s *PushUserAnalyzerEntriesResponse) SetBody(v *PushUserAnalyzerEntriesResponseBody) *PushUserAnalyzerEntriesResponse {
	s.Body = v
	return s
}

func (s *PushUserAnalyzerEntriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
