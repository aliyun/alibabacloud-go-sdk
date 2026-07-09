// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceCustomizedDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCertId(v string) *GetInstanceCustomizedDomainResponseBody
	GetCertId() *string
	SetCode(v string) *GetInstanceCustomizedDomainResponseBody
	GetCode() *string
	SetCreateTime(v int64) *GetInstanceCustomizedDomainResponseBody
	GetCreateTime() *int64
	SetDomain(v string) *GetInstanceCustomizedDomainResponseBody
	GetDomain() *string
	SetDomainType(v string) *GetInstanceCustomizedDomainResponseBody
	GetDomainType() *string
	SetEndpointType(v string) *GetInstanceCustomizedDomainResponseBody
	GetEndpointType() *string
	SetInstanceId(v string) *GetInstanceCustomizedDomainResponseBody
	GetInstanceId() *string
	SetIsSuccess(v bool) *GetInstanceCustomizedDomainResponseBody
	GetIsSuccess() *bool
	SetModifiedTime(v int64) *GetInstanceCustomizedDomainResponseBody
	GetModifiedTime() *int64
	SetModuleName(v string) *GetInstanceCustomizedDomainResponseBody
	GetModuleName() *string
	SetRegionId(v string) *GetInstanceCustomizedDomainResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetInstanceCustomizedDomainResponseBody
	GetRequestId() *string
}

type GetInstanceCustomizedDomainResponseBody struct {
	// example:
	//
	// 24858802
	CertId *string `json:"CertId,omitempty" xml:"CertId,omitempty"`
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 1571926439000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// docker-images.qu-in.club
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// USER
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	// example:
	//
	// internet
	EndpointType *string `json:"EndpointType,omitempty" xml:"EndpointType,omitempty"`
	// example:
	//
	// cri-4ec5xvj4j0l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// example:
	//
	// 1638259914000
	ModifiedTime *int64 `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// example:
	//
	// Chart
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 4CE1F661-75DD-4EBD-A4AD-057B26834ABB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceCustomizedDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceCustomizedDomainResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceCustomizedDomainResponseBody) GetCertId() *string {
	return s.CertId
}

func (s *GetInstanceCustomizedDomainResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceCustomizedDomainResponseBody) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *GetInstanceCustomizedDomainResponseBody) GetDomain() *string {
	return s.Domain
}

func (s *GetInstanceCustomizedDomainResponseBody) GetDomainType() *string {
	return s.DomainType
}

func (s *GetInstanceCustomizedDomainResponseBody) GetEndpointType() *string {
	return s.EndpointType
}

func (s *GetInstanceCustomizedDomainResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceCustomizedDomainResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetInstanceCustomizedDomainResponseBody) GetModifiedTime() *int64 {
	return s.ModifiedTime
}

func (s *GetInstanceCustomizedDomainResponseBody) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetInstanceCustomizedDomainResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceCustomizedDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceCustomizedDomainResponseBody) SetCertId(v string) *GetInstanceCustomizedDomainResponseBody {
	s.CertId = &v
	return s
}

func (s *GetInstanceCustomizedDomainResponseBody) SetCode(v string) *GetInstanceCustomizedDomainResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceCustomizedDomainResponseBody) SetCreateTime(v int64) *GetInstanceCustomizedDomainResponseBody {
	s.CreateTime = &v
	return s
}

func (s *GetInstanceCustomizedDomainResponseBody) SetDomain(v string) *GetInstanceCustomizedDomainResponseBody {
	s.Domain = &v
	return s
}

func (s *GetInstanceCustomizedDomainResponseBody) SetDomainType(v string) *GetInstanceCustomizedDomainResponseBody {
	s.DomainType = &v
	return s
}

func (s *GetInstanceCustomizedDomainResponseBody) SetEndpointType(v string) *GetInstanceCustomizedDomainResponseBody {
	s.EndpointType = &v
	return s
}

func (s *GetInstanceCustomizedDomainResponseBody) SetInstanceId(v string) *GetInstanceCustomizedDomainResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceCustomizedDomainResponseBody) SetIsSuccess(v bool) *GetInstanceCustomizedDomainResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetInstanceCustomizedDomainResponseBody) SetModifiedTime(v int64) *GetInstanceCustomizedDomainResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetInstanceCustomizedDomainResponseBody) SetModuleName(v string) *GetInstanceCustomizedDomainResponseBody {
	s.ModuleName = &v
	return s
}

func (s *GetInstanceCustomizedDomainResponseBody) SetRegionId(v string) *GetInstanceCustomizedDomainResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetInstanceCustomizedDomainResponseBody) SetRequestId(v string) *GetInstanceCustomizedDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceCustomizedDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
