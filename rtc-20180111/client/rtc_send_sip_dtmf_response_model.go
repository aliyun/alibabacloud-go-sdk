// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRtcSendSipDtmfResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RtcSendSipDtmfResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RtcSendSipDtmfResponse
	GetStatusCode() *int32
	SetBody(v *RtcSendSipDtmfResponseBody) *RtcSendSipDtmfResponse
	GetBody() *RtcSendSipDtmfResponseBody
}

type RtcSendSipDtmfResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RtcSendSipDtmfResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RtcSendSipDtmfResponse) String() string {
	return dara.Prettify(s)
}

func (s RtcSendSipDtmfResponse) GoString() string {
	return s.String()
}

func (s *RtcSendSipDtmfResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RtcSendSipDtmfResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RtcSendSipDtmfResponse) GetBody() *RtcSendSipDtmfResponseBody {
	return s.Body
}

func (s *RtcSendSipDtmfResponse) SetHeaders(v map[string]*string) *RtcSendSipDtmfResponse {
	s.Headers = v
	return s
}

func (s *RtcSendSipDtmfResponse) SetStatusCode(v int32) *RtcSendSipDtmfResponse {
	s.StatusCode = &v
	return s
}

func (s *RtcSendSipDtmfResponse) SetBody(v *RtcSendSipDtmfResponseBody) *RtcSendSipDtmfResponse {
	s.Body = v
	return s
}

func (s *RtcSendSipDtmfResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
