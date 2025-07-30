// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTensorboardResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopTensorboardResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopTensorboardResponse
	GetStatusCode() *int32
	SetBody(v *StopTensorboardResponseBody) *StopTensorboardResponse
	GetBody() *StopTensorboardResponseBody
}

type StopTensorboardResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopTensorboardResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopTensorboardResponse) String() string {
	return dara.Prettify(s)
}

func (s StopTensorboardResponse) GoString() string {
	return s.String()
}

func (s *StopTensorboardResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopTensorboardResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopTensorboardResponse) GetBody() *StopTensorboardResponseBody {
	return s.Body
}

func (s *StopTensorboardResponse) SetHeaders(v map[string]*string) *StopTensorboardResponse {
	s.Headers = v
	return s
}

func (s *StopTensorboardResponse) SetStatusCode(v int32) *StopTensorboardResponse {
	s.StatusCode = &v
	return s
}

func (s *StopTensorboardResponse) SetBody(v *StopTensorboardResponseBody) *StopTensorboardResponse {
	s.Body = v
	return s
}

func (s *StopTensorboardResponse) Validate() error {
	return dara.Validate(s)
}
