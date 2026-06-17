// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSaasServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCu(v string) *GetSaasServiceResponseBody
	GetCu() *string
	SetPayType(v string) *GetSaasServiceResponseBody
	GetPayType() *string
	SetRegionId(v string) *GetSaasServiceResponseBody
	GetRegionId() *string
	SetRequestId(v string) *GetSaasServiceResponseBody
	GetRequestId() *string
	SetServiceId(v string) *GetSaasServiceResponseBody
	GetServiceId() *string
	SetServiceName(v string) *GetSaasServiceResponseBody
	GetServiceName() *string
	SetServiceType(v string) *GetSaasServiceResponseBody
	GetServiceType() *string
	SetStatus(v string) *GetSaasServiceResponseBody
	GetStatus() *string
	SetWorkspaceId(v string) *GetSaasServiceResponseBody
	GetWorkspaceId() *string
}

type GetSaasServiceResponseBody struct {
	// example:
	//
	// 1
	Cu *string `json:"Cu,omitempty" xml:"Cu,omitempty"`
	// example:
	//
	// POSTPAY
	PayType *string `json:"PayType,omitempty" xml:"PayType,omitempty"`
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
	// example:
	//
	// agdb-xxxxx
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
	// example:
	//
	// drama
	ServiceType *string `json:"ServiceType,omitempty" xml:"ServiceType,omitempty"`
	// example:
	//
	// creating
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// ws-*****
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s GetSaasServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSaasServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetSaasServiceResponseBody) GetCu() *string {
	return s.Cu
}

func (s *GetSaasServiceResponseBody) GetPayType() *string {
	return s.PayType
}

func (s *GetSaasServiceResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSaasServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSaasServiceResponseBody) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetSaasServiceResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *GetSaasServiceResponseBody) GetServiceType() *string {
	return s.ServiceType
}

func (s *GetSaasServiceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSaasServiceResponseBody) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *GetSaasServiceResponseBody) SetCu(v string) *GetSaasServiceResponseBody {
	s.Cu = &v
	return s
}

func (s *GetSaasServiceResponseBody) SetPayType(v string) *GetSaasServiceResponseBody {
	s.PayType = &v
	return s
}

func (s *GetSaasServiceResponseBody) SetRegionId(v string) *GetSaasServiceResponseBody {
	s.RegionId = &v
	return s
}

func (s *GetSaasServiceResponseBody) SetRequestId(v string) *GetSaasServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSaasServiceResponseBody) SetServiceId(v string) *GetSaasServiceResponseBody {
	s.ServiceId = &v
	return s
}

func (s *GetSaasServiceResponseBody) SetServiceName(v string) *GetSaasServiceResponseBody {
	s.ServiceName = &v
	return s
}

func (s *GetSaasServiceResponseBody) SetServiceType(v string) *GetSaasServiceResponseBody {
	s.ServiceType = &v
	return s
}

func (s *GetSaasServiceResponseBody) SetStatus(v string) *GetSaasServiceResponseBody {
	s.Status = &v
	return s
}

func (s *GetSaasServiceResponseBody) SetWorkspaceId(v string) *GetSaasServiceResponseBody {
	s.WorkspaceId = &v
	return s
}

func (s *GetSaasServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
