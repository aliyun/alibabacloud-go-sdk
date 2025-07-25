// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iScaleQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuotaId(v string) *ScaleQuotaResponseBody
	GetQuotaId() *string
	SetRequestId(v string) *ScaleQuotaResponseBody
	GetRequestId() *string
}

type ScaleQuotaResponseBody struct {
	// Quota Id
	//
	// example:
	//
	// quotamtl37ge7gkvdz
	QuotaId *string `json:"QuotaId,omitempty" xml:"QuotaId,omitempty"`
	// example:
	//
	// F2D0392B-D749-5C48-A98A-3FAE5C9444A6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ScaleQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ScaleQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *ScaleQuotaResponseBody) GetQuotaId() *string {
	return s.QuotaId
}

func (s *ScaleQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ScaleQuotaResponseBody) SetQuotaId(v string) *ScaleQuotaResponseBody {
	s.QuotaId = &v
	return s
}

func (s *ScaleQuotaResponseBody) SetRequestId(v string) *ScaleQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *ScaleQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
