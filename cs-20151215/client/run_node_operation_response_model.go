// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunNodeOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunNodeOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunNodeOperationResponse
	GetStatusCode() *int32
	SetBody(v *RunNodeOperationResponseBody) *RunNodeOperationResponse
	GetBody() *RunNodeOperationResponseBody
}

type RunNodeOperationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunNodeOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunNodeOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s RunNodeOperationResponse) GoString() string {
	return s.String()
}

func (s *RunNodeOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunNodeOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunNodeOperationResponse) GetBody() *RunNodeOperationResponseBody {
	return s.Body
}

func (s *RunNodeOperationResponse) SetHeaders(v map[string]*string) *RunNodeOperationResponse {
	s.Headers = v
	return s
}

func (s *RunNodeOperationResponse) SetStatusCode(v int32) *RunNodeOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *RunNodeOperationResponse) SetBody(v *RunNodeOperationResponseBody) *RunNodeOperationResponse {
	s.Body = v
	return s
}

func (s *RunNodeOperationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
