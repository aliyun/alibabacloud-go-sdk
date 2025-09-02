// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRemindResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRemindResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRemindResponse
	GetStatusCode() *int32
	SetBody(v *GetRemindResponseBody) *GetRemindResponse
	GetBody() *GetRemindResponseBody
}

type GetRemindResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRemindResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRemindResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRemindResponse) GoString() string {
	return s.String()
}

func (s *GetRemindResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRemindResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRemindResponse) GetBody() *GetRemindResponseBody {
	return s.Body
}

func (s *GetRemindResponse) SetHeaders(v map[string]*string) *GetRemindResponse {
	s.Headers = v
	return s
}

func (s *GetRemindResponse) SetStatusCode(v int32) *GetRemindResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRemindResponse) SetBody(v *GetRemindResponseBody) *GetRemindResponse {
	s.Body = v
	return s
}

func (s *GetRemindResponse) Validate() error {
	return dara.Validate(s)
}
