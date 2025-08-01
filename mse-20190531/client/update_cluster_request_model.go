// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateClusterRequest
	GetAcceptLanguage() *string
	SetClusterAliasName(v string) *UpdateClusterRequest
	GetClusterAliasName() *string
	SetInstanceId(v string) *UpdateClusterRequest
	GetInstanceId() *string
	SetMaintenanceEndTime(v string) *UpdateClusterRequest
	GetMaintenanceEndTime() *string
	SetMaintenanceStartTime(v string) *UpdateClusterRequest
	GetMaintenanceStartTime() *string
	SetRequestPars(v string) *UpdateClusterRequest
	GetRequestPars() *string
}

type UpdateClusterRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The alias of the instance.
	//
	// example:
	//
	// cluster-1
	ClusterAliasName *string `json:"ClusterAliasName,omitempty" xml:"ClusterAliasName,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-78v1l83****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The end time of the O\\&M window.
	//
	// example:
	//
	// 06:00
	MaintenanceEndTime *string `json:"MaintenanceEndTime,omitempty" xml:"MaintenanceEndTime,omitempty"`
	// The start time of the O\\&M window.
	//
	// example:
	//
	// 02:00
	MaintenanceStartTime *string `json:"MaintenanceStartTime,omitempty" xml:"MaintenanceStartTime,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
}

func (s UpdateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateClusterRequest) GoString() string {
	return s.String()
}

func (s *UpdateClusterRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateClusterRequest) GetClusterAliasName() *string {
	return s.ClusterAliasName
}

func (s *UpdateClusterRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateClusterRequest) GetMaintenanceEndTime() *string {
	return s.MaintenanceEndTime
}

func (s *UpdateClusterRequest) GetMaintenanceStartTime() *string {
	return s.MaintenanceStartTime
}

func (s *UpdateClusterRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *UpdateClusterRequest) SetAcceptLanguage(v string) *UpdateClusterRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateClusterRequest) SetClusterAliasName(v string) *UpdateClusterRequest {
	s.ClusterAliasName = &v
	return s
}

func (s *UpdateClusterRequest) SetInstanceId(v string) *UpdateClusterRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateClusterRequest) SetMaintenanceEndTime(v string) *UpdateClusterRequest {
	s.MaintenanceEndTime = &v
	return s
}

func (s *UpdateClusterRequest) SetMaintenanceStartTime(v string) *UpdateClusterRequest {
	s.MaintenanceStartTime = &v
	return s
}

func (s *UpdateClusterRequest) SetRequestPars(v string) *UpdateClusterRequest {
	s.RequestPars = &v
	return s
}

func (s *UpdateClusterRequest) Validate() error {
	return dara.Validate(s)
}
