// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartAndroidInstanceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartAndroidInstanceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartAndroidInstanceResponse
	GetStatusCode() *int32
	SetBody(v *StartAndroidInstanceResponseBody) *StartAndroidInstanceResponse
	GetBody() *StartAndroidInstanceResponseBody
}

type StartAndroidInstanceResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartAndroidInstanceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartAndroidInstanceResponse) String() string {
	return dara.Prettify(s)
}

func (s StartAndroidInstanceResponse) GoString() string {
	return s.String()
}

func (s *StartAndroidInstanceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartAndroidInstanceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartAndroidInstanceResponse) GetBody() *StartAndroidInstanceResponseBody {
	return s.Body
}

func (s *StartAndroidInstanceResponse) SetHeaders(v map[string]*string) *StartAndroidInstanceResponse {
	s.Headers = v
	return s
}

func (s *StartAndroidInstanceResponse) SetStatusCode(v int32) *StartAndroidInstanceResponse {
	s.StatusCode = &v
	return s
}

func (s *StartAndroidInstanceResponse) SetBody(v *StartAndroidInstanceResponseBody) *StartAndroidInstanceResponse {
	s.Body = v
	return s
}

func (s *StartAndroidInstanceResponse) Validate() error {
	return dara.Validate(s)
}
