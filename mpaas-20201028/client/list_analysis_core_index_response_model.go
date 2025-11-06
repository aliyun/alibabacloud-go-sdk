// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnalysisCoreIndexResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAnalysisCoreIndexResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAnalysisCoreIndexResponse
	GetStatusCode() *int32
	SetBody(v *ListAnalysisCoreIndexResponseBody) *ListAnalysisCoreIndexResponse
	GetBody() *ListAnalysisCoreIndexResponseBody
}

type ListAnalysisCoreIndexResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnalysisCoreIndexResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnalysisCoreIndexResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAnalysisCoreIndexResponse) GoString() string {
	return s.String()
}

func (s *ListAnalysisCoreIndexResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAnalysisCoreIndexResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAnalysisCoreIndexResponse) GetBody() *ListAnalysisCoreIndexResponseBody {
	return s.Body
}

func (s *ListAnalysisCoreIndexResponse) SetHeaders(v map[string]*string) *ListAnalysisCoreIndexResponse {
	s.Headers = v
	return s
}

func (s *ListAnalysisCoreIndexResponse) SetStatusCode(v int32) *ListAnalysisCoreIndexResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnalysisCoreIndexResponse) SetBody(v *ListAnalysisCoreIndexResponseBody) *ListAnalysisCoreIndexResponse {
	s.Body = v
	return s
}

func (s *ListAnalysisCoreIndexResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
