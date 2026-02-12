// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsTraceQueryByMsgKeyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsTraceQueryByMsgKeyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsTraceQueryByMsgKeyResponse
	GetStatusCode() *int32
	SetBody(v *OnsTraceQueryByMsgKeyResponseBody) *OnsTraceQueryByMsgKeyResponse
	GetBody() *OnsTraceQueryByMsgKeyResponseBody
}

type OnsTraceQueryByMsgKeyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsTraceQueryByMsgKeyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsTraceQueryByMsgKeyResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsTraceQueryByMsgKeyResponse) GoString() string {
	return s.String()
}

func (s *OnsTraceQueryByMsgKeyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsTraceQueryByMsgKeyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsTraceQueryByMsgKeyResponse) GetBody() *OnsTraceQueryByMsgKeyResponseBody {
	return s.Body
}

func (s *OnsTraceQueryByMsgKeyResponse) SetHeaders(v map[string]*string) *OnsTraceQueryByMsgKeyResponse {
	s.Headers = v
	return s
}

func (s *OnsTraceQueryByMsgKeyResponse) SetStatusCode(v int32) *OnsTraceQueryByMsgKeyResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsTraceQueryByMsgKeyResponse) SetBody(v *OnsTraceQueryByMsgKeyResponseBody) *OnsTraceQueryByMsgKeyResponse {
	s.Body = v
	return s
}

func (s *OnsTraceQueryByMsgKeyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
