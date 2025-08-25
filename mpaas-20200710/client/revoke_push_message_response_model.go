// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokePushMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RevokePushMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RevokePushMessageResponse
	GetStatusCode() *int32
	SetBody(v *RevokePushMessageResponseBody) *RevokePushMessageResponse
	GetBody() *RevokePushMessageResponseBody
}

type RevokePushMessageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RevokePushMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RevokePushMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s RevokePushMessageResponse) GoString() string {
	return s.String()
}

func (s *RevokePushMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RevokePushMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RevokePushMessageResponse) GetBody() *RevokePushMessageResponseBody {
	return s.Body
}

func (s *RevokePushMessageResponse) SetHeaders(v map[string]*string) *RevokePushMessageResponse {
	s.Headers = v
	return s
}

func (s *RevokePushMessageResponse) SetStatusCode(v int32) *RevokePushMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *RevokePushMessageResponse) SetBody(v *RevokePushMessageResponseBody) *RevokePushMessageResponse {
	s.Body = v
	return s
}

func (s *RevokePushMessageResponse) Validate() error {
	return dara.Validate(s)
}
