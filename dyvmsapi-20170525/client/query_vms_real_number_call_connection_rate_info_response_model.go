// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVmsRealNumberCallConnectionRateInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVmsRealNumberCallConnectionRateInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVmsRealNumberCallConnectionRateInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryVmsRealNumberCallConnectionRateInfoResponseBody) *QueryVmsRealNumberCallConnectionRateInfoResponse
	GetBody() *QueryVmsRealNumberCallConnectionRateInfoResponseBody
}

type QueryVmsRealNumberCallConnectionRateInfoResponse struct {
	Headers    map[string]*string                                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVmsRealNumberCallConnectionRateInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVmsRealNumberCallConnectionRateInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVmsRealNumberCallConnectionRateInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponse) GetBody() *QueryVmsRealNumberCallConnectionRateInfoResponseBody {
	return s.Body
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponse) SetHeaders(v map[string]*string) *QueryVmsRealNumberCallConnectionRateInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponse) SetStatusCode(v int32) *QueryVmsRealNumberCallConnectionRateInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponse) SetBody(v *QueryVmsRealNumberCallConnectionRateInfoResponseBody) *QueryVmsRealNumberCallConnectionRateInfoResponse {
	s.Body = v
	return s
}

func (s *QueryVmsRealNumberCallConnectionRateInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
