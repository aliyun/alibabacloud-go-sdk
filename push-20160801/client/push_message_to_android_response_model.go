// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMessageToAndroidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushMessageToAndroidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushMessageToAndroidResponse
	GetStatusCode() *int32
	SetBody(v *PushMessageToAndroidResponseBody) *PushMessageToAndroidResponse
	GetBody() *PushMessageToAndroidResponseBody
}

type PushMessageToAndroidResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushMessageToAndroidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushMessageToAndroidResponse) String() string {
	return dara.Prettify(s)
}

func (s PushMessageToAndroidResponse) GoString() string {
	return s.String()
}

func (s *PushMessageToAndroidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushMessageToAndroidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushMessageToAndroidResponse) GetBody() *PushMessageToAndroidResponseBody {
	return s.Body
}

func (s *PushMessageToAndroidResponse) SetHeaders(v map[string]*string) *PushMessageToAndroidResponse {
	s.Headers = v
	return s
}

func (s *PushMessageToAndroidResponse) SetStatusCode(v int32) *PushMessageToAndroidResponse {
	s.StatusCode = &v
	return s
}

func (s *PushMessageToAndroidResponse) SetBody(v *PushMessageToAndroidResponseBody) *PushMessageToAndroidResponse {
	s.Body = v
	return s
}

func (s *PushMessageToAndroidResponse) Validate() error {
	return dara.Validate(s)
}
