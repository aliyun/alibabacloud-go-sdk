// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrometheusViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthFreeReadPolicy(v string) *UpdatePrometheusViewRequest
	GetAuthFreeReadPolicy() *string
	SetEnableAuthFreeRead(v bool) *UpdatePrometheusViewRequest
	GetEnableAuthFreeRead() *bool
	SetEnableAuthToken(v bool) *UpdatePrometheusViewRequest
	GetEnableAuthToken() *bool
	SetPrometheusInstances(v []*UpdatePrometheusViewRequestPrometheusInstances) *UpdatePrometheusViewRequest
	GetPrometheusInstances() []*UpdatePrometheusViewRequestPrometheusInstances
	SetPrometheusViewName(v string) *UpdatePrometheusViewRequest
	GetPrometheusViewName() *string
	SetStatus(v string) *UpdatePrometheusViewRequest
	GetStatus() *string
	SetWorkspace(v string) *UpdatePrometheusViewRequest
	GetWorkspace() *string
}

type UpdatePrometheusViewRequest struct {
	// Password-free read policy (supports IP segments and VpcId).
	//
	// example:
	//
	// {
	//
	//   "SourceIp": [
	//
	//     "192.168.1.0/24",
	//
	//     "172.168.2.22"
	//
	//   ],
	//
	//   "SourceVpc": [
	//
	//     "vpc-xx1",
	//
	//     "vpc-xx2"
	//
	//   ]
	//
	// }
	AuthFreeReadPolicy *string `json:"authFreeReadPolicy,omitempty" xml:"authFreeReadPolicy,omitempty"`
	// Whether to support password-free read.
	//
	// example:
	//
	// true
	EnableAuthFreeRead *bool `json:"enableAuthFreeRead,omitempty" xml:"enableAuthFreeRead,omitempty"`
	// Whether to support authToken.
	//
	// example:
	//
	// false
	EnableAuthToken *bool `json:"enableAuthToken,omitempty" xml:"enableAuthToken,omitempty"`
	// List of Prometheus instances.
	PrometheusInstances []*UpdatePrometheusViewRequestPrometheusInstances `json:"prometheusInstances,omitempty" xml:"prometheusInstances,omitempty" type:"Repeated"`
	// Prometheus view name.
	//
	// example:
	//
	// test-prom-view-name
	PrometheusViewName *string `json:"prometheusViewName,omitempty" xml:"prometheusViewName,omitempty"`
	// Running status.
	//
	// example:
	//
	// Running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// Belonging workspace.
	//
	// example:
	//
	// default-cms-108490012345-cn-heyuan
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s UpdatePrometheusViewRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusViewRequest) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusViewRequest) GetAuthFreeReadPolicy() *string {
	return s.AuthFreeReadPolicy
}

func (s *UpdatePrometheusViewRequest) GetEnableAuthFreeRead() *bool {
	return s.EnableAuthFreeRead
}

func (s *UpdatePrometheusViewRequest) GetEnableAuthToken() *bool {
	return s.EnableAuthToken
}

func (s *UpdatePrometheusViewRequest) GetPrometheusInstances() []*UpdatePrometheusViewRequestPrometheusInstances {
	return s.PrometheusInstances
}

func (s *UpdatePrometheusViewRequest) GetPrometheusViewName() *string {
	return s.PrometheusViewName
}

func (s *UpdatePrometheusViewRequest) GetStatus() *string {
	return s.Status
}

func (s *UpdatePrometheusViewRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *UpdatePrometheusViewRequest) SetAuthFreeReadPolicy(v string) *UpdatePrometheusViewRequest {
	s.AuthFreeReadPolicy = &v
	return s
}

func (s *UpdatePrometheusViewRequest) SetEnableAuthFreeRead(v bool) *UpdatePrometheusViewRequest {
	s.EnableAuthFreeRead = &v
	return s
}

func (s *UpdatePrometheusViewRequest) SetEnableAuthToken(v bool) *UpdatePrometheusViewRequest {
	s.EnableAuthToken = &v
	return s
}

func (s *UpdatePrometheusViewRequest) SetPrometheusInstances(v []*UpdatePrometheusViewRequestPrometheusInstances) *UpdatePrometheusViewRequest {
	s.PrometheusInstances = v
	return s
}

func (s *UpdatePrometheusViewRequest) SetPrometheusViewName(v string) *UpdatePrometheusViewRequest {
	s.PrometheusViewName = &v
	return s
}

func (s *UpdatePrometheusViewRequest) SetStatus(v string) *UpdatePrometheusViewRequest {
	s.Status = &v
	return s
}

func (s *UpdatePrometheusViewRequest) SetWorkspace(v string) *UpdatePrometheusViewRequest {
	s.Workspace = &v
	return s
}

func (s *UpdatePrometheusViewRequest) Validate() error {
	if s.PrometheusInstances != nil {
		for _, item := range s.PrometheusInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdatePrometheusViewRequestPrometheusInstances struct {
	// Instance ID.
	//
	// example:
	//
	// c7ba84651c71e442c8d0653085d862164
	PrometheusInstanceId *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
	// Region.
	//
	// example:
	//
	// cn-north-2-gov-1
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// User ID.
	//
	// example:
	//
	// 16727123456789
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s UpdatePrometheusViewRequestPrometheusInstances) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrometheusViewRequestPrometheusInstances) GoString() string {
	return s.String()
}

func (s *UpdatePrometheusViewRequestPrometheusInstances) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *UpdatePrometheusViewRequestPrometheusInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdatePrometheusViewRequestPrometheusInstances) GetUserId() *string {
	return s.UserId
}

func (s *UpdatePrometheusViewRequestPrometheusInstances) SetPrometheusInstanceId(v string) *UpdatePrometheusViewRequestPrometheusInstances {
	s.PrometheusInstanceId = &v
	return s
}

func (s *UpdatePrometheusViewRequestPrometheusInstances) SetRegionId(v string) *UpdatePrometheusViewRequestPrometheusInstances {
	s.RegionId = &v
	return s
}

func (s *UpdatePrometheusViewRequestPrometheusInstances) SetUserId(v string) *UpdatePrometheusViewRequestPrometheusInstances {
	s.UserId = &v
	return s
}

func (s *UpdatePrometheusViewRequestPrometheusInstances) Validate() error {
	return dara.Validate(s)
}
