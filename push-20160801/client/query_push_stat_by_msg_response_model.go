// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPushStatByMsgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPushStatByMsgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPushStatByMsgResponse
	GetStatusCode() *int32
	SetBody(v *QueryPushStatByMsgResponseBody) *QueryPushStatByMsgResponse
	GetBody() *QueryPushStatByMsgResponseBody
}

type QueryPushStatByMsgResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPushStatByMsgResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPushStatByMsgResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPushStatByMsgResponse) GoString() string {
	return s.String()
}

func (s *QueryPushStatByMsgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPushStatByMsgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPushStatByMsgResponse) GetBody() *QueryPushStatByMsgResponseBody {
	return s.Body
}

func (s *QueryPushStatByMsgResponse) SetHeaders(v map[string]*string) *QueryPushStatByMsgResponse {
	s.Headers = v
	return s
}

func (s *QueryPushStatByMsgResponse) SetStatusCode(v int32) *QueryPushStatByMsgResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPushStatByMsgResponse) SetBody(v *QueryPushStatByMsgResponseBody) *QueryPushStatByMsgResponse {
	s.Body = v
	return s
}

func (s *QueryPushStatByMsgResponse) Validate() error {
	return dara.Validate(s)
}
