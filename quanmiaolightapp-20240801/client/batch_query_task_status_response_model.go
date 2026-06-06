// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryTaskStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchQueryTaskStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchQueryTaskStatusResponse
	GetStatusCode() *int32
	SetBody(v *BatchQueryTaskStatusResponseBody) *BatchQueryTaskStatusResponse
	GetBody() *BatchQueryTaskStatusResponseBody
}

type BatchQueryTaskStatusResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryTaskStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryTaskStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchQueryTaskStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchQueryTaskStatusResponse) GetBody() *BatchQueryTaskStatusResponseBody {
	return s.Body
}

func (s *BatchQueryTaskStatusResponse) SetHeaders(v map[string]*string) *BatchQueryTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryTaskStatusResponse) SetStatusCode(v int32) *BatchQueryTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryTaskStatusResponse) SetBody(v *BatchQueryTaskStatusResponseBody) *BatchQueryTaskStatusResponse {
	s.Body = v
	return s
}

func (s *BatchQueryTaskStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
