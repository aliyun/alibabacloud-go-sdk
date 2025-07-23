// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstanceSpec4ModifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInstanceSpec4ModifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInstanceSpec4ModifyResponse
	GetStatusCode() *int32
	SetBody(v *QueryInstanceSpec4ModifyResponseBody) *QueryInstanceSpec4ModifyResponse
	GetBody() *QueryInstanceSpec4ModifyResponseBody
}

type QueryInstanceSpec4ModifyResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstanceSpec4ModifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstanceSpec4ModifyResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInstanceSpec4ModifyResponse) GoString() string {
	return s.String()
}

func (s *QueryInstanceSpec4ModifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInstanceSpec4ModifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInstanceSpec4ModifyResponse) GetBody() *QueryInstanceSpec4ModifyResponseBody {
	return s.Body
}

func (s *QueryInstanceSpec4ModifyResponse) SetHeaders(v map[string]*string) *QueryInstanceSpec4ModifyResponse {
	s.Headers = v
	return s
}

func (s *QueryInstanceSpec4ModifyResponse) SetStatusCode(v int32) *QueryInstanceSpec4ModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstanceSpec4ModifyResponse) SetBody(v *QueryInstanceSpec4ModifyResponseBody) *QueryInstanceSpec4ModifyResponse {
	s.Body = v
	return s
}

func (s *QueryInstanceSpec4ModifyResponse) Validate() error {
	return dara.Validate(s)
}
