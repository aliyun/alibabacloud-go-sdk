// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iParseBatchTaskDependencyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ParseBatchTaskDependencyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ParseBatchTaskDependencyResponse
	GetStatusCode() *int32
	SetBody(v *ParseBatchTaskDependencyResponseBody) *ParseBatchTaskDependencyResponse
	GetBody() *ParseBatchTaskDependencyResponseBody
}

type ParseBatchTaskDependencyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ParseBatchTaskDependencyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ParseBatchTaskDependencyResponse) String() string {
	return dara.Prettify(s)
}

func (s ParseBatchTaskDependencyResponse) GoString() string {
	return s.String()
}

func (s *ParseBatchTaskDependencyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ParseBatchTaskDependencyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ParseBatchTaskDependencyResponse) GetBody() *ParseBatchTaskDependencyResponseBody {
	return s.Body
}

func (s *ParseBatchTaskDependencyResponse) SetHeaders(v map[string]*string) *ParseBatchTaskDependencyResponse {
	s.Headers = v
	return s
}

func (s *ParseBatchTaskDependencyResponse) SetStatusCode(v int32) *ParseBatchTaskDependencyResponse {
	s.StatusCode = &v
	return s
}

func (s *ParseBatchTaskDependencyResponse) SetBody(v *ParseBatchTaskDependencyResponseBody) *ParseBatchTaskDependencyResponse {
	s.Body = v
	return s
}

func (s *ParseBatchTaskDependencyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
