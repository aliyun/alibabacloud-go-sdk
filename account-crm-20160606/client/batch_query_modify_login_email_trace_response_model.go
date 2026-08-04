// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchQueryModifyLoginEmailTraceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchQueryModifyLoginEmailTraceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchQueryModifyLoginEmailTraceResponse
	GetStatusCode() *int32
	SetBody(v *BatchQueryModifyLoginEmailTraceResponseBody) *BatchQueryModifyLoginEmailTraceResponse
	GetBody() *BatchQueryModifyLoginEmailTraceResponseBody
}

type BatchQueryModifyLoginEmailTraceResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchQueryModifyLoginEmailTraceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchQueryModifyLoginEmailTraceResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchQueryModifyLoginEmailTraceResponse) GoString() string {
	return s.String()
}

func (s *BatchQueryModifyLoginEmailTraceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchQueryModifyLoginEmailTraceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchQueryModifyLoginEmailTraceResponse) GetBody() *BatchQueryModifyLoginEmailTraceResponseBody {
	return s.Body
}

func (s *BatchQueryModifyLoginEmailTraceResponse) SetHeaders(v map[string]*string) *BatchQueryModifyLoginEmailTraceResponse {
	s.Headers = v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponse) SetStatusCode(v int32) *BatchQueryModifyLoginEmailTraceResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponse) SetBody(v *BatchQueryModifyLoginEmailTraceResponseBody) *BatchQueryModifyLoginEmailTraceResponse {
	s.Body = v
	return s
}

func (s *BatchQueryModifyLoginEmailTraceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
