// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeServiceInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeServiceInstanceResponseBody
	GetRequestId() *string
	SetServiceInstanceId(v string) *UpgradeServiceInstanceResponseBody
	GetServiceInstanceId() *string
	SetStatus(v string) *UpgradeServiceInstanceResponseBody
	GetStatus() *string
	SetUpgradeRequiredParameters(v []*string) *UpgradeServiceInstanceResponseBody
	GetUpgradeRequiredParameters() []*string
}

type UpgradeServiceInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F224E002-AB0E-5FD9-A87E-54AEE56F6CAE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the service instance.
	//
	// example:
	//
	// si-5cbae874da0e47xxxxxx
	ServiceInstanceId *string `json:"ServiceInstanceId,omitempty" xml:"ServiceInstanceId,omitempty"`
	// The deployment state of the service instance. Valid values:
	//
	// 	- Created
	//
	// 	- Deploying
	//
	// 	- DeployedFailed
	//
	// 	- Deployed
	//
	// 	- Upgrading
	//
	// 	- Deleting
	//
	// 	- Deleted
	//
	// 	- DeletedFailed
	//
	// example:
	//
	// Created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The parameters required for the upgrade.
	UpgradeRequiredParameters []*string `json:"UpgradeRequiredParameters,omitempty" xml:"UpgradeRequiredParameters,omitempty" type:"Repeated"`
}

func (s UpgradeServiceInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeServiceInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeServiceInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeServiceInstanceResponseBody) GetServiceInstanceId() *string {
	return s.ServiceInstanceId
}

func (s *UpgradeServiceInstanceResponseBody) GetStatus() *string {
	return s.Status
}

func (s *UpgradeServiceInstanceResponseBody) GetUpgradeRequiredParameters() []*string {
	return s.UpgradeRequiredParameters
}

func (s *UpgradeServiceInstanceResponseBody) SetRequestId(v string) *UpgradeServiceInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeServiceInstanceResponseBody) SetServiceInstanceId(v string) *UpgradeServiceInstanceResponseBody {
	s.ServiceInstanceId = &v
	return s
}

func (s *UpgradeServiceInstanceResponseBody) SetStatus(v string) *UpgradeServiceInstanceResponseBody {
	s.Status = &v
	return s
}

func (s *UpgradeServiceInstanceResponseBody) SetUpgradeRequiredParameters(v []*string) *UpgradeServiceInstanceResponseBody {
	s.UpgradeRequiredParameters = v
	return s
}

func (s *UpgradeServiceInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
