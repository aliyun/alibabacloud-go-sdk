// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryCreateAccountTraceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchQueryCreateAccountTraceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchQueryCreateAccountTraceResponse
	GetStatusCode() *int32
	SetBody(v *BatchQueryCreateAccountTraceResponseBody) *BatchQueryCreateAccountTraceResponse
	GetBody() *BatchQueryCreateAccountTraceResponseBody
}

type BatchQueryCreateAccountTraceResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryCreateAccountTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryCreateAccountTraceResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryCreateAccountTraceResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryCreateAccountTraceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchQueryCreateAccountTraceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchQueryCreateAccountTraceResponse) GetBody() *BatchQueryCreateAccountTraceResponseBody {
	return s.Body
}

func (s *BatchQueryCreateAccountTraceResponse) SetHeaders(v map[string]*string) *BatchQueryCreateAccountTraceResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryCreateAccountTraceResponse) SetStatusCode(v int32) *BatchQueryCreateAccountTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryCreateAccountTraceResponse) SetBody(v *BatchQueryCreateAccountTraceResponseBody) *BatchQueryCreateAccountTraceResponse {
	s.Body = v
	return s
}

func (s *BatchQueryCreateAccountTraceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
