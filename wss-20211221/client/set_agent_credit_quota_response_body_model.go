// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAgentCreditQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetAgentCreditQuotaResponseBody
	GetRequestId() *string
}

type SetAgentCreditQuotaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9B353D6D-53C9-1C24-95D5-2D24596DBCF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetAgentCreditQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetAgentCreditQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *SetAgentCreditQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetAgentCreditQuotaResponseBody) SetRequestId(v string) *SetAgentCreditQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAgentCreditQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
