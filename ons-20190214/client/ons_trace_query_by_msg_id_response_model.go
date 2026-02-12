// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTraceQueryByMsgIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsTraceQueryByMsgIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsTraceQueryByMsgIdResponse
	GetStatusCode() *int32
	SetBody(v *OnsTraceQueryByMsgIdResponseBody) *OnsTraceQueryByMsgIdResponse
	GetBody() *OnsTraceQueryByMsgIdResponseBody
}

type OnsTraceQueryByMsgIdResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTraceQueryByMsgIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTraceQueryByMsgIdResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceQueryByMsgIdResponse) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsTraceQueryByMsgIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsTraceQueryByMsgIdResponse) GetBody() *OnsTraceQueryByMsgIdResponseBody {
	return s.Body
}

func (s *OnsTraceQueryByMsgIdResponse) SetHeaders(v map[string]*string) *OnsTraceQueryByMsgIdResponse {
	s.Headers = v
	return s
}

func (s *OnsTraceQueryByMsgIdResponse) SetStatusCode(v int32) *OnsTraceQueryByMsgIdResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTraceQueryByMsgIdResponse) SetBody(v *OnsTraceQueryByMsgIdResponseBody) *OnsTraceQueryByMsgIdResponse {
	s.Body = v
	return s
}

func (s *OnsTraceQueryByMsgIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
