// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySignInRecordPopResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySignInRecordPopResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySignInRecordPopResponse
	GetStatusCode() *int32
	SetBody(v *QuerySignInRecordPopResponseBody) *QuerySignInRecordPopResponse
	GetBody() *QuerySignInRecordPopResponseBody
}

type QuerySignInRecordPopResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySignInRecordPopResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySignInRecordPopResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySignInRecordPopResponse) GoString() string {
	return s.String()
}

func (s *QuerySignInRecordPopResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySignInRecordPopResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySignInRecordPopResponse) GetBody() *QuerySignInRecordPopResponseBody {
	return s.Body
}

func (s *QuerySignInRecordPopResponse) SetHeaders(v map[string]*string) *QuerySignInRecordPopResponse {
	s.Headers = v
	return s
}

func (s *QuerySignInRecordPopResponse) SetStatusCode(v int32) *QuerySignInRecordPopResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySignInRecordPopResponse) SetBody(v *QuerySignInRecordPopResponseBody) *QuerySignInRecordPopResponse {
	s.Body = v
	return s
}

func (s *QuerySignInRecordPopResponse) Validate() error {
	return dara.Validate(s)
}
