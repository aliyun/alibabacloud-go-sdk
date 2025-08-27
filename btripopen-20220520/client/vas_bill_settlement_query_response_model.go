// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVasBillSettlementQueryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VasBillSettlementQueryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VasBillSettlementQueryResponse
	GetStatusCode() *int32
	SetBody(v *VasBillSettlementQueryResponseBody) *VasBillSettlementQueryResponse
	GetBody() *VasBillSettlementQueryResponseBody
}

type VasBillSettlementQueryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VasBillSettlementQueryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VasBillSettlementQueryResponse) String() string {
	return dara.Prettify(s)
}

func (s VasBillSettlementQueryResponse) GoString() string {
	return s.String()
}

func (s *VasBillSettlementQueryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VasBillSettlementQueryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VasBillSettlementQueryResponse) GetBody() *VasBillSettlementQueryResponseBody {
	return s.Body
}

func (s *VasBillSettlementQueryResponse) SetHeaders(v map[string]*string) *VasBillSettlementQueryResponse {
	s.Headers = v
	return s
}

func (s *VasBillSettlementQueryResponse) SetStatusCode(v int32) *VasBillSettlementQueryResponse {
	s.StatusCode = &v
	return s
}

func (s *VasBillSettlementQueryResponse) SetBody(v *VasBillSettlementQueryResponseBody) *VasBillSettlementQueryResponse {
	s.Body = v
	return s
}

func (s *VasBillSettlementQueryResponse) Validate() error {
	return dara.Validate(s)
}
