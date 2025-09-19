// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProcessSoarStrategyTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ProcessSoarStrategyTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ProcessSoarStrategyTaskResponse
	GetStatusCode() *int32
	SetBody(v *ProcessSoarStrategyTaskResponseBody) *ProcessSoarStrategyTaskResponse
	GetBody() *ProcessSoarStrategyTaskResponseBody
}

type ProcessSoarStrategyTaskResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ProcessSoarStrategyTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ProcessSoarStrategyTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ProcessSoarStrategyTaskResponse) GoString() string {
	return s.String()
}

func (s *ProcessSoarStrategyTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ProcessSoarStrategyTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ProcessSoarStrategyTaskResponse) GetBody() *ProcessSoarStrategyTaskResponseBody {
	return s.Body
}

func (s *ProcessSoarStrategyTaskResponse) SetHeaders(v map[string]*string) *ProcessSoarStrategyTaskResponse {
	s.Headers = v
	return s
}

func (s *ProcessSoarStrategyTaskResponse) SetStatusCode(v int32) *ProcessSoarStrategyTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ProcessSoarStrategyTaskResponse) SetBody(v *ProcessSoarStrategyTaskResponseBody) *ProcessSoarStrategyTaskResponse {
	s.Body = v
	return s
}

func (s *ProcessSoarStrategyTaskResponse) Validate() error {
	return dara.Validate(s)
}
