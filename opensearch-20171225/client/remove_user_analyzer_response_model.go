// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveUserAnalyzerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveUserAnalyzerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveUserAnalyzerResponse
	GetStatusCode() *int32
	SetBody(v *RemoveUserAnalyzerResponseBody) *RemoveUserAnalyzerResponse
	GetBody() *RemoveUserAnalyzerResponseBody
}

type RemoveUserAnalyzerResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveUserAnalyzerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveUserAnalyzerResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveUserAnalyzerResponse) GoString() string {
	return s.String()
}

func (s *RemoveUserAnalyzerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveUserAnalyzerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveUserAnalyzerResponse) GetBody() *RemoveUserAnalyzerResponseBody {
	return s.Body
}

func (s *RemoveUserAnalyzerResponse) SetHeaders(v map[string]*string) *RemoveUserAnalyzerResponse {
	s.Headers = v
	return s
}

func (s *RemoveUserAnalyzerResponse) SetStatusCode(v int32) *RemoveUserAnalyzerResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveUserAnalyzerResponse) SetBody(v *RemoveUserAnalyzerResponseBody) *RemoveUserAnalyzerResponse {
	s.Body = v
	return s
}

func (s *RemoveUserAnalyzerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
