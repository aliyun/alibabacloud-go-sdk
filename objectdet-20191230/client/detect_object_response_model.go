// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetectObjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DetectObjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DetectObjectResponse
	GetStatusCode() *int32
	SetBody(v *DetectObjectResponseBody) *DetectObjectResponse
	GetBody() *DetectObjectResponseBody
}

type DetectObjectResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DetectObjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DetectObjectResponse) String() string {
	return dara.Prettify(s)
}

func (s DetectObjectResponse) GoString() string {
	return s.String()
}

func (s *DetectObjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DetectObjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DetectObjectResponse) GetBody() *DetectObjectResponseBody {
	return s.Body
}

func (s *DetectObjectResponse) SetHeaders(v map[string]*string) *DetectObjectResponse {
	s.Headers = v
	return s
}

func (s *DetectObjectResponse) SetStatusCode(v int32) *DetectObjectResponse {
	s.StatusCode = &v
	return s
}

func (s *DetectObjectResponse) SetBody(v *DetectObjectResponseBody) *DetectObjectResponse {
	s.Body = v
	return s
}

func (s *DetectObjectResponse) Validate() error {
	return dara.Validate(s)
}
