// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNotifyMeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNotifyMeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNotifyMeResponse
	GetStatusCode() *int32
	SetBody(v *GetNotifyMeResponseBody) *GetNotifyMeResponse
	GetBody() *GetNotifyMeResponseBody
}

type GetNotifyMeResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNotifyMeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNotifyMeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNotifyMeResponse) GoString() string {
	return s.String()
}

func (s *GetNotifyMeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNotifyMeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNotifyMeResponse) GetBody() *GetNotifyMeResponseBody {
	return s.Body
}

func (s *GetNotifyMeResponse) SetHeaders(v map[string]*string) *GetNotifyMeResponse {
	s.Headers = v
	return s
}

func (s *GetNotifyMeResponse) SetStatusCode(v int32) *GetNotifyMeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNotifyMeResponse) SetBody(v *GetNotifyMeResponseBody) *GetNotifyMeResponse {
	s.Body = v
	return s
}

func (s *GetNotifyMeResponse) Validate() error {
	return dara.Validate(s)
}
