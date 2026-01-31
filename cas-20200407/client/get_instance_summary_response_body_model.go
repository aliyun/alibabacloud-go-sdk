// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceSummaryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoReissueCount(v int32) *GetInstanceSummaryResponseBody
	GetAutoReissueCount() *int32
	SetCertificateCount(v int32) *GetInstanceSummaryResponseBody
	GetCertificateCount() *int32
	SetInactiveCount(v int32) *GetInstanceSummaryResponseBody
	GetInactiveCount() *int32
	SetRequestId(v string) *GetInstanceSummaryResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *GetInstanceSummaryResponseBody
	GetTotalCount() *int32
	SetWillExpireCount(v int32) *GetInstanceSummaryResponseBody
	GetWillExpireCount() *int32
}

type GetInstanceSummaryResponseBody struct {
	// example:
	//
	// 1
	AutoReissueCount *int32 `json:"AutoReissueCount,omitempty" xml:"AutoReissueCount,omitempty"`
	// example:
	//
	// 1
	CertificateCount *int32 `json:"CertificateCount,omitempty" xml:"CertificateCount,omitempty"`
	// example:
	//
	// 1
	InactiveCount *int32 `json:"InactiveCount,omitempty" xml:"InactiveCount,omitempty"`
	// example:
	//
	// 09470F19-CEE8-5C63-BF2C-02B5E3F07A17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// example:
	//
	// 1
	WillExpireCount *int32 `json:"WillExpireCount,omitempty" xml:"WillExpireCount,omitempty"`
}

func (s GetInstanceSummaryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceSummaryResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceSummaryResponseBody) GetAutoReissueCount() *int32 {
	return s.AutoReissueCount
}

func (s *GetInstanceSummaryResponseBody) GetCertificateCount() *int32 {
	return s.CertificateCount
}

func (s *GetInstanceSummaryResponseBody) GetInactiveCount() *int32 {
	return s.InactiveCount
}

func (s *GetInstanceSummaryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceSummaryResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetInstanceSummaryResponseBody) GetWillExpireCount() *int32 {
	return s.WillExpireCount
}

func (s *GetInstanceSummaryResponseBody) SetAutoReissueCount(v int32) *GetInstanceSummaryResponseBody {
	s.AutoReissueCount = &v
	return s
}

func (s *GetInstanceSummaryResponseBody) SetCertificateCount(v int32) *GetInstanceSummaryResponseBody {
	s.CertificateCount = &v
	return s
}

func (s *GetInstanceSummaryResponseBody) SetInactiveCount(v int32) *GetInstanceSummaryResponseBody {
	s.InactiveCount = &v
	return s
}

func (s *GetInstanceSummaryResponseBody) SetRequestId(v string) *GetInstanceSummaryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceSummaryResponseBody) SetTotalCount(v int32) *GetInstanceSummaryResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetInstanceSummaryResponseBody) SetWillExpireCount(v int32) *GetInstanceSummaryResponseBody {
	s.WillExpireCount = &v
	return s
}

func (s *GetInstanceSummaryResponseBody) Validate() error {
	return dara.Validate(s)
}
