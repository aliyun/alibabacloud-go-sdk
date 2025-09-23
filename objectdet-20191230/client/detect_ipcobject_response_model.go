// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectIPCObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectIPCObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectIPCObjectResponse
	GetStatusCode() *int32
	SetBody(v *DetectIPCObjectResponseBody) *DetectIPCObjectResponse
	GetBody() *DetectIPCObjectResponseBody
}

type DetectIPCObjectResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectIPCObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectIPCObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectIPCObjectResponse) GoString() string {
	return s.String()
}

func (s *DetectIPCObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectIPCObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectIPCObjectResponse) GetBody() *DetectIPCObjectResponseBody {
	return s.Body
}

func (s *DetectIPCObjectResponse) SetHeaders(v map[string]*string) *DetectIPCObjectResponse {
	s.Headers = v
	return s
}

func (s *DetectIPCObjectResponse) SetStatusCode(v int32) *DetectIPCObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectIPCObjectResponse) SetBody(v *DetectIPCObjectResponseBody) *DetectIPCObjectResponse {
	s.Body = v
	return s
}

func (s *DetectIPCObjectResponse) Validate() error {
	return dara.Validate(s)
}
