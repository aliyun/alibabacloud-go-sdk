// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelExecutionResponse
	GetStatusCode() *int32
	SetBody(v *CancelExecutionResponseBody) *CancelExecutionResponse
	GetBody() *CancelExecutionResponseBody
}

type CancelExecutionResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelExecutionResponse) GoString() string {
	return s.String()
}

func (s *CancelExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelExecutionResponse) GetBody() *CancelExecutionResponseBody {
	return s.Body
}

func (s *CancelExecutionResponse) SetHeaders(v map[string]*string) *CancelExecutionResponse {
	s.Headers = v
	return s
}

func (s *CancelExecutionResponse) SetStatusCode(v int32) *CancelExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelExecutionResponse) SetBody(v *CancelExecutionResponseBody) *CancelExecutionResponse {
	s.Body = v
	return s
}

func (s *CancelExecutionResponse) Validate() error {
	return dara.Validate(s)
}
