// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDomainSpecialBizCancelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DomainSpecialBizCancelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DomainSpecialBizCancelResponse
	GetStatusCode() *int32
	SetBody(v *DomainSpecialBizCancelResponseBody) *DomainSpecialBizCancelResponse
	GetBody() *DomainSpecialBizCancelResponseBody
}

type DomainSpecialBizCancelResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DomainSpecialBizCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DomainSpecialBizCancelResponse) String() string {
	return dara.Prettify(s)
}

func (s DomainSpecialBizCancelResponse) GoString() string {
	return s.String()
}

func (s *DomainSpecialBizCancelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DomainSpecialBizCancelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DomainSpecialBizCancelResponse) GetBody() *DomainSpecialBizCancelResponseBody {
	return s.Body
}

func (s *DomainSpecialBizCancelResponse) SetHeaders(v map[string]*string) *DomainSpecialBizCancelResponse {
	s.Headers = v
	return s
}

func (s *DomainSpecialBizCancelResponse) SetStatusCode(v int32) *DomainSpecialBizCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *DomainSpecialBizCancelResponse) SetBody(v *DomainSpecialBizCancelResponseBody) *DomainSpecialBizCancelResponse {
	s.Body = v
	return s
}

func (s *DomainSpecialBizCancelResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
