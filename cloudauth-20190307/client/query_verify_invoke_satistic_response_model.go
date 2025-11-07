// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVerifyInvokeSatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVerifyInvokeSatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVerifyInvokeSatisticResponse
	GetStatusCode() *int32
	SetBody(v *QueryVerifyInvokeSatisticResponseBody) *QueryVerifyInvokeSatisticResponse
	GetBody() *QueryVerifyInvokeSatisticResponseBody
}

type QueryVerifyInvokeSatisticResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVerifyInvokeSatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVerifyInvokeSatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVerifyInvokeSatisticResponse) GoString() string {
	return s.String()
}

func (s *QueryVerifyInvokeSatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVerifyInvokeSatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVerifyInvokeSatisticResponse) GetBody() *QueryVerifyInvokeSatisticResponseBody {
	return s.Body
}

func (s *QueryVerifyInvokeSatisticResponse) SetHeaders(v map[string]*string) *QueryVerifyInvokeSatisticResponse {
	s.Headers = v
	return s
}

func (s *QueryVerifyInvokeSatisticResponse) SetStatusCode(v int32) *QueryVerifyInvokeSatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVerifyInvokeSatisticResponse) SetBody(v *QueryVerifyInvokeSatisticResponseBody) *QueryVerifyInvokeSatisticResponse {
	s.Body = v
	return s
}

func (s *QueryVerifyInvokeSatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
