// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListQuotasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetQuotas(v []*Quota) *ListQuotasResponseBody
	GetQuotas() []*Quota
	SetRequestId(v string) *ListQuotasResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListQuotasResponseBody
	GetTotalCount() *int32
}

type ListQuotasResponseBody struct {
	Quotas []*Quota `json:"Quotas,omitempty" xml:"Quotas,omitempty" type:"Repeated"`
	// example:
	//
	// F082BD0D-21E1-5F9B-81A0-AB07485B03CD
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListQuotasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListQuotasResponseBody) GoString() string {
	return s.String()
}

func (s *ListQuotasResponseBody) GetQuotas() []*Quota {
	return s.Quotas
}

func (s *ListQuotasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListQuotasResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListQuotasResponseBody) SetQuotas(v []*Quota) *ListQuotasResponseBody {
	s.Quotas = v
	return s
}

func (s *ListQuotasResponseBody) SetRequestId(v string) *ListQuotasResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListQuotasResponseBody) SetTotalCount(v int32) *ListQuotasResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListQuotasResponseBody) Validate() error {
	return dara.Validate(s)
}
