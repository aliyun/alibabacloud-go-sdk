// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnsMessageGetByMsgIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnsMessageGetByMsgIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnsMessageGetByMsgIdResponse
	GetStatusCode() *int32
	SetBody(v *OnsMessageGetByMsgIdResponseBody) *OnsMessageGetByMsgIdResponse
	GetBody() *OnsMessageGetByMsgIdResponseBody
}

type OnsMessageGetByMsgIdResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnsMessageGetByMsgIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnsMessageGetByMsgIdResponse) String() string {
	return dara.Prettify(s)
}

func (s OnsMessageGetByMsgIdResponse) GoString() string {
	return s.String()
}

func (s *OnsMessageGetByMsgIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnsMessageGetByMsgIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnsMessageGetByMsgIdResponse) GetBody() *OnsMessageGetByMsgIdResponseBody {
	return s.Body
}

func (s *OnsMessageGetByMsgIdResponse) SetHeaders(v map[string]*string) *OnsMessageGetByMsgIdResponse {
	s.Headers = v
	return s
}

func (s *OnsMessageGetByMsgIdResponse) SetStatusCode(v int32) *OnsMessageGetByMsgIdResponse {
	s.StatusCode = &v
	return s
}

func (s *OnsMessageGetByMsgIdResponse) SetBody(v *OnsMessageGetByMsgIdResponseBody) *OnsMessageGetByMsgIdResponse {
	s.Body = v
	return s
}

func (s *OnsMessageGetByMsgIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
