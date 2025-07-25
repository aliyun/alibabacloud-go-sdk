// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaId(v string) *UpdateQuotaResponseBody
	GetQuotaId() *string
	SetRequestId(v string) *UpdateQuotaResponseBody
	GetRequestId() *string
}

type UpdateQuotaResponseBody struct {
	// Quota Id
	//
	// example:
	//
	// quota-20210126170216-mtl37ge7gkvdz
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// 96496E6E-00B4-5F55-80F6-1844FA9E92DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateQuotaResponseBody) GetQuotaId() *string {
	return s.QuotaId
}

func (s *UpdateQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateQuotaResponseBody) SetQuotaId(v string) *UpdateQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *UpdateQuotaResponseBody) SetRequestId(v string) *UpdateQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
