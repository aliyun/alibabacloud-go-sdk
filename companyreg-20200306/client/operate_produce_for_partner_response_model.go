// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateProduceForPartnerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateProduceForPartnerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateProduceForPartnerResponse
	GetStatusCode() *int32
	SetBody(v *OperateProduceForPartnerResponseBody) *OperateProduceForPartnerResponse
	GetBody() *OperateProduceForPartnerResponseBody
}

type OperateProduceForPartnerResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateProduceForPartnerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateProduceForPartnerResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateProduceForPartnerResponse) GoString() string {
	return s.String()
}

func (s *OperateProduceForPartnerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateProduceForPartnerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateProduceForPartnerResponse) GetBody() *OperateProduceForPartnerResponseBody {
	return s.Body
}

func (s *OperateProduceForPartnerResponse) SetHeaders(v map[string]*string) *OperateProduceForPartnerResponse {
	s.Headers = v
	return s
}

func (s *OperateProduceForPartnerResponse) SetStatusCode(v int32) *OperateProduceForPartnerResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateProduceForPartnerResponse) SetBody(v *OperateProduceForPartnerResponseBody) *OperateProduceForPartnerResponse {
	s.Body = v
	return s
}

func (s *OperateProduceForPartnerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
