// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendDtmfSignalingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendDtmfSignalingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendDtmfSignalingResponse
	GetStatusCode() *int32
	SetBody(v *SendDtmfSignalingResponseBody) *SendDtmfSignalingResponse
	GetBody() *SendDtmfSignalingResponseBody
}

type SendDtmfSignalingResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendDtmfSignalingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendDtmfSignalingResponse) String() string {
	return dara.Prettify(s)
}

func (s SendDtmfSignalingResponse) GoString() string {
	return s.String()
}

func (s *SendDtmfSignalingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendDtmfSignalingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendDtmfSignalingResponse) GetBody() *SendDtmfSignalingResponseBody {
	return s.Body
}

func (s *SendDtmfSignalingResponse) SetHeaders(v map[string]*string) *SendDtmfSignalingResponse {
	s.Headers = v
	return s
}

func (s *SendDtmfSignalingResponse) SetStatusCode(v int32) *SendDtmfSignalingResponse {
	s.StatusCode = &v
	return s
}

func (s *SendDtmfSignalingResponse) SetBody(v *SendDtmfSignalingResponseBody) *SendDtmfSignalingResponse {
	s.Body = v
	return s
}

func (s *SendDtmfSignalingResponse) Validate() error {
	return dara.Validate(s)
}
