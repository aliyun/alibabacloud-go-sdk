// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPerformanceViewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPerformanceViewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPerformanceViewResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPerformanceViewResponseBody) *ModifyPerformanceViewResponse
	GetBody() *ModifyPerformanceViewResponseBody
}

type ModifyPerformanceViewResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPerformanceViewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPerformanceViewResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPerformanceViewResponse) GoString() string {
	return s.String()
}

func (s *ModifyPerformanceViewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPerformanceViewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPerformanceViewResponse) GetBody() *ModifyPerformanceViewResponseBody {
	return s.Body
}

func (s *ModifyPerformanceViewResponse) SetHeaders(v map[string]*string) *ModifyPerformanceViewResponse {
	s.Headers = v
	return s
}

func (s *ModifyPerformanceViewResponse) SetStatusCode(v int32) *ModifyPerformanceViewResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPerformanceViewResponse) SetBody(v *ModifyPerformanceViewResponseBody) *ModifyPerformanceViewResponse {
	s.Body = v
	return s
}

func (s *ModifyPerformanceViewResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
