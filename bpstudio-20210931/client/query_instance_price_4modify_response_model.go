// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryInstancePrice4ModifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryInstancePrice4ModifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryInstancePrice4ModifyResponse
	GetStatusCode() *int32
	SetBody(v *QueryInstancePrice4ModifyResponseBody) *QueryInstancePrice4ModifyResponse
	GetBody() *QueryInstancePrice4ModifyResponseBody
}

type QueryInstancePrice4ModifyResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryInstancePrice4ModifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryInstancePrice4ModifyResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryInstancePrice4ModifyResponse) GoString() string {
	return s.String()
}

func (s *QueryInstancePrice4ModifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryInstancePrice4ModifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryInstancePrice4ModifyResponse) GetBody() *QueryInstancePrice4ModifyResponseBody {
	return s.Body
}

func (s *QueryInstancePrice4ModifyResponse) SetHeaders(v map[string]*string) *QueryInstancePrice4ModifyResponse {
	s.Headers = v
	return s
}

func (s *QueryInstancePrice4ModifyResponse) SetStatusCode(v int32) *QueryInstancePrice4ModifyResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryInstancePrice4ModifyResponse) SetBody(v *QueryInstancePrice4ModifyResponseBody) *QueryInstancePrice4ModifyResponse {
	s.Body = v
	return s
}

func (s *QueryInstancePrice4ModifyResponse) Validate() error {
	return dara.Validate(s)
}
