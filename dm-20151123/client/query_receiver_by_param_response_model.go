// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryReceiverByParamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryReceiverByParamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryReceiverByParamResponse
	GetStatusCode() *int32
	SetBody(v *QueryReceiverByParamResponseBody) *QueryReceiverByParamResponse
	GetBody() *QueryReceiverByParamResponseBody
}

type QueryReceiverByParamResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryReceiverByParamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryReceiverByParamResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryReceiverByParamResponse) GoString() string {
	return s.String()
}

func (s *QueryReceiverByParamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryReceiverByParamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryReceiverByParamResponse) GetBody() *QueryReceiverByParamResponseBody {
	return s.Body
}

func (s *QueryReceiverByParamResponse) SetHeaders(v map[string]*string) *QueryReceiverByParamResponse {
	s.Headers = v
	return s
}

func (s *QueryReceiverByParamResponse) SetStatusCode(v int32) *QueryReceiverByParamResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryReceiverByParamResponse) SetBody(v *QueryReceiverByParamResponseBody) *QueryReceiverByParamResponse {
	s.Body = v
	return s
}

func (s *QueryReceiverByParamResponse) Validate() error {
	return dara.Validate(s)
}
