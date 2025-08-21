// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushMessageToiOSResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushMessageToiOSResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushMessageToiOSResponse
	GetStatusCode() *int32
	SetBody(v *PushMessageToiOSResponseBody) *PushMessageToiOSResponse
	GetBody() *PushMessageToiOSResponseBody
}

type PushMessageToiOSResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushMessageToiOSResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushMessageToiOSResponse) String() string {
	return dara.Prettify(s)
}

func (s PushMessageToiOSResponse) GoString() string {
	return s.String()
}

func (s *PushMessageToiOSResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushMessageToiOSResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushMessageToiOSResponse) GetBody() *PushMessageToiOSResponseBody {
	return s.Body
}

func (s *PushMessageToiOSResponse) SetHeaders(v map[string]*string) *PushMessageToiOSResponse {
	s.Headers = v
	return s
}

func (s *PushMessageToiOSResponse) SetStatusCode(v int32) *PushMessageToiOSResponse {
	s.StatusCode = &v
	return s
}

func (s *PushMessageToiOSResponse) SetBody(v *PushMessageToiOSResponseBody) *PushMessageToiOSResponse {
	s.Body = v
	return s
}

func (s *PushMessageToiOSResponse) Validate() error {
	return dara.Validate(s)
}
