// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateCallCenterForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateCallCenterForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateCallCenterForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *OperateCallCenterForPartnerResponseBody) *OperateCallCenterForPartnerResponse
	GetBody() *OperateCallCenterForPartnerResponseBody
}

type OperateCallCenterForPartnerResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateCallCenterForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateCallCenterForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateCallCenterForPartnerResponse) GoString() string {
	return s.String()
}

func (s *OperateCallCenterForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateCallCenterForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateCallCenterForPartnerResponse) GetBody() *OperateCallCenterForPartnerResponseBody {
	return s.Body
}

func (s *OperateCallCenterForPartnerResponse) SetHeaders(v map[string]*string) *OperateCallCenterForPartnerResponse {
	s.Headers = v
	return s
}

func (s *OperateCallCenterForPartnerResponse) SetStatusCode(v int32) *OperateCallCenterForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateCallCenterForPartnerResponse) SetBody(v *OperateCallCenterForPartnerResponseBody) *OperateCallCenterForPartnerResponse {
	s.Body = v
	return s
}

func (s *OperateCallCenterForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
