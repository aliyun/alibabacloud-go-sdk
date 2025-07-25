// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaId(v string) *CreateQuotaResponseBody
	GetQuotaId() *string
	SetRequestId(v string) *CreateQuotaResponseBody
	GetRequestId() *string
}

type CreateQuotaResponseBody struct {
	// Quota Id
	//
	// example:
	//
	// quotad2kd8ljpsno
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// CBF05F13-B24C-5129-9048-4FA684DCD579
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateQuotaResponseBody) GetQuotaId() *string {
	return s.QuotaId
}

func (s *CreateQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateQuotaResponseBody) SetQuotaId(v string) *CreateQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *CreateQuotaResponseBody) SetRequestId(v string) *CreateQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
