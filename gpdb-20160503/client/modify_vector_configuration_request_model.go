// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVectorConfigurationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyVectorConfigurationRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *ModifyVectorConfigurationRequest
	GetOwnerId() *int64
	SetVectorConfigurationStatus(v string) *ModifyVectorConfigurationRequest
	GetVectorConfigurationStatus() *string
}

type ModifyVectorConfigurationRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the IDs of all AnalyticDB for PostgreSQL instances in a region.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to enable vector engine optimization. Valid values:
	//
	// 	- **enabled**
	//
	// 	- **disabled**
	//
	// > 	- We recommend that you **do not enable*	- vector engine optimization in mainstream analysis and real-time data warehousing scenarios.
	//
	// > 	- We recommend that you **enable*	- vector engine optimization in AI Generated Content (AIGC) and vector retrieval scenarios that require the vector analysis engine.
	//
	// example:
	//
	// enabled
	VectorConfigurationStatus *string `json:"VectorConfigurationStatus,omitempty" xml:"VectorConfigurationStatus,omitempty"`
}

func (s ModifyVectorConfigurationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVectorConfigurationRequest) GoString() string {
	return s.String()
}

func (s *ModifyVectorConfigurationRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyVectorConfigurationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVectorConfigurationRequest) GetVectorConfigurationStatus() *string {
	return s.VectorConfigurationStatus
}

func (s *ModifyVectorConfigurationRequest) SetDBInstanceId(v string) *ModifyVectorConfigurationRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyVectorConfigurationRequest) SetOwnerId(v int64) *ModifyVectorConfigurationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVectorConfigurationRequest) SetVectorConfigurationStatus(v string) *ModifyVectorConfigurationRequest {
	s.VectorConfigurationStatus = &v
	return s
}

func (s *ModifyVectorConfigurationRequest) Validate() error {
	return dara.Validate(s)
}
