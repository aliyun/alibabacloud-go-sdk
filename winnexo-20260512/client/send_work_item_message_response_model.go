// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSendWorkItemMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendWorkItemMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendWorkItemMessageResponse
	GetStatusCode() *int32
	SetBody(v *SendWorkItemMessageResponseBody) *SendWorkItemMessageResponse
	GetBody() *SendWorkItemMessageResponseBody
}

type SendWorkItemMessageResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendWorkItemMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendWorkItemMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s SendWorkItemMessageResponse) GoString() string {
	return s.String()
}

func (s *SendWorkItemMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendWorkItemMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendWorkItemMessageResponse) GetBody() *SendWorkItemMessageResponseBody {
	return s.Body
}

func (s *SendWorkItemMessageResponse) SetHeaders(v map[string]*string) *SendWorkItemMessageResponse {
	s.Headers = v
	return s
}

func (s *SendWorkItemMessageResponse) SetStatusCode(v int32) *SendWorkItemMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *SendWorkItemMessageResponse) SetBody(v *SendWorkItemMessageResponseBody) *SendWorkItemMessageResponse {
	s.Body = v
	return s
}

func (s *SendWorkItemMessageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
