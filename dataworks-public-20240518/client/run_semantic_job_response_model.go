// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSemanticJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunSemanticJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunSemanticJobResponse
	GetStatusCode() *int32
	SetBody(v *RunSemanticJobResponseBody) *RunSemanticJobResponse
	GetBody() *RunSemanticJobResponseBody
}

type RunSemanticJobResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunSemanticJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunSemanticJobResponse) String() string {
	return dara.Prettify(s)
}

func (s RunSemanticJobResponse) GoString() string {
	return s.String()
}

func (s *RunSemanticJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunSemanticJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunSemanticJobResponse) GetBody() *RunSemanticJobResponseBody {
	return s.Body
}

func (s *RunSemanticJobResponse) SetHeaders(v map[string]*string) *RunSemanticJobResponse {
	s.Headers = v
	return s
}

func (s *RunSemanticJobResponse) SetStatusCode(v int32) *RunSemanticJobResponse {
	s.StatusCode = &v
	return s
}

func (s *RunSemanticJobResponse) SetBody(v *RunSemanticJobResponseBody) *RunSemanticJobResponse {
	s.Body = v
	return s
}

func (s *RunSemanticJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
