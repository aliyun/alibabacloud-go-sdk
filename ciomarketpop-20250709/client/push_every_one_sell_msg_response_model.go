// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushEveryOneSellMsgResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushEveryOneSellMsgResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushEveryOneSellMsgResponse
	GetStatusCode() *int32
	SetBody(v string) *PushEveryOneSellMsgResponse
	GetBody() *string
}

type PushEveryOneSellMsgResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *string            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushEveryOneSellMsgResponse) String() string {
	return dara.Prettify(s)
}

func (s PushEveryOneSellMsgResponse) GoString() string {
	return s.String()
}

func (s *PushEveryOneSellMsgResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushEveryOneSellMsgResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushEveryOneSellMsgResponse) GetBody() *string {
	return s.Body
}

func (s *PushEveryOneSellMsgResponse) SetHeaders(v map[string]*string) *PushEveryOneSellMsgResponse {
	s.Headers = v
	return s
}

func (s *PushEveryOneSellMsgResponse) SetStatusCode(v int32) *PushEveryOneSellMsgResponse {
	s.StatusCode = &v
	return s
}

func (s *PushEveryOneSellMsgResponse) SetBody(v string) *PushEveryOneSellMsgResponse {
	s.Body = &v
	return s
}

func (s *PushEveryOneSellMsgResponse) Validate() error {
	return dara.Validate(s)
}
