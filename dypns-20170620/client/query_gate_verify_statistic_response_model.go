// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryGateVerifyStatisticResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryGateVerifyStatisticResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryGateVerifyStatisticResponse
	GetStatusCode() *int32
	SetBody(v *QueryGateVerifyStatisticResponseBody) *QueryGateVerifyStatisticResponse
	GetBody() *QueryGateVerifyStatisticResponseBody
}

type QueryGateVerifyStatisticResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryGateVerifyStatisticResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryGateVerifyStatisticResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryGateVerifyStatisticResponse) GoString() string {
	return s.String()
}

func (s *QueryGateVerifyStatisticResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryGateVerifyStatisticResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryGateVerifyStatisticResponse) GetBody() *QueryGateVerifyStatisticResponseBody {
	return s.Body
}

func (s *QueryGateVerifyStatisticResponse) SetHeaders(v map[string]*string) *QueryGateVerifyStatisticResponse {
	s.Headers = v
	return s
}

func (s *QueryGateVerifyStatisticResponse) SetStatusCode(v int32) *QueryGateVerifyStatisticResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryGateVerifyStatisticResponse) SetBody(v *QueryGateVerifyStatisticResponseBody) *QueryGateVerifyStatisticResponse {
	s.Body = v
	return s
}

func (s *QueryGateVerifyStatisticResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
