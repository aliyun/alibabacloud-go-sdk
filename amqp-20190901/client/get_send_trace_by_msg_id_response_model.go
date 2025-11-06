// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSendTraceByMsgIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSendTraceByMsgIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSendTraceByMsgIdResponse
	GetStatusCode() *int32
	SetBody(v *GetSendTraceByMsgIdResponseBody) *GetSendTraceByMsgIdResponse
	GetBody() *GetSendTraceByMsgIdResponseBody
}

type GetSendTraceByMsgIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSendTraceByMsgIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSendTraceByMsgIdResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSendTraceByMsgIdResponse) GoString() string {
	return s.String()
}

func (s *GetSendTraceByMsgIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSendTraceByMsgIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSendTraceByMsgIdResponse) GetBody() *GetSendTraceByMsgIdResponseBody {
	return s.Body
}

func (s *GetSendTraceByMsgIdResponse) SetHeaders(v map[string]*string) *GetSendTraceByMsgIdResponse {
	s.Headers = v
	return s
}

func (s *GetSendTraceByMsgIdResponse) SetStatusCode(v int32) *GetSendTraceByMsgIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSendTraceByMsgIdResponse) SetBody(v *GetSendTraceByMsgIdResponseBody) *GetSendTraceByMsgIdResponse {
	s.Body = v
	return s
}

func (s *GetSendTraceByMsgIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
