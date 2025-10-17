// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnalysisTagDetailByTaskIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAnalysisTagDetailByTaskIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAnalysisTagDetailByTaskIdResponse
	GetStatusCode() *int32
	SetBody(v *ListAnalysisTagDetailByTaskIdResponseBody) *ListAnalysisTagDetailByTaskIdResponse
	GetBody() *ListAnalysisTagDetailByTaskIdResponseBody
}

type ListAnalysisTagDetailByTaskIdResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnalysisTagDetailByTaskIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnalysisTagDetailByTaskIdResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisTagDetailByTaskIdResponse) GoString() string {
	return s.String()
}

func (s *ListAnalysisTagDetailByTaskIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAnalysisTagDetailByTaskIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAnalysisTagDetailByTaskIdResponse) GetBody() *ListAnalysisTagDetailByTaskIdResponseBody {
	return s.Body
}

func (s *ListAnalysisTagDetailByTaskIdResponse) SetHeaders(v map[string]*string) *ListAnalysisTagDetailByTaskIdResponse {
	s.Headers = v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponse) SetStatusCode(v int32) *ListAnalysisTagDetailByTaskIdResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponse) SetBody(v *ListAnalysisTagDetailByTaskIdResponseBody) *ListAnalysisTagDetailByTaskIdResponse {
	s.Body = v
	return s
}

func (s *ListAnalysisTagDetailByTaskIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
