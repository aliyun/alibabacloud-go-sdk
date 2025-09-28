// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyStatisticPublicResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGateVerifyStatisticPublicResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGateVerifyStatisticPublicResponse
	GetStatusCode() *int32
	SetBody(v *QueryGateVerifyStatisticPublicResponseBody) *QueryGateVerifyStatisticPublicResponse
	GetBody() *QueryGateVerifyStatisticPublicResponseBody
}

type QueryGateVerifyStatisticPublicResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGateVerifyStatisticPublicResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGateVerifyStatisticPublicResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyStatisticPublicResponse) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticPublicResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGateVerifyStatisticPublicResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGateVerifyStatisticPublicResponse) GetBody() *QueryGateVerifyStatisticPublicResponseBody {
	return s.Body
}

func (s *QueryGateVerifyStatisticPublicResponse) SetHeaders(v map[string]*string) *QueryGateVerifyStatisticPublicResponse {
	s.Headers = v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponse) SetStatusCode(v int32) *QueryGateVerifyStatisticPublicResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponse) SetBody(v *QueryGateVerifyStatisticPublicResponseBody) *QueryGateVerifyStatisticPublicResponse {
	s.Body = v
	return s
}

func (s *QueryGateVerifyStatisticPublicResponse) Validate() error {
	return dara.Validate(s)
}
