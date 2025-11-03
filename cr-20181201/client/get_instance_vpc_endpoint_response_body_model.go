// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceVpcEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceVpcEndpointResponseBody
	GetCode() *string
	SetDomains(v []*string) *GetInstanceVpcEndpointResponseBody
	GetDomains() []*string
	SetEnable(v bool) *GetInstanceVpcEndpointResponseBody
	GetEnable() *bool
	SetIsSuccess(v bool) *GetInstanceVpcEndpointResponseBody
	GetIsSuccess() *bool
	SetLinkedVpcs(v []*GetInstanceVpcEndpointResponseBodyLinkedVpcs) *GetInstanceVpcEndpointResponseBody
	GetLinkedVpcs() []*GetInstanceVpcEndpointResponseBodyLinkedVpcs
	SetModuleName(v string) *GetInstanceVpcEndpointResponseBody
	GetModuleName() *string
	SetRequestId(v string) *GetInstanceVpcEndpointResponseBody
	GetRequestId() *string
}

type GetInstanceVpcEndpointResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Domain names.
	Domains []*string `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
	// Indicates whether the VPC endpoint is enabled. Valid values:
	//
	// 	- `true`
	//
	// 	- `false`
	//
	// example:
	//
	// true
	Enable *bool `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- `true`: The request is successful.
	//
	// 	- `false`: The request fails.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// List of linked VPCs
	LinkedVpcs []*GetInstanceVpcEndpointResponseBodyLinkedVpcs `json:"LinkedVpcs,omitempty" xml:"LinkedVpcs,omitempty" type:"Repeated"`
	// The name of the modules that can be accessed. Valid values:
	//
	// 	- `Registry`: image repositories.
	//
	// 	- `Chart`: Helm charts.
	//
	// example:
	//
	// Registry
	ModuleName *string `json:"ModuleName,omitempty" xml:"ModuleName,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// BAE9349D-A587-4F9A-B574-9DA0EF2638D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceVpcEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceVpcEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceVpcEndpointResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceVpcEndpointResponseBody) GetDomains() []*string {
	return s.Domains
}

func (s *GetInstanceVpcEndpointResponseBody) GetEnable() *bool {
	return s.Enable
}

func (s *GetInstanceVpcEndpointResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *GetInstanceVpcEndpointResponseBody) GetLinkedVpcs() []*GetInstanceVpcEndpointResponseBodyLinkedVpcs {
	return s.LinkedVpcs
}

func (s *GetInstanceVpcEndpointResponseBody) GetModuleName() *string {
	return s.ModuleName
}

func (s *GetInstanceVpcEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceVpcEndpointResponseBody) SetCode(v string) *GetInstanceVpcEndpointResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBody) SetDomains(v []*string) *GetInstanceVpcEndpointResponseBody {
	s.Domains = v
	return s
}

func (s *GetInstanceVpcEndpointResponseBody) SetEnable(v bool) *GetInstanceVpcEndpointResponseBody {
	s.Enable = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBody) SetIsSuccess(v bool) *GetInstanceVpcEndpointResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBody) SetLinkedVpcs(v []*GetInstanceVpcEndpointResponseBodyLinkedVpcs) *GetInstanceVpcEndpointResponseBody {
	s.LinkedVpcs = v
	return s
}

func (s *GetInstanceVpcEndpointResponseBody) SetModuleName(v string) *GetInstanceVpcEndpointResponseBody {
	s.ModuleName = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBody) SetRequestId(v string) *GetInstanceVpcEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBody) Validate() error {
	if s.LinkedVpcs != nil {
		for _, item := range s.LinkedVpcs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceVpcEndpointResponseBodyLinkedVpcs struct {
	// Indicates whether the VPC is the default VPC over which the Container Registry instance is accessed.
	//
	// example:
	//
	// false
	DefaultAccess *bool `json:"DefaultAccess,omitempty" xml:"DefaultAccess,omitempty"`
	// IP address.
	//
	// example:
	//
	// 192.168.10.11
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The error message detected in the linked VPC access control.
	//
	// example:
	//
	// PRIVATE_ZONE_CONFLICT_AT_{private_zone_id}
	Issue *string `json:"Issue,omitempty" xml:"Issue,omitempty"`
	// The status of the VPC. Valid values:
	//
	// 	- `CREATING`
	//
	// 	- `RUNNING`
	//
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-uf6aamu2nomfr1thd****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-uf62m5vmxl2e72dk7****
	VswitchId *string `json:"VswitchId,omitempty" xml:"VswitchId,omitempty"`
}

func (s GetInstanceVpcEndpointResponseBodyLinkedVpcs) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceVpcEndpointResponseBodyLinkedVpcs) GoString() string {
	return s.String()
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) GetDefaultAccess() *bool {
	return s.DefaultAccess
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) GetIp() *string {
	return s.Ip
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) GetIssue() *string {
	return s.Issue
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) GetStatus() *string {
	return s.Status
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) GetVpcId() *string {
	return s.VpcId
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) GetVswitchId() *string {
	return s.VswitchId
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) SetDefaultAccess(v bool) *GetInstanceVpcEndpointResponseBodyLinkedVpcs {
	s.DefaultAccess = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) SetIp(v string) *GetInstanceVpcEndpointResponseBodyLinkedVpcs {
	s.Ip = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) SetIssue(v string) *GetInstanceVpcEndpointResponseBodyLinkedVpcs {
	s.Issue = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) SetStatus(v string) *GetInstanceVpcEndpointResponseBodyLinkedVpcs {
	s.Status = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) SetVpcId(v string) *GetInstanceVpcEndpointResponseBodyLinkedVpcs {
	s.VpcId = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) SetVswitchId(v string) *GetInstanceVpcEndpointResponseBodyLinkedVpcs {
	s.VswitchId = &v
	return s
}

func (s *GetInstanceVpcEndpointResponseBodyLinkedVpcs) Validate() error {
	return dara.Validate(s)
}
