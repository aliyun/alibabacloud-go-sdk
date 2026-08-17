// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaInstanceCrlAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCaInstanceStatus(v string) *GetCaInstanceCrlAddressResponseBody
	GetCaInstanceStatus() *string
	SetCaType(v string) *GetCaInstanceCrlAddressResponseBody
	GetCaType() *string
	SetCrlUrl(v string) *GetCaInstanceCrlAddressResponseBody
	GetCrlUrl() *string
	SetHashCode(v string) *GetCaInstanceCrlAddressResponseBody
	GetHashCode() *string
	SetNextUpdateTime(v string) *GetCaInstanceCrlAddressResponseBody
	GetNextUpdateTime() *string
	SetRequestId(v string) *GetCaInstanceCrlAddressResponseBody
	GetRequestId() *string
}

type GetCaInstanceCrlAddressResponseBody struct {
	// The status of the CA instance.
	//
	// example:
	//
	// normal
	CaInstanceStatus *string `json:"CaInstanceStatus,omitempty" xml:"CaInstanceStatus,omitempty"`
	CaType           *string `json:"CaType,omitempty" xml:"CaType,omitempty"`
	// The CRL URL.
	//
	// example:
	//
	// https://crl-cn-publish.oss-cn-hangzhou.aliyuncs.com/pca/crl/35118048/1f0be094-14bd-6caa-bd7f-db45730d510a.crl
	CrlUrl *string `json:"CrlUrl,omitempty" xml:"CrlUrl,omitempty"`
	// The hash code used to identify whether the CRL contains new revoked certificates.
	//
	// example:
	//
	// 5481d1b1228fXXX40ee70dc8cd
	HashCode *string `json:"HashCode,omitempty" xml:"HashCode,omitempty"`
	// The next update time of the CRL.
	//
	// example:
	//
	// 1778688000000
	NextUpdateTime *string `json:"NextUpdateTime,omitempty" xml:"NextUpdateTime,omitempty"`
	// Id of the request
	//
	// example:
	//
	// C53C2341-F321-55A5-895C-D0746E6DA58E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetCaInstanceCrlAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetCaInstanceCrlAddressResponseBody) GoString() string {
	return s.String()
}

func (s *GetCaInstanceCrlAddressResponseBody) GetCaInstanceStatus() *string {
	return s.CaInstanceStatus
}

func (s *GetCaInstanceCrlAddressResponseBody) GetCaType() *string {
	return s.CaType
}

func (s *GetCaInstanceCrlAddressResponseBody) GetCrlUrl() *string {
	return s.CrlUrl
}

func (s *GetCaInstanceCrlAddressResponseBody) GetHashCode() *string {
	return s.HashCode
}

func (s *GetCaInstanceCrlAddressResponseBody) GetNextUpdateTime() *string {
	return s.NextUpdateTime
}

func (s *GetCaInstanceCrlAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetCaInstanceCrlAddressResponseBody) SetCaInstanceStatus(v string) *GetCaInstanceCrlAddressResponseBody {
	s.CaInstanceStatus = &v
	return s
}

func (s *GetCaInstanceCrlAddressResponseBody) SetCaType(v string) *GetCaInstanceCrlAddressResponseBody {
	s.CaType = &v
	return s
}

func (s *GetCaInstanceCrlAddressResponseBody) SetCrlUrl(v string) *GetCaInstanceCrlAddressResponseBody {
	s.CrlUrl = &v
	return s
}

func (s *GetCaInstanceCrlAddressResponseBody) SetHashCode(v string) *GetCaInstanceCrlAddressResponseBody {
	s.HashCode = &v
	return s
}

func (s *GetCaInstanceCrlAddressResponseBody) SetNextUpdateTime(v string) *GetCaInstanceCrlAddressResponseBody {
	s.NextUpdateTime = &v
	return s
}

func (s *GetCaInstanceCrlAddressResponseBody) SetRequestId(v string) *GetCaInstanceCrlAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetCaInstanceCrlAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
