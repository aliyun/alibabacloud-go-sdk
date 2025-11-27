// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckDuckDBDependencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PrecheckDuckDBDependencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PrecheckDuckDBDependencyResponse
	GetStatusCode() *int32
	SetBody(v *PrecheckDuckDBDependencyResponseBody) *PrecheckDuckDBDependencyResponse
	GetBody() *PrecheckDuckDBDependencyResponseBody
}

type PrecheckDuckDBDependencyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PrecheckDuckDBDependencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PrecheckDuckDBDependencyResponse) String() string {
	return dara.Prettify(s)
}

func (s PrecheckDuckDBDependencyResponse) GoString() string {
	return s.String()
}

func (s *PrecheckDuckDBDependencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PrecheckDuckDBDependencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PrecheckDuckDBDependencyResponse) GetBody() *PrecheckDuckDBDependencyResponseBody {
	return s.Body
}

func (s *PrecheckDuckDBDependencyResponse) SetHeaders(v map[string]*string) *PrecheckDuckDBDependencyResponse {
	s.Headers = v
	return s
}

func (s *PrecheckDuckDBDependencyResponse) SetStatusCode(v int32) *PrecheckDuckDBDependencyResponse {
	s.StatusCode = &v
	return s
}

func (s *PrecheckDuckDBDependencyResponse) SetBody(v *PrecheckDuckDBDependencyResponseBody) *PrecheckDuckDBDependencyResponse {
	s.Body = v
	return s
}

func (s *PrecheckDuckDBDependencyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
