// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaId(v string) *DeleteQuotaResponseBody
	GetQuotaId() *string
	SetRequestId(v string) *DeleteQuotaResponseBody
	GetRequestId() *string
}

type DeleteQuotaResponseBody struct {
	// Quota Id
	//
	// example:
	//
	// quotamtl37ge7gkvdz
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQuotaResponseBody) GetQuotaId() *string {
	return s.QuotaId
}

func (s *DeleteQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQuotaResponseBody) SetQuotaId(v string) *DeleteQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *DeleteQuotaResponseBody) SetRequestId(v string) *DeleteQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
