// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeDBInstanceMajorVersionPrecheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *UpgradeDBInstanceMajorVersionPrecheckRequest
	GetDBInstanceId() *string
	SetResourceOwnerId(v int64) *UpgradeDBInstanceMajorVersionPrecheckRequest
	GetResourceOwnerId() *int64
	SetTargetMajorVersion(v string) *UpgradeDBInstanceMajorVersionPrecheckRequest
	GetTargetMajorVersion() *string
	SetUpgradeMode(v string) *UpgradeDBInstanceMajorVersionPrecheckRequest
	GetUpgradeMode() *string
}

type UpgradeDBInstanceMajorVersionPrecheckRequest struct {
	// The ID of the instance. You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/610396.html) operation to query the ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// pgm-bp1c808s731l****
	DBInstanceId    *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	ResourceOwnerId *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The new major engine version of the instance. The new major engine version must be later than the original major engine version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 12.0
	TargetMajorVersion *string `json:"TargetMajorVersion,omitempty" xml:"TargetMajorVersion,omitempty"`
	UpgradeMode        *string `json:"UpgradeMode,omitempty" xml:"UpgradeMode,omitempty"`
}

func (s UpgradeDBInstanceMajorVersionPrecheckRequest) String() string {
	return dara.Prettify(s)
}

func (s UpgradeDBInstanceMajorVersionPrecheckRequest) GoString() string {
	return s.String()
}

func (s *UpgradeDBInstanceMajorVersionPrecheckRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *UpgradeDBInstanceMajorVersionPrecheckRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpgradeDBInstanceMajorVersionPrecheckRequest) GetTargetMajorVersion() *string {
	return s.TargetMajorVersion
}

func (s *UpgradeDBInstanceMajorVersionPrecheckRequest) GetUpgradeMode() *string {
	return s.UpgradeMode
}

func (s *UpgradeDBInstanceMajorVersionPrecheckRequest) SetDBInstanceId(v string) *UpgradeDBInstanceMajorVersionPrecheckRequest {
	s.DBInstanceId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionPrecheckRequest) SetResourceOwnerId(v int64) *UpgradeDBInstanceMajorVersionPrecheckRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionPrecheckRequest) SetTargetMajorVersion(v string) *UpgradeDBInstanceMajorVersionPrecheckRequest {
	s.TargetMajorVersion = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionPrecheckRequest) SetUpgradeMode(v string) *UpgradeDBInstanceMajorVersionPrecheckRequest {
	s.UpgradeMode = &v
	return s
}

func (s *UpgradeDBInstanceMajorVersionPrecheckRequest) Validate() error {
	return dara.Validate(s)
}
