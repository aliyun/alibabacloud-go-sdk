// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectVideoIPCObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectVideoIPCObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectVideoIPCObjectResponse
	GetStatusCode() *int32
	SetBody(v *DetectVideoIPCObjectResponseBody) *DetectVideoIPCObjectResponse
	GetBody() *DetectVideoIPCObjectResponseBody
}

type DetectVideoIPCObjectResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectVideoIPCObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectVideoIPCObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectVideoIPCObjectResponse) GoString() string {
	return s.String()
}

func (s *DetectVideoIPCObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectVideoIPCObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectVideoIPCObjectResponse) GetBody() *DetectVideoIPCObjectResponseBody {
	return s.Body
}

func (s *DetectVideoIPCObjectResponse) SetHeaders(v map[string]*string) *DetectVideoIPCObjectResponse {
	s.Headers = v
	return s
}

func (s *DetectVideoIPCObjectResponse) SetStatusCode(v int32) *DetectVideoIPCObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectVideoIPCObjectResponse) SetBody(v *DetectVideoIPCObjectResponseBody) *DetectVideoIPCObjectResponse {
	s.Body = v
	return s
}

func (s *DetectVideoIPCObjectResponse) Validate() error {
	return dara.Validate(s)
}
