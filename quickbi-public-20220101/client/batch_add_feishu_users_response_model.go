// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchAddFeishuUsersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BatchAddFeishuUsersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BatchAddFeishuUsersResponse
	GetStatusCode() *int32
	SetBody(v *BatchAddFeishuUsersResponseBody) *BatchAddFeishuUsersResponse
	GetBody() *BatchAddFeishuUsersResponseBody
}

type BatchAddFeishuUsersResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchAddFeishuUsersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchAddFeishuUsersResponse) String() string {
	return dara.Prettify(s)
}

func (s BatchAddFeishuUsersResponse) GoString() string {
	return s.String()
}

func (s *BatchAddFeishuUsersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BatchAddFeishuUsersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BatchAddFeishuUsersResponse) GetBody() *BatchAddFeishuUsersResponseBody {
	return s.Body
}

func (s *BatchAddFeishuUsersResponse) SetHeaders(v map[string]*string) *BatchAddFeishuUsersResponse {
	s.Headers = v
	return s
}

func (s *BatchAddFeishuUsersResponse) SetStatusCode(v int32) *BatchAddFeishuUsersResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchAddFeishuUsersResponse) SetBody(v *BatchAddFeishuUsersResponseBody) *BatchAddFeishuUsersResponse {
	s.Body = v
	return s
}

func (s *BatchAddFeishuUsersResponse) Validate() error {
	return dara.Validate(s)
}
