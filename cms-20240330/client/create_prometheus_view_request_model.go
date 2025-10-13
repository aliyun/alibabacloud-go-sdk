// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePrometheusViewRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthFreeReadPolicy(v string) *CreatePrometheusViewRequest
	GetAuthFreeReadPolicy() *string
	SetEnableAuthFreeRead(v bool) *CreatePrometheusViewRequest
	GetEnableAuthFreeRead() *bool
	SetEnableAuthToken(v bool) *CreatePrometheusViewRequest
	GetEnableAuthToken() *bool
	SetPrometheusInstances(v []*CreatePrometheusViewRequestPrometheusInstances) *CreatePrometheusViewRequest
	GetPrometheusInstances() []*CreatePrometheusViewRequestPrometheusInstances
	SetPrometheusViewName(v string) *CreatePrometheusViewRequest
	GetPrometheusViewName() *string
	SetResourceGroupId(v string) *CreatePrometheusViewRequest
	GetResourceGroupId() *string
	SetStatus(v string) *CreatePrometheusViewRequest
	GetStatus() *string
	SetTags(v []*CreatePrometheusViewRequestTags) *CreatePrometheusViewRequest
	GetTags() []*CreatePrometheusViewRequestTags
	SetVersion(v string) *CreatePrometheusViewRequest
	GetVersion() *string
	SetWorkspace(v string) *CreatePrometheusViewRequest
	GetWorkspace() *string
}

type CreatePrometheusViewRequest struct {
	// example:
	//
	// null
	AuthFreeReadPolicy *string `json:"authFreeReadPolicy,omitempty" xml:"authFreeReadPolicy,omitempty"`
	// example:
	//
	// true
	EnableAuthFreeRead *bool `json:"enableAuthFreeRead,omitempty" xml:"enableAuthFreeRead,omitempty"`
	// example:
	//
	// true
	EnableAuthToken *bool `json:"enableAuthToken,omitempty" xml:"enableAuthToken,omitempty"`
	// This parameter is required.
	PrometheusInstances []*CreatePrometheusViewRequestPrometheusInstances `json:"prometheusInstances,omitempty" xml:"prometheusInstances,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test-prom-view-name
	PrometheusViewName *string `json:"prometheusViewName,omitempty" xml:"prometheusViewName,omitempty"`
	// example:
	//
	// rg-acfm3gn5i6bigbi
	ResourceGroupId *string `json:"resourceGroupId,omitempty" xml:"resourceGroupId,omitempty"`
	// example:
	//
	// null
	Status *string                            `json:"status,omitempty" xml:"status,omitempty"`
	Tags   []*CreatePrometheusViewRequestTags `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// V2
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
	// example:
	//
	// cms-monitor-test-aysls-pub-cn-zhangjiakou-spe-monitor
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s CreatePrometheusViewRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusViewRequest) GoString() string {
	return s.String()
}

func (s *CreatePrometheusViewRequest) GetAuthFreeReadPolicy() *string {
	return s.AuthFreeReadPolicy
}

func (s *CreatePrometheusViewRequest) GetEnableAuthFreeRead() *bool {
	return s.EnableAuthFreeRead
}

func (s *CreatePrometheusViewRequest) GetEnableAuthToken() *bool {
	return s.EnableAuthToken
}

func (s *CreatePrometheusViewRequest) GetPrometheusInstances() []*CreatePrometheusViewRequestPrometheusInstances {
	return s.PrometheusInstances
}

func (s *CreatePrometheusViewRequest) GetPrometheusViewName() *string {
	return s.PrometheusViewName
}

func (s *CreatePrometheusViewRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePrometheusViewRequest) GetStatus() *string {
	return s.Status
}

func (s *CreatePrometheusViewRequest) GetTags() []*CreatePrometheusViewRequestTags {
	return s.Tags
}

func (s *CreatePrometheusViewRequest) GetVersion() *string {
	return s.Version
}

func (s *CreatePrometheusViewRequest) GetWorkspace() *string {
	return s.Workspace
}

func (s *CreatePrometheusViewRequest) SetAuthFreeReadPolicy(v string) *CreatePrometheusViewRequest {
	s.AuthFreeReadPolicy = &v
	return s
}

func (s *CreatePrometheusViewRequest) SetEnableAuthFreeRead(v bool) *CreatePrometheusViewRequest {
	s.EnableAuthFreeRead = &v
	return s
}

func (s *CreatePrometheusViewRequest) SetEnableAuthToken(v bool) *CreatePrometheusViewRequest {
	s.EnableAuthToken = &v
	return s
}

func (s *CreatePrometheusViewRequest) SetPrometheusInstances(v []*CreatePrometheusViewRequestPrometheusInstances) *CreatePrometheusViewRequest {
	s.PrometheusInstances = v
	return s
}

func (s *CreatePrometheusViewRequest) SetPrometheusViewName(v string) *CreatePrometheusViewRequest {
	s.PrometheusViewName = &v
	return s
}

func (s *CreatePrometheusViewRequest) SetResourceGroupId(v string) *CreatePrometheusViewRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePrometheusViewRequest) SetStatus(v string) *CreatePrometheusViewRequest {
	s.Status = &v
	return s
}

func (s *CreatePrometheusViewRequest) SetTags(v []*CreatePrometheusViewRequestTags) *CreatePrometheusViewRequest {
	s.Tags = v
	return s
}

func (s *CreatePrometheusViewRequest) SetVersion(v string) *CreatePrometheusViewRequest {
	s.Version = &v
	return s
}

func (s *CreatePrometheusViewRequest) SetWorkspace(v string) *CreatePrometheusViewRequest {
	s.Workspace = &v
	return s
}

func (s *CreatePrometheusViewRequest) Validate() error {
	if s.PrometheusInstances != nil {
		for _, item := range s.PrometheusInstances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreatePrometheusViewRequestPrometheusInstances struct {
	// example:
	//
	// arms-1d581fac20a462dcde743d9628
	PrometheusInstanceId *string `json:"prometheusInstanceId,omitempty" xml:"prometheusInstanceId,omitempty"`
	// example:
	//
	// cn-wulanchabu
	RegionId *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	// example:
	//
	// 167271234567890
	UserId *string `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CreatePrometheusViewRequestPrometheusInstances) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusViewRequestPrometheusInstances) GoString() string {
	return s.String()
}

func (s *CreatePrometheusViewRequestPrometheusInstances) GetPrometheusInstanceId() *string {
	return s.PrometheusInstanceId
}

func (s *CreatePrometheusViewRequestPrometheusInstances) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePrometheusViewRequestPrometheusInstances) GetUserId() *string {
	return s.UserId
}

func (s *CreatePrometheusViewRequestPrometheusInstances) SetPrometheusInstanceId(v string) *CreatePrometheusViewRequestPrometheusInstances {
	s.PrometheusInstanceId = &v
	return s
}

func (s *CreatePrometheusViewRequestPrometheusInstances) SetRegionId(v string) *CreatePrometheusViewRequestPrometheusInstances {
	s.RegionId = &v
	return s
}

func (s *CreatePrometheusViewRequestPrometheusInstances) SetUserId(v string) *CreatePrometheusViewRequestPrometheusInstances {
	s.UserId = &v
	return s
}

func (s *CreatePrometheusViewRequestPrometheusInstances) Validate() error {
	return dara.Validate(s)
}

type CreatePrometheusViewRequestTags struct {
	// example:
	//
	// test-key
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// example:
	//
	// test-value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CreatePrometheusViewRequestTags) String() string {
	return dara.Prettify(s)
}

func (s CreatePrometheusViewRequestTags) GoString() string {
	return s.String()
}

func (s *CreatePrometheusViewRequestTags) GetKey() *string {
	return s.Key
}

func (s *CreatePrometheusViewRequestTags) GetValue() *string {
	return s.Value
}

func (s *CreatePrometheusViewRequestTags) SetKey(v string) *CreatePrometheusViewRequestTags {
	s.Key = &v
	return s
}

func (s *CreatePrometheusViewRequestTags) SetValue(v string) *CreatePrometheusViewRequestTags {
	s.Value = &v
	return s
}

func (s *CreatePrometheusViewRequestTags) Validate() error {
	return dara.Validate(s)
}
